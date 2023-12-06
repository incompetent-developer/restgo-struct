package header

import (
	"fmt"
	"strconv"
	"strings"
)

// GetAuthorizationWithPrefix is merge token with authorization type
func GetAuthorizationWithPrefix(token string) string {
	return fmt.Sprintf("%s %s", AuthorizationType, token)
}

// GetScopes is decode scopes from header
func GetScopes(head string) (scopes []string) {
	return strings.Split(head, " ")
}

// MergeScopes is merge scopes to header
func MergeScopes(head string, scopes ...string) (scope string) {
	return strings.Join(append(GetScopes(head), scopes...), " ")
}

// GetClassAndAuthenticatedID cecode class and authenticated id
func GetClassAndAuthenticatedID(head string) (classID string, authenticatedID uint64, err error) {
	heads := strings.Split(head, "|")
	authenticatorID, err := strconv.ParseUint(heads[1], 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("Decode authenticator id error : %s", err)
	}
	return heads[0], authenticatorID, nil
}
