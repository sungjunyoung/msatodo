package cache

import (
	"sort"
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

func (m *memory) ListJobs() []Job {
	jobs := []Job{}

	for _, job := range m.jobs {
		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].CreatedAt.After(jobs[j].CreatedAt)
	})

	return jobs
}
