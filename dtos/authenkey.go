package dtos

import "time"

type AuthenKey struct {
	Key       string     `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
}
