package common

import "github.com/go-co-op/gocron/v2"

type Inject struct {
	V         *Values
	Scheduler gocron.Scheduler
}
