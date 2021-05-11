package listing

import (
	"github.com/sungjunyoung/msatodo/pkg/manager/cache"
)

type Cache interface {
	ListJobs() []cache.Job
}

type Service interface {
	ListJobs() ([]Job, error)
}

type service struct {
	cache Cache
}

func NewService(c Cache) *service {
	return &service{
		cache: c,
	}
}

func (s *service) ListJobs() ([]Job, error) {
	jobs := []Job{}
	cacheJobs := s.cache.ListJobs()
	for _, cacheJob := range cacheJobs {
		jobs = append(jobs, Job{
			ID:        cacheJob.ID,
			Name:      cacheJob.Name,
			Email:     cacheJob.Email,
			CreatedAt: cacheJob.CreatedAt,
			DoneAt:    cacheJob.DoneAt,
		})
	}

	return jobs, nil
}
