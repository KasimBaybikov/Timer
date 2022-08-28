package main

import "time"

type Timer struct {
	start   time.Time
	started bool
}

func (t *Timer) Start() {
	t.started = true
}

func (t *Timer) Stop() {
	t.started = false
}

func (t Timer) Running() bool {
	return t.started
}

func newTimer() Timer {
	return Timer{}
}
