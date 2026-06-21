package shortcut

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrValidation          = errors.New("validation error")
	ErrDuplicateKey        = errors.New("duplicate key")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckViolation      = errors.New("check violation")
	ErrExclusionViolation  = errors.New("exclusion violation")

	ErrTimeout       = errors.New("operation timeout")
	ErrCanceled      = errors.New("operation canceled")
	ErrRetryable     = errors.New("retryable database error")
	ErrQueryCanceled = errors.New("query canceled")

	ErrNotInExperiment       = errors.New("device is not in experiment")
	ErrGroupNotFoundByBucket = errors.New("group not found by bucket")
	ErrFailedToUpdateGroup   = errors.New("failed to update group")
	ErrFailedToDeleteGroup   = errors.New("failed to delete group")
	ErrFailedToGetGroup      = errors.New("failed to get group")
	ErrFailedToScanGroup     = errors.New("failed to scan group")
	ErrRowsIteration         = errors.New("rows iteration error")

	ErrExperimentEndDateInPast        = errors.New("experiment end date is in past")
	ErrExperimentStartDateAfterEnd    = errors.New("experiment start date must be before end date")
	ErrExperimentNameRequired         = errors.New("experiment name is required")
	ErrExperimentStatusRequired       = errors.New("experiment status is required")
	ErrExperimentGroupsMinCount       = errors.New("experiment must have at least 2 groups")
	ErrExperimentLayersRequired       = errors.New("experiment must have at least 1 layer")
	ErrExperimentRolloutOutOfRange    = errors.New("experiment rollout percentage must be between 0 and 100")
	ErrExperimentGroupsRolloutTooBig  = errors.New("experiment groups rollout percentage must be <= 100")
	ErrExperimentGroupsRolloutNotFull = errors.New("experiment groups rollout percentage must be == 100")

	ErrFailedToUpdateExperiment     = errors.New("failed to update experiment")
	ErrExperimentLayerRolloutTooBig = errors.New("experiment layer rollout percentage must be <= 100")
	ErrFailedToGetExperiment        = errors.New("failed to get experiment")
	ErrFailedToGetLayerExperiments  = errors.New("failed to get layer experiments")

	ErrLayerNotFound                 = errors.New("layer not found")
	ErrFailedToGetLayer              = errors.New("failed to get layer")
	ErrStorage                       = errors.New("storage error")
	ErrExperimentAlreadyRunning      = errors.New("experiment status is already ready or active")
	ErrDifferentNamespaces           = errors.New("different namespaces")
	ErrTypeCast                      = errors.New("error while trying to cast to a specific type")
	ErrNoParamsInGroup               = errors.New("no params in group")
	ErrNameSpaseCountInGroup         = errors.New("name spase count in group")
	ErrNamespaceIDsDontMatch         = errors.New("namespace IDs don't match")
	ErrCustomParamsNamespaceMismatch = errors.New("custom params namespace does not match experiment namespace")

	ErrNotSupportedConditions = errors.New("not supported conditions")
	ErrNotSupportedType       = errors.New("not supported type")

	ErrFeatureToggleNameRequired        = errors.New("feature toggle name required")
	ErrFeatureToggleNamespaceIDRequired = errors.New("feature toggle namespace id required")
	ErrFeatureToggleRolloutOutOfRange   = errors.New("feature toggle rollout percentage out of range")
	ErrFeatureToggleInvalidStatus       = errors.New("feature toggle status is invalid")
)

const (
	ExpStatusInTest  = "in test"
	ExpStatusReady   = "ready"
	ExpStatusActive  = "active"
	ExpStatusStopped = "stopped"
	ExpStatusEnded   = "ended"
)

func MapStorageError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return ErrNotFound

	case errors.Is(err, context.DeadlineExceeded):
		return ErrTimeout

	case errors.Is(err, context.Canceled):
		return ErrCanceled

	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ErrDuplicateKey

		case "23503":
			return ErrForeignKeyViolation

		case "23514":
			return ErrCheckViolation

		case "23P01":
			return ErrExclusionViolation

		case "23502",
			"22001",
			"22P02",
			"22003",
			"22007",
			"22008",
			"22023":
			return ErrValidation

		case "42P01",
			"42703",
			"42883",
			"42P10",
			"42601":
			return ErrStorage

		case "40001",
			"40P01",
			"53300",
			"08000",
			"08003",
			"08006": // connection_failure
			return ErrRetryable

		case "57014":
			return ErrQueryCanceled
		}
	}

	return ErrStorage
}
