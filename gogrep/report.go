package gogrep

import "time"

type Report struct {
	id int
	Elapsed time.Duration
	Remaining time.Duration
	ByteCnt int
	Status  string
}
