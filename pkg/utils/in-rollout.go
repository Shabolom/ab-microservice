package utils

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

func IsUserInRolloutPercentage(id, splitID int64, buckets []int64) bool {
	key := fmt.Sprintf("experiment:%d:SplitID:%v", id, splitID)

	bucket := getBucket(key)

	for _, expBucket := range buckets {
		if bucket == expBucket {
			return true
		}
	}

	return false
}

func getBucket(key string) int64 {
	hash := murmur3.Sum32([]byte(key))
	bucket := int64(hash%100) + 1

	return bucket
}

func IsUserInRolloutPercentageWithoutBuckets(id, splitID, percentage int64) bool {
	hash := fmt.Sprintf("experiment:%d:SplitID:%v", id, splitID)

	key := getBucket(hash)

	if key < percentage {
		return true
	}

	return false
}
