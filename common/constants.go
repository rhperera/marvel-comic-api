package common

const JWTContextKey string = "jwtConfig"
const JWTTokenLookup string = "jwt"
const UserContext string = "user"

const ErrorEchoServerInit string = "Error - Echo Server has not been initialized. Instance is nil. Run Init() first"

//HTTP errors
const ErrorCodeInvalidJWT int = 0
const ErrorInvalidJWT string = "Error - Invalid JWT token"
