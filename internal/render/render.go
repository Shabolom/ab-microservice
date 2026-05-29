package render

import (
	"ab/pkg/shortcut"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorValidator(err error) error {
	switch {
	case errors.Is(err, shortcut.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, shortcut.ErrValidation),
		errors.Is(err, shortcut.ErrCheckViolation):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, shortcut.ErrDuplicateKey),
		errors.Is(err, shortcut.ErrExclusionViolation):
		return status.Error(codes.AlreadyExists, err.Error())

	case errors.Is(err, shortcut.ErrForeignKeyViolation),
		errors.Is(err, shortcut.ErrNotInExperiment):
		return status.Error(codes.FailedPrecondition, err.Error())

	case errors.Is(err, shortcut.ErrGroupNotFoundByBucket):
		return status.Error(codes.Internal, err.Error())

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
