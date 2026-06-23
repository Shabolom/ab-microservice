package dto

type UpdateFeatureToggleBuckets struct {
	FeatureToggleID int64

	ExceptionBuckets []int64
	IOSBuckets       []int64
	AndroidBuckets   []int64
	WebBuckets       []int64
}
