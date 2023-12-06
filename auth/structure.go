package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/incompetent-developer/restgo-struct/redis"
)

// Claims are custom claims extending default ones.
type Claims struct {
	ID      uint64          `json:"id"`
	ClassID string          `json:"class_id"`
	Scopes  map[string]bool `json:"scopes"`
	jwt.StandardClaims
}

// Token structure
type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// Client is keep configuration
type Client struct {
	Redis redis.Redis
}
