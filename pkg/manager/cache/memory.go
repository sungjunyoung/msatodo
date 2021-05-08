package cache

import (
	"sync"
)

var memoryCache *memory
var once sync.Once

type memory struct {
	jobs map[string]Job
}

func NewMemory() *memory {
	once.Do(func() {
		memoryCache = &memory{
			jobs: map[string]Job{},
		}
	})

	return memoryCache
}

func (m *memory) AddJob(job Job) {
	m.jobs[job.ID] = job
}
