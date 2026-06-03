package shortcut

import (
	"errors"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrValidation          = errors.New("validation error")
	ErrDuplicateKey        = errors.New("duplicate key")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckViolation      = errors.New("check violation")
	ErrExclusionViolation  = errors.New("exclusion violation")

	ErrNotInExperiment       = errors.New("device is not in experiment")
	ErrGroupNotFoundByBucket = errors.New("group not found by bucket")
	ErrFailedToUpdateGroup   = errors.New("failed to update group")
	ErrFailedToDeleteGroup   = errors.New("failed to delete group")
	ErrFailedToGetGroup      = errors.New("failed to get group")
	ErrFailedToScanGroup     = errors.New("failed to scan group")
	ErrRowsIteration         = errors.New("rows iteration error")

	ErrExperimentEndDateInPast       = errors.New("experiment end date is in past")
	ErrExperimentStartDateAfterEnd   = errors.New("experiment start date must be before end date")
	ErrExperimentNameRequired        = errors.New("experiment name is required")
	ErrExperimentStatusRequired      = errors.New("experiment status is required")
	ErrExperimentGroupsMinCount      = errors.New("experiment must have at least 2 groups")
	ErrExperimentLayersRequired      = errors.New("experiment must have at least 1 layer")
	ErrExperimentRolloutOutOfRange   = errors.New("experiment rollout percentage must be between 0 and 100")
	ErrExperimentGroupsRolloutTooBig = errors.New("experiment groups rollout percentage must be <= 100")
)

const (
	ExpStatusInTest  = "in test"
	ExpStatusReady   = "ready"
	ExpStatusActive  = "active"
	ExpStatusStopped = "stopped"
	ExpStatusEnded   = "ended"
)
