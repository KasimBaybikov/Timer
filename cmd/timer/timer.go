package main

import (
	"time"
)

type Timer struct {
	start           time.Time
	started         bool
	historyFilePath string
}

func (t *Timer) SetHistoryFilePath(historyFilePath string) {
	t.historyFilePath = historyFilePath
}

func (t *Timer) Start() {
	t.started = true
}

func (t *Timer) Stop() {
	t.started = false
}

func (t Timer) IsRunning() bool {
	return t.started
}

func newTimer() Timer {
	return Timer{}
}
