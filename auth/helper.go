package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/incompetent-developer/restgo-struct/access"
	"github.com/incompetent-developer/restgo-struct/header"
	"github.com/incompetent-developer/restgo-struct/logger"
)

// GetTokenExpiresIn is get token expires in second
func GetTokenExpiresIn() int64 {
	return tokenExpiresIn
}

// SetTokenExpiresIn is set token expires in second
func SetTokenExpiresIn(duration int64) {
	tokenExpiresIn = duration
}

// GetRefreshTokenExpiresIn is get refresh token expires in second
func GetRefreshTokenExpiresIn() int64 {
	return refreshTokenExpiresIn
}

// SetRefreshTokenExpiresIn is set refresh token expires in second
func SetRefreshTokenExpiresIn(duration int64) {
	refreshTokenExpiresIn = duration
}

// GetConfirmationTokenExpiresIn is get confirmation token expires in second
func GetConfirmationTokenExpiresIn() int64 {
	return confirmationTokenExpiresIn
}

// SetConfirmationTokenExpiresIn is set confirmation token expires in second
func SetConfirmationTokenExpiresIn(duration int64) {
	confirmationTokenExpiresIn = duration
}

// GetResetPasswordTokenExpiresIn is get resetPassword token expires in second
func GetResetPasswordTokenExpiresIn() int64 {
	return resetPasswordTokenExpiresIn
}

// SetResetPasswordTokenExpiresIn is set resetPassword token expires in second
func SetResetPasswordTokenExpiresIn(duration int64) {
	resetPasswordTokenExpiresIn = duration
}

// Generate a token with grant type
func (client *Client) generateToken(grantType string, signingMethodName string, authenticatorID uint64, authenticatorClassID string, secret []byte, permissions ...string) (resp string, err error) {

	/*
		Expires in
	*/
	var (
		expiresIn time.Duration
	)

	/*
		Scopes
	*/
	var (
		scopes = make(map[string]bool)
	)

	for _, permission := range permissions {
		scopes[permission] = true
	}

	/*
		Set custom claims
	*/
	claims := &Claims{}

	switch grantType {
	case AccessToken:
		expiresIn = time.Duration(tokenExpiresIn) * time.Second
		claims.ID = authenticatorID
		claims.ClassID = authenticatorClassID
		claims.Scopes = scopes
		claims.ExpiresAt = time.Now().Add(expiresIn).Unix()

	case RefreshToken:
		expiresIn = time.Duration(refreshTokenExpiresIn) * time.Second
		claims.ID = authenticatorID
		claims.ClassID = authenticatorClassID
		claims.Scopes = scopes
		claims.ExpiresAt = time.Now().Add(expiresIn).Unix()

	case ConfirmationToken:
		expiresIn = time.Duration(confirmationTokenExpiresIn) * time.Second
		claims.ID = authenticatorID
		claims.ClassID = authenticatorClassID
		claims.Scopes = scopes
		claims.ExpiresAt = time.Now().Add(expiresIn).Unix()

	case ResetPasswordToken:
		expiresIn = time.Duration(resetPasswordTokenExpiresIn) * time.Second
		claims.ID = authenticatorID
		claims.ClassID = authenticatorClassID
		claims.Scopes = scopes
		claims.ExpiresAt = time.Now().Add(expiresIn).Unix()

	default:
		return resp, fmt.Errorf("Grant type not support yet")

	}

	var (
		tokenSigned string
	)

	switch signingMethodName {
	case jwt.SigningMethodRS256.Name:
		/*
			Create token with claims
		*/
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

		/*
			Generate encoded token and send it as response.
		*/
		signedKey, err := jwt.ParseRSAPrivateKeyFromPEM(secret)
		if err != nil {
			return resp, err
		}

		tokenEncoded, err := token.SignedString(signedKey)
		if err != nil {
			return resp, err
		}

		tokenSigned = tokenEncoded

	case jwt.SigningMethodHS256.Name:
		/*
			Create token with claims
		*/
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		/*
			Generate encoded token and send it as response.
		*/
		tokenEncoded, err := token.SignedString(secret)
		if err != nil {
			return resp, err
		}

		tokenSigned = tokenEncoded

	default:
		return resp, fmt.Errorf("Signing method not support yet")

	}

	/*
		Cache
	*/
	ckey := fmt.Sprintf("%s:%d:%s", authenticatorClassID, authenticatorID, grantType)
	ctx := context.Background()

	if _, err := client.Redis.Client.Set(ctx, ckey, tokenSigned, expiresIn).Result(); err != nil {
		return resp, err
	}

	/*
		Response
	*/
	return tokenSigned, nil
}

// GetClaims is return jwt map claims
func GetClaims(ctx echo.Context) (accessDetail access.AccessDetail, err error) {
	/*
		JWT token
	*/
	user, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return accessDetail, fmt.Errorf("Cannot parse token")
	}

	/*
		JWT claims
	*/
	claims := user.Claims.(jwt.MapClaims)
	if !ok {
		return accessDetail, fmt.Errorf("Cannot claims token")
	}

	/*
		Encode to access detail
	*/
	claimsByte, err := json.Marshal(claims)
	if err != nil {
		return accessDetail, err
	}

	if err := json.Unmarshal(claimsByte, &accessDetail); err != nil {
		return accessDetail, err
	}

	return accessDetail, nil
}

