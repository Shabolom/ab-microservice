package di

import inMemmoryCashe "ab/internal/in-memory-cache"

func (d *DI) GetInMemoryCache() *inMemmoryCashe.RawExperimentSessionStorage {
	if d.inMemoryCache != nil {
		return d.inMemoryCache
	}

	d.inMemoryCache = inMemmoryCashe.New()
	return d.inMemoryCache
}
