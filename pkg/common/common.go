package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/erik-sostenes/notifications-api/pkg/domain/wrongs"
	"github.com/google/uuid"
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

type Map map[string]any

// Identifier receives a value to verify if the format is correct
type Identifier string

// Validate method validates if the value is an Uuid, if incorrect returns an wrongs.StatusUnprocessableEntity
func (i Identifier) Validate() (string, error) {
	u, err := ParseUuID(string(i))
	if err != nil {
		return u.String(), wrongs.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s uuid unique identifier, %v", string(i), err))
	}
	return u.String(), nil
}

// TimeStampLayout format the dates
const TimeStampLayout = "2006-01-02 15:04:05"

// Timestamp receives a value to verify if the format is correct
type Timestamp string

// Validate method validates if the value is a time.Time, if incorrect returns an wrongs.StatusUnprocessableEntity
func (t Timestamp) Validate() (int64, error) {
	v, err := time.Parse(TimeStampLayout, string(t))
	if err != nil {
		return 0, wrongs.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s value format, %v", string(t), err))
	}
	return v.Unix(), nil
}

// String receives a value to verify if the format is correct
type String string

// The Validate method validates if the value is a string and is not empty, if incorrect returns an wrongs.StatusUnprocessableEntity
func (s String) Validate() (string, error) {
	if strings.TrimSpace(string(s)) == "" {
		return "", wrongs.StatusUnprocessableEntity("Value not found")
	}
	return string(s), nil
}
