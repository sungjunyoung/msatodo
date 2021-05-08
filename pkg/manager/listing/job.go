package listing

import "time"

type Job struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	DoneAt    time.Time
}
