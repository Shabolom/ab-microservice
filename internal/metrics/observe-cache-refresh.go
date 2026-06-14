package metrics

import "time"

func (m *Metrics) ObserveCacheRefresh(startedAt time.Time) {
	m.cacheRefreshTime.
		Observe(time.Since(startedAt).Seconds())
}
