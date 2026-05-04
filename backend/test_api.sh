#!/bin/bash

# API Test Script for User Management
# Usage: ./test_api.sh

set -e

BASE_URL="http://localhost:8081"
API_BASE="${BASE_URL}/api/v1"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_test() {
    echo -e "${BLUE}[TEST]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_info() {
    echo -e "${YELLOW}[INFO]${NC} $1"
}

# Function to check if server is running
check_server() {
    print_info "Checking if server is running..."
    if curl -s "${BASE_URL}/health" > /dev/null; then
        print_success "Server is running!"
    else
        print_error "Server is not running. Please start the server first:"
        echo "cd backend && go run ./cmd/main.go"
        exit 1
    fi
}

# Function to make HTTP request and show response
make_request() {
    local method=$1
    local url=$2
    local data=$3
    local description=$4

    print_test "$description"
    echo "Request: $method $url"

    if [ -n "$data" ]; then
        echo "Body: $data"
        response=$(curl -s -w "\n%{http_code}" -X "$method" \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$url")
    else
        response=$(curl -s -w "\n%{http_code}" -X "$method" "$url")
    fi

    # Extract response body and status code
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')

    echo "Status: $http_code"
    echo "Response:"
    echo "$body" | jq '.' 2>/dev/null || echo "$body"
    echo ""

    # Check if request was successful (2xx status codes)
    if [[ $http_code =~ ^2[0-9][0-9]$ ]]; then
        print_success "Request successful"
    else
        print_error "Request failed with status $http_code"
    fi

    echo "----------------------------------------"
}

# Main test execution
main() {
    echo "================================================="
    echo "         User API Testing Script"
    echo "================================================="
    echo ""

    # Check if jq is installed (for JSON formatting)
    if ! command -v jq &> /dev/null; then
        print_info "jq not found. JSON responses won't be formatted."
        print_info "Install jq for better output: sudo apt-get install jq"
        echo ""
    fi

    # Check if server is running
    check_server
    echo ""

    # Test 1: Health Check
    make_request "GET" "${BASE_URL}/health" "" "Health Check Endpoint"

    # Test 2: API Routes Documentation
    make_request "GET" "${BASE_URL}/api/routes" "" "API Documentation Endpoint"

    # Test 3: Get All Users
    make_request "GET" "${API_BASE}/users" "" "Get All Users"

    # Test 4: Get User by ID (existing)
    make_request "GET" "${API_BASE}/users/1" "" "Get User by ID (existing user)"

    # Test 5: Get User by ID (non-existing)
    make_request "GET" "${API_BASE}/users/999" "" "Get User by ID (non-existing user)"

    # Test 6: Create New User (valid data)
    user_data='{
        "email": "testuser@example.com",
        "password": "password123",
        "name": "Test User"
    }'
    make_request "POST" "${API_BASE}/users" "$user_data" "Create New User (valid data)"

    # Test 7: Create New User (invalid email)
    invalid_user_data='{
        "email": "invalid-email",
        "password": "password123",
        "name": "Test User"
    }'
    make_request "POST" "${API_BASE}/users" "$invalid_user_data" "Create New User (invalid email)"

    # Test 8: Create New User (short password)
    short_password_data='{
        "email": "test2@example.com",
        "password": "123",
        "name": "Test User"
    }'
    make_request "POST" "${API_BASE}/users" "$short_password_data" "Create New User (password too short)"

    # Test 9: Create New User (missing name)
    missing_name_data='{
        "email": "test3@example.com",
        "password": "password123"
    }'
    make_request "POST" "${API_BASE}/users" "$missing_name_data" "Create New User (missing name)"

    # Test 10: Update User (existing)
    update_data='{
        "email": "updated@example.com",
        "name": "Updated Name"
    }'
    make_request "PUT" "${API_BASE}/users/1" "$update_data" "Update User (existing user)"

    # Test 11: Update User (non-existing)
    make_request "PUT" "${API_BASE}/users/999" "$update_data" "Update User (non-existing user)"

    # Test 12: Delete User (existing)
    make_request "DELETE" "${API_BASE}/users/2" "" "Delete User (existing user)"

    # Test 13: Delete User (non-existing)
    make_request "DELETE" "${API_BASE}/users/999" "" "Delete User (non-existing user)"

    # Test 14: Invalid HTTP Method
    make_request "PATCH" "${API_BASE}/users" "" "Invalid HTTP Method (PATCH not implemented)"

    # Test 15: Test CORS Headers (OPTIONS request)
    print_test "Testing CORS Headers (OPTIONS request)"
    echo "Request: OPTIONS ${API_BASE}/users"

    cors_response=$(curl -s -I -X OPTIONS \
        -H "Origin: http://localhost:3000" \
        -H "Access-Control-Request-Method: POST" \
        -H "Access-Control-Request-Headers: Content-Type" \
        "${API_BASE}/users")

    echo "CORS Headers:"
    echo "$cors_response" | grep -i "access-control" || print_info "No CORS headers found"
    echo ""
    print_success "CORS test completed"
    echo "----------------------------------------"

    echo ""
    echo "================================================="
    echo "                Test Summary"
    echo "================================================="
    print_info "All API tests completed!"
    print_info "Check the responses above for any errors."
    print_info "Successful requests should return 2xx status codes."
    echo ""
    print_info "To run the server: cd backend && go run ./cmd/main.go"
    print_info "To build the server: cd backend && go build -o bin/server ./cmd/main.go"
    echo "================================================="
}

# Run if script is executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
