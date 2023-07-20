package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcError(err error) error {
	appErr := &Error{}
	if errors.As(err, &appErr) {
		return status.Error(errGrpcCodeMap[appErr.Code()], appErr.Error())
	}
	return status.Error(codes.Internal, err.Error())
}

func HandleServiceError(err error) error {
	switch err {
	case ErrValidationError:
		return NewError(ErrValidationError.Error(), ErrCodeInvalidArgument)
	default:
		return NewError(ErrUnknown.Error(), ErrCodeUnknown)
	}
}
