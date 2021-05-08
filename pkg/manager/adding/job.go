package adding

import "time"

type Job struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	DoneAt    time.Time
}
