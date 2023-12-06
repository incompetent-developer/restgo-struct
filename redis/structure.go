package redis

import "github.com/go-redis/redis/v8"

type Redis struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	Database       int    `json:"database"`
	CredentialPath string `json:"credential_path"`
	Client         *redis.Client
}
