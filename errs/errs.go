package errs

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	// SuccessCode means request success = 1000
	SuccessCode = iota + 1000
	// ServerInternalErrorCode means 500 internal server error = 1001
	ServerInternalErrorCode
	// InvalidAttributeErrorCode means 400 bad request = 1002
	InvalidAttributeErrorCode
	// RecordNotFoundErrorCode means 404 not found = 1003
	RecordNotFoundErrorCode
	// SuccessCreatedCode means 201 created = 1004
	SuccessCreatedCode
	// UnauthorizedErrCode means 401 unauthorized = 1005
	UnauthorizedErrCode
)

const (
	// ServerInternalErrorMsg is the ServerInternalErrorCode message
	ServerInternalErrorMsg = "Internal Server Error"
	// InvalidAttributeErrorMsg is the InvalidAttributeErrorCode message
	InvalidAttributeErrorMsg = "You passed an invalid value for the attributes."
	// RecordNotFoundErrorMsg is the RecordNotFoundErrorCode message
	RecordNotFoundErrorMsg = "Record Not Found"
	// UnauthorizedErrMsg is the UnauthorizedErrCode message
	UnauthorizedErrMsg = "Unauthorized"
)

// Error represents an error with an associated ExternalAPI status code.
type Error struct {
	InternalErr error      `json:"-"`
	Status      int        `json:"-"`
	GRPCStatus  codes.Code `json:"-"`
	OutputErr   string     `json:"error"`
}

// NewErr returns a Error instance.
func NewErr(code int, err error) *Error {
	e := &Error{
		InternalErr: err,
	}
	e.setup(code)
	return e
}

// Error allows handler Error struct to satisfy the build-in error interface.
func (e *Error) Error() string {
	return e.InternalErr.Error()
}

func (e *Error) setup(code int) {
	switch code {
	case SuccessCode:
		e.Status = http.StatusOK
		e.GRPCStatus = codes.OK
	case SuccessCreatedCode:
		e.Status = http.StatusCreated
		e.GRPCStatus = codes.OK
	case InvalidAttributeErrorCode:
		e.Status = http.StatusBadRequest
		e.GRPCStatus = codes.InvalidArgument
		e.OutputErr = InvalidAttributeErrorMsg
	case RecordNotFoundErrorCode:
		e.Status = http.StatusNotFound
		e.GRPCStatus = codes.NotFound
		e.OutputErr = RecordNotFoundErrorMsg
	case UnauthorizedErrCode:
		e.Status = http.StatusUnauthorized
		e.GRPCStatus = codes.Unauthenticated
		e.OutputErr = UnauthorizedErrMsg
	default:
		e.Status = http.StatusInternalServerError
		e.GRPCStatus = codes.Internal
		e.OutputErr = ServerInternalErrorMsg
	}
}
