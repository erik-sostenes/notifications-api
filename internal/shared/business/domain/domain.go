package domain

import (
	"encoding/json"
	"fmt"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/errors"
	"strconv"
	"strings"
	"time"
)

type Map map[string]any

func (m Map) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// Identifier receives a value to verify if the format is correct
type Identifier string

// Validate method validates if the value is an Uuid, if incorrect returns an errors.StatusUnprocessableEntity
func (i Identifier) Validate() (string, error) {
	u, err := shared.ParseUuID(string(i))
	if err != nil {
		return u.String(), errors.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s uuid unique identifier, %v", string(i), err))
	}
	return u.String(), nil
}

// TimeStampLayout format the dates
const TimeStampLayout = "2006-01-02 15:04:05"

// Timestamp receives a value to verify if the format is correct
type Timestamp string

// Validate method validates if the value is a time.Time, if incorrect returns an errors.StatusUnprocessableEntity
func (t Timestamp) Validate() (int64, error) {
	v, err := time.Parse(TimeStampLayout, string(t))
	if err != nil {
		return v.Unix(), errors.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s value format, %v", string(t), err))
	}
	return v.Unix(), nil
}

// Bool receives a value to verify if the format is correct
type Bool string

// Validate method validates if the value is a bool, if incorrect returns an errors.StatusUnprocessableEntity
func (b Bool) Validate() (bool, error) {
	v, err := strconv.ParseBool(string(b))
	if err != nil {
		return v, errors.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s value format, %v", string(b), err))
	}
	return v, nil
}

// Float receives a value to verify if the format is correct
type Float string

// Validate method validates if the value is a float, if incorrect returns an errors.StatusUnprocessableEntity
func (f Float) Validate() (float64, error) {
	v, err := strconv.ParseFloat(string(f), 64)
	if err != nil {
		return v, errors.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s value format, %v", string(f), err))
	}
	return v, nil
}

// String receives a value to verify if the format is correct
type String string

// The Validate method validates if the value is a string and is not empty, if incorrect returns an errors.StatusUnprocessableEntity
func (s String) Validate() (string, error) {
	if strings.TrimSpace(string(s)) == "" {
		return "", errors.StatusUnprocessableEntity("Value not found")
	}
	return string(s), nil
}
