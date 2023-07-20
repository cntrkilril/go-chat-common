package errors

import "google.golang.org/grpc/codes"

var errGrpcCodeMap = map[ErrCode]codes.Code{
	ErrCodeNotFound:        codes.NotFound,
	ErrCodeUnknown:         codes.Unknown,
	ErrCodeInvalidArgument: codes.InvalidArgument,
}
