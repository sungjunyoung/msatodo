package adding

import (
	"github.com/hashicorp/go-uuid"
	"github.com/sirupsen/logrus"
	"github.com/sungjunyoung/msatodo/pkg/manager/cache"
	"time"
)

type Cache interface {
	AddJob(cache.Job)
}

type Service interface {
	AddJob(Job) (string, error)
}

type service struct {
	cache Cache
}

func NewService(c Cache) *service {
	return &service{
		cache: c,
	}
}

func (s *service) AddJob(j Job) (string, error) {
	id, err := uuid.GenerateUUID()
	if err != nil {
		logrus.Errorf("cannot generate uuid for job: %+v", err)
		return "", err
	}
	createdAt := time.Now()

	job := cache.Job{
		ID:        id,
		Name:      j.Name,
		Email:     j.Email,
		CreatedAt: createdAt,
		DoneAt:    j.DoneAt,
	}

	s.cache.AddJob(job)
	return id, nil
}
