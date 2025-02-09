package randomness

import "github.com/google/uuid"

func NewRandomID() string {
	uuid := uuid.New()
	return uuid.String()
}
