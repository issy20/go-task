package utils

import "github.com/google/uuid"

func GenRandomUUID() string {
	uuid := uuid.Must(uuid.NewRandom()).String()
	return string(uuid)
}
