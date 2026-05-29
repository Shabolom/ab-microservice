package shortcut

import (
	"errors"
	"log"
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
)

func FatalIfErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
