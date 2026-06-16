package metrics

const (
	ExperimentTypeWithGroups    = "with_groups"
	ExperimentTypeWithoutGroups = "without_groups"
)

func (m *Metrics) SetActiveExperimentsWithGroup(count int) {
	m.activeExperiments.
		WithLabelValues(ExperimentTypeWithGroups).
		Set(float64(count))
}

func (m *Metrics) SetActiveExperimentsWithoutGroup(count int) {
	m.activeExperiments.
		WithLabelValues(ExperimentTypeWithoutGroups).
		Set(float64(count))
}
