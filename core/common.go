package core

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"log"
)

// Encoded function encodes content with salt string.
func Encoded(content string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	hash.Write([]byte(salt))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// GetUUID function generates an uuid string
// A panic is triggered if the current time cannot be determined.
func GetUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	return id.String()
}