// // GetClaims is return jwt map claims
// func GetClaims(ctx echo.Context) (*Claims, error) {
// 	/*
// 		JWT token
// 	*/
// 	user, ok := ctx.Get("user").(*jwt.Token)
// 	if !ok {
// 		return nil, fmt.Errorf("Cannot parse token")
// 	}

// 	/*
// 		JWT claims
// 	*/
// 	claims := user.Claims.(*Claims)
// 	if !ok {
// 		return nil, fmt.Errorf("Cannot claims token")
// 	}

// 	return claims, nil
// }

// IsTokenAlive return access token grant type is alive or not
func (client *Client) IsTokenAlive(ctx echo.Context, claims *Claims) bool {

	return client.IsAlive(AccessToken, ctx, claims)

}

// IsRefreshTokenAlive return refresh token grant type is alive or not
func (client *Client) IsRefreshTokenAlive(ctx echo.Context, claims *Claims) bool {

	return client.IsAlive(RefreshToken, ctx, claims)

}

// IsAlive return token with grant type is alive or not
func (client *Client) IsAlive(grantType string, ctx echo.Context, claims *Claims) bool {
	/*
		Check from cache
	*/
	key := fmt.Sprintf("%s:%d:%s", claims.ClassID, claims.ID, grantType)
	tctx := context.Background()
	token, err := client.Redis.Client.Get(tctx, key).Result()
	if err != nil {

		return false

	}

	/*
		Get header token
	*/
	authorization := ctx.Request().Header.Get(header.Authorization)
	seperate := strings.Split(authorization, " ")
	headerToken := seperate[(len(seperate) - 1)]

	if token != headerToken {

		return false

	}

	return true
}

// FetchAuth
func (client *Client) FetchAuth(accessDetail *access.AccessDetail) (string, bool) {

	accountID, err := client.Redis.Client.Get(context.Background(), accessDetail.AccessUuid).Result()
	if err != nil {
		return "", false
	}

	// Decrypt account_id
	accountIDDecrypt, err := DecryptData(accountID)
	if err != nil {
		logger.Errorf("Error when decrypt account : %v", err)
		return "", false
	}

	// Decrypt account_id from AccessDetail
	accountIDAccessDetailDecrypt, err := DecryptData(accessDetail.AccountID)
	if err != nil {
		logger.Errorf("Error when decrypt access account : %v", err)
		return "", false
	}

	if accountIDAccessDetailDecrypt != accountIDDecrypt {
		logger.Errorf("Error when decrypt access account : %v", err)
		return "", false
	}

	return accountIDDecrypt, true
}

// DecryptData decrypt permission to string
func DecryptData(value string) (string, error) {

	if value == "" {
		return "", fmt.Errorf("Value not found")
	}

	strValue := value

	for i := 1; i <= 2; i++ {
		strByte, err := base64.StdEncoding.DecodeString(string(strValue))
		if err != nil {
			return "", fmt.Errorf("Error when decode value : %v", err)
		}

		strValue = string(strByte)
	}

	return strValue, nil
}

// RevokeTokenWithClaims is revoke a token with claims
func (client *Client) RevokeTokenWithClaims(claims *Claims) error {

	return client.RevokeTokenWithAuthenticator(claims.ID, claims.ClassID)

}

// RevokeGrantTypeTokenWithClaims is revoke a confirmation token with claims
func (client *Client) RevokeGrantTypeTokenWithClaims(grantType string, claims *Claims) error {

	return client.RevokeGrantTypeTokenWithAuthenticator(grantType, claims.ID, claims.ClassID)

}

// RevokeTokenWithAuthenticator is revoke a token with authenticator
func (client *Client) RevokeTokenWithAuthenticator(authenticatorID uint64, authenticatorClassID string) error {

	if err := client.RevokeGrantTypeTokenWithAuthenticator(AccessToken, authenticatorID, authenticatorClassID); err != nil {
		return err
	}

	if err := client.RevokeGrantTypeTokenWithAuthenticator(RefreshToken, authenticatorID, authenticatorClassID); err != nil {
		return err
	}

	return nil
}

// RevokeGrantTypeTokenWithAuthenticator is revoke a token with authenticator
func (client *Client) RevokeGrantTypeTokenWithAuthenticator(grantType string, authenticatorID uint64, authenticatorClassID string) error {
	/*
		Remove from cache
	*/
	ckey := fmt.Sprintf("%s:%d:%s", authenticatorClassID, authenticatorID, grantType)
	ctx := context.Background()

	if _, err := client.Redis.Client.Del(ctx, ckey).Result(); err != nil {
		return err
	}

	return nil
}

// GenerateDigestParts is generate digest part for digest authentication
func GenerateDigestParts(resp *http.Response) map[string]string {
	var (
		result = make(map[string]string)
	)

	if len(resp.Header["Www-Authenticate"]) > 0 {
		wantedHeaders := []string{"nonce", "realm", "qop"}
		responseHeaders := strings.Split(resp.Header["Www-Authenticate"][0], ",")
		for _, r := range responseHeaders {
			for _, w := range wantedHeaders {
				if strings.Contains(r, w) {
					result[w] = strings.Split(r, `"`)[1]
				}
			}
		}
	}

	return result
}
