package errors

type ErrorType int

const (
	BadRequest ErrorType = iota
	Unauthorized
	Forbidden
	Suspended
	NotFound
	Conflict
	UnprocessableEntity

	InternalServerError
	ServiceUnavailable
)

var ErrorTypes = map[ErrorType]string{
	BadRequest:          "bad_request",
	Unauthorized:        "unauthorized",
	Forbidden:           "forbidden",
	Suspended:           "suspended",
	NotFound:            "not_found",
	Conflict:            "conflict",
	UnprocessableEntity: "invalid_params",
}

type ErrorCode int

const (
	Missing ErrorCode = iota
	MissingField
	Invalid
	AlreadyExists
)

var ErrorCodes = map[ErrorCode]string{
	Missing:       "missing",
	MissingField:  "missing_field",
	Invalid:       "invalid",
	AlreadyExists: "already_exists",
}
