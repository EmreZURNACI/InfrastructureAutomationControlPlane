package token

import "fmt"

var errTokenSignatureInvalid error = fmt.Errorf("token signature invalid")
var errTokenNotProvided error = fmt.Errorf("token not provided")
var errTokenInvalid error = fmt.Errorf("token invalid")
var errTokenClaimsInvalid error = fmt.Errorf("token token claims invalid")
var errTokenExpInvalid error = fmt.Errorf("token expire time invalid")
var errTokenExpired error = fmt.Errorf("token expired")
var errTokenNickNotFound error = fmt.Errorf("token signature invalid")
var errTokenRoleNotFound error = fmt.Errorf("token signature invalid")
var errTokenNotValid error = fmt.Errorf("token not valid")
