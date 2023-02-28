package shared

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

// GetEnv method that reads the environment variables needed in the project.
//
// Note: if an environment variable is not found, a panic will occur.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf("missing environment variable '%s'", key))
	}
	return value
}

// GenerateUuID generate a new UuID.
func GenerateUuID() string {
	return uuid.New().String()
}

// ParseUuID validate if the format the values is a UuID
func ParseUuID(value string) (uuid.UUID, error) {
	return uuid.Parse(value)
}
