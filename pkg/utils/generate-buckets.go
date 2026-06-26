package utils

import (
	"math/rand"
	"slices"
)

func GenerateBuckets(usedBuckets []int64, count int64) []int64 {
	used := make(map[int64]struct{}, len(usedBuckets))

	for _, bucket := range usedBuckets {
		used[bucket] = struct{}{}
	}

	freeBuckets := make([]int64, 0, 100-len(used))

	for bucket := int64(1); bucket <= 100; bucket++ {
		if _, ok := used[bucket]; !ok {
			freeBuckets = append(freeBuckets, bucket)
		}
	}

	rand.Shuffle(len(freeBuckets), func(i, j int) {
		freeBuckets[i], freeBuckets[j] = freeBuckets[j], freeBuckets[i]
	})

	if count > int64(len(freeBuckets)) {
		count = int64(len(freeBuckets))
	}

	result := freeBuckets[:count]

	slices.Sort(result)

	return result
}
