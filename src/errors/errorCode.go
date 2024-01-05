package errors

// ErrorCode is a custom type to explicitly link codes to enums in either the application or in capila
type ErrorCode = string

const (
	ErrorCodeBadGateway          ErrorCode = "ERROR_BAD_GATEWAY"
	ErrorCodeBadRequest          ErrorCode = "ERROR_BAD_REQUEST"
	ErrorCodeGatewayTimeout      ErrorCode = "ERROR_GATEWAY_TIMEOUT"
	ErrorCodeInternalServerError ErrorCode = "ERROR_INTERNAL_SERVER_ERROR"
	ErrorCodeNotFound            ErrorCode = "ERROR_NOT_FOUND"
	ErrorCodeServiceUnavailable  ErrorCode = "ERROR_SERVICE_UNAVAILABLE"

	ErrorCodeInputValidation ErrorCode = "ERROR_INPUT_VALIDATION"
	ErrorCodeUnknown         ErrorCode = "ERROR_UNKNOWN"
)
