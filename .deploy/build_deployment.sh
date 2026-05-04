#!/bin/bash

# Setting colours
greenColour="\e[0;32m\033[1m"
endColour="\033[0m\e[0m"
redColour="\e[0;31m\033[1m"
blueColour="\e[0;34m\033[1m"
yellowColour="\e[0;33m\033[1m"
purpleColour="\e[0;35m\033[1m"
turquoiseColour="\e[0;36m\033[1m"
grayColour="\e[0;37m\033[1m"

deployment_dir=".deploy"

# message with format
message() {
  echo -e "${yellowColour}[*]${endColour}${grayColour} $1${endColour}"
}

errorMessage() {
  echo -e "${redColour}[x] ${endColour}${grayColour}$1${endColour}" >&2
}

validate_dependencies() {
  if ! command -v yq &> /dev/null; then
    errorMessage "Missing dep: yq not found binary"
    message "Install with: go install github.com/mikefarah/yq/v4@latest"
    exit 1
  fi
}

validate_args() {
  local missing=()
  for var in stack environment image host; do
    if [[ -z "${!var}" ]]; then
      missing+=("$var")
    fi
  done

  if (( ${#missing[@]} > 0 )); then
    errorMessage "Missing required args: ${missing[*]}"
    exit 1
  fi
}

validate_env_format() {
  local key="$1"
  if ! [[ "$key" =~ ^DEPLOY_[A-Za-z0-9_]+_[A-Za-z0-9_]+$ ]]; then
    errorMessage "Invalid format on enviroment variable: $key"
    errorMessage "The correct format is: DEPLOY_<SERVICE>_<VARIABLE>"
    return 1
  fi
  return 0
}

run_build_template() {
  validate_args
  validate_dependencies

  # Default values
  local tls="${tls:-internal}"

  # Read env files
  declare -A service_envs
  local relevant_envs=$(env | grep -E '^DEPLOY_[^=]+=')
  local relevant_envs_count=$(echo "$relevant_envs" | wc -l)

  # validate if not empy vars
  message "processing $relevant_envs_count vars"

  while IFS='=' read -r key value; do
    # Validate var struct
    if ! validate_env_format "$key"; then
      continue
    fi

    # Extract components (DEPLOY_<SERVICIO>_<VARIABLE>)
    local service_part="${key#DEPLOY_}"
    local service_name="${service_part%%_*}"
    local var_name="${service_part#*_}"

    # Save on associative array
    service_envs["$service_name"]+="$var_name=$value\n"
  done <<< "$relevant_envs"

  # message "find array service_envs:"
  # for service in "${!service_envs[@]}"; do
  #     echo -e "${yellowColour}Servicio: $service${endColour}"
  #     echo -e "${service_envs[$service]}"
  # done

  # Validate template
  local template_file="$deployment_dir/deployment.template.yml"
  if [[ ! -f "$template_file" ]]; then
    errorMessage "Template file not found: $template_file"
    exit 1
  fi

  # Create secure temp file
  local temp_yaml
  temp_yaml=$(mktemp --suffix=.yml) || {
    errorMessage "Error on create temp file"
    exit 1
  }

  trap 'rm -f "$temp_yaml"; message "Cleaned temp file"' EXIT

  message "Has been created temp file $temp_yaml"

  # copy temp file
  cp "$template_file" "$temp_yaml" || {
    errorMessage "Error on copy template"
    exit 1
  }

  # Replace generic values
  sed -i -e "s|STACK_PLACEHOLDER|$stack-$environment|g" \
         -e "s|IMAGE_PLACEHOLDER|$image|g" \
         -e "s|IMAGE_API_PLACEHOLDER|$image_api|g" \
         -e "s|HOST_PLACEHOLDER|$host|g" \
         -e "s|TLS_PLACEHOLDER|$tls|g" "$temp_yaml"

  # Inject vars on service
  for service_name in "${!service_envs[@]}"; do
    local full_service_name="$stack-$environment-$service_name"
    local env_list=$(echo -e "${service_envs[$service_name]}" | sort | uniq | sed '/^$/d')

    message "Start processing var inject on service $full_service_name"

    # validate if service exist on template
    if ! yq eval ".services.$full_service_name" $temp_yaml | grep -q 'null'; then
      while IFS= read -r env_var; do
        [[ -n "$env_var" ]] || continue
        message "var debug: $en_var"
        yq eval -i ".services.$full_service_name.environment += [\"$env_var\"]" $temp_yaml
      done <<< "$env_list"
    else
      errorMessage "Advertencia: Servicio '$full_service_name' no encontrado en template"
    fi
  done

  #
  message "File deployment generated succesfully"

  cp "$temp_yaml" "$output"
}

# process args
while getopts "s:e:i:a:h:t:o:" arg; do
  case $arg in
    s) stack="$OPTARG" ;;
    e) environment="$OPTARG" ;;
    i) image="$OPTARG" ;;
    a) image_api="$OPTARG" ;;
    h) host="$OPTARG" ;;
    t) tls="$OPTARG" ;;
    o) output="$OPTARG" ;;
    *)
      errorMessage "Usage: $0 -s <stack> -e <environment> -i <image> -a <image_api> -h <host> -t <tls> -o <output.yml>"
      errorMessage "Example: $0 -s stack -e dev -i gr.images.com -a gr.images.com -h caddy-host.com -t examples@email.com -o output_file.yml"
      exit 1
      ;;
  esac
done

# exec define work
run_build_template
