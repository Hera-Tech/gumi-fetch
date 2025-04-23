package env

import (
	"os"
	"strconv"
	"time"
)

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return boolVal
}
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
func GetDuration(key string, fallback time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsDuration, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}

	return valAsDuration
}

func GetUInt32(key string, fallback uint32) uint32 {

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsUInt32, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return fallback
	}

	return uint32(valAsUInt32)
}

func GetUInt8(key string, fallback uint8) uint8 {

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsUInt8, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return fallback
	}

	return uint8(valAsUInt8)
}
