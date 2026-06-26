package dto

type NamespaceFeatureToggle struct {
	Namespace      string
	FeatureToggles []RawFeatureToggle
}
