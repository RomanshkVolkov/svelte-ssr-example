package repository

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// generateSalt creates a cryptographically secure random salt of the specified size.
func generateSalt(saltSize uint32) ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

// hashPassword hashes a plaintext password using Argon2id and returns the base64-encoded hash string.
func HashPassword(password string) (string, error) {
	// Recommended Argon2id parameters
	timeCost := uint32(1)           // Number of passes over the memory
	memoryCost := uint32(64 * 1024) // Memory in KiB (e.g., 64MB)
	threads := uint8(4)             // Number of threads/lanes
	keyLength := uint32(32)         // Length of the derived key

	salt, err := generateSalt(16) // 16 bytes is a common salt length
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, threads, keyLength)

	// Encode the parameters, salt, and hash into the PHC String Format
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memoryCost,
		timeCost,
		threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encodedHash, nil
}
