package core

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"log"
)

// Response is a common response content for ticktok application
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// ResponseError creates a common response content represents error
func ResponseError(err error) Response {
	return Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}

// ResponseOK creates a common response content represents success
func ResponseOK() Response {
	return Response{
		StatusCode: 0,
	}
}

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
