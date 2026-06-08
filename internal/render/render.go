package render

import (
	"ab/pkg/shortcut"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorValidator(err error) error {
	switch {
	case errors.Is(err, shortcut.ErrNotFound),
		errors.Is(err, shortcut.ErrLayerNotFound),
		errors.Is(err, shortcut.ErrFailedToGetExperiment):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, shortcut.ErrValidation),
		errors.Is(err, shortcut.ErrCheckViolation),
		errors.Is(err, shortcut.ErrExperimentEndDateInPast),
		errors.Is(err, shortcut.ErrExperimentStartDateAfterEnd),
		errors.Is(err, shortcut.ErrExperimentNameRequired),
		errors.Is(err, shortcut.ErrExperimentStatusRequired),
		errors.Is(err, shortcut.ErrExperimentGroupsMinCount),
		errors.Is(err, shortcut.ErrExperimentLayersRequired),
		errors.Is(err, shortcut.ErrExperimentRolloutOutOfRange),
		errors.Is(err, shortcut.ErrExperimentGroupsRolloutTooBig),
		errors.Is(err, shortcut.ErrExperimentAlreadyRunning),
		errors.Is(err, shortcut.ErrDifferentNamespaces),
		errors.Is(err, shortcut.ErrExperimentGroupsRolloutNotFull),
		errors.Is(err, shortcut.ErrNoParamsInGroup),
		errors.Is(err, shortcut.ErrNameSpaseCountInGroup),
		errors.Is(err, shortcut.ErrNamespaceIDsDontMatch),
		errors.Is(err, shortcut.ErrCustomParamsNamespaceMismatch),
		errors.Is(err, shortcut.ErrNotSupportedConditions),
		errors.Is(err, shortcut.ErrNotSupportedType),
		errors.Is(err, shortcut.ErrExperimentLayerRolloutTooBig):
		return status.Error(codes.InvalidArgument, err.Error())

	case errors.Is(err, shortcut.ErrDuplicateKey),
		errors.Is(err, shortcut.ErrExclusionViolation):
		return status.Error(codes.AlreadyExists, err.Error())

	case errors.Is(err, shortcut.ErrForeignKeyViolation),
		errors.Is(err, shortcut.ErrNotInExperiment):
		return status.Error(codes.FailedPrecondition, err.Error())

	case errors.Is(err, shortcut.ErrTimeout):
		return status.Error(codes.DeadlineExceeded, err.Error())

	case errors.Is(err, shortcut.ErrCanceled):
		return status.Error(codes.Canceled, err.Error())

	case errors.Is(err, shortcut.ErrRetryable):
		return status.Error(codes.Aborted, err.Error())

	case errors.Is(err, shortcut.ErrQueryCanceled):
		return status.Error(codes.Canceled, err.Error())

	case errors.Is(err, shortcut.ErrFailedToGetLayerExperiments),
		errors.Is(err, shortcut.ErrFailedToUpdateExperiment),
		errors.Is(err, shortcut.ErrFailedToGetLayer),
		errors.Is(err, shortcut.ErrStorage),
		errors.Is(err, shortcut.ErrFailedToUpdateGroup),
		errors.Is(err, shortcut.ErrFailedToDeleteGroup),
		errors.Is(err, shortcut.ErrFailedToGetGroup),
		errors.Is(err, shortcut.ErrFailedToScanGroup),
		errors.Is(err, shortcut.ErrRowsIteration),
		errors.Is(err, shortcut.ErrGroupNotFoundByBucket):
		return status.Error(codes.Internal, err.Error())

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
