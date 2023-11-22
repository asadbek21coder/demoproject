package entities

type Response struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Body         interface{} `json:"body"`
}

const (
	ErrorSuccessCode                  = 1000
	ErrorCodeAccessTokenExpired       = 1001
	ErrorCodeRefreshTokenExpired      = 1002
	ErrorCodeOTPExpired               = 1003
	ErrorCodeImageExtensionNotAllowed = 1004
	ErrorCodeImageSizeExceed          = 1005
	ErrorCodeNotFound                 = 1006
	ErrorCodeInvalidJson              = 1007
	ErrorCodeWrongOTP                 = 1008
	ErrorCodeWrongPassword            = 1009
	ErrorCodeInternal                 = 1010
	ErrorCodeBadRequest               = 1011
	ErrorCodeUnauthorized             = 1012
	ErrorCodeNotAllowed               = 1013
	ErrorCodeAlreadyRegistered        = 1014
	ErrorCodeAlreadyLiked             = 1015
	ErrorCodeNotLiked                 = 1016
)
