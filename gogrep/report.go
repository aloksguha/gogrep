package gogrep

import "time"

type STATUS string
const (
	SUCEESS STATUS = "SUCCESS"
	FAILURE STATUS = "FAILURE"
	TIMEOUT STATUS = "TIMEOUT"
)

type Report struct {
	id int
	Elapsed time.Duration
	Remaining time.Duration
	ByteCnt int
	Status  STATUS
}
