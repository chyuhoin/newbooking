package utils

import "errors"

var (
	TokenExpired     error = errors.New("token is expired")
	TokenNotValidYet error = errors.New("token is not active yet")
	TokenMalformed   error = errors.New("malformed token")
	TokenInvalid     error = errors.New("can't handle this token")
	NoSuchUser       error = errors.New("no such user")
)

const (
	// SignKey 这个是需要保密的一段信息
	SignKey       string = "a87x80wfebei90f8532f16f423b125616dea9b75"
	GinContextKey string = "claims"
)
