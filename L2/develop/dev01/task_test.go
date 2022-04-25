package main

import (
	"develop/current_time"
	"testing"
	"time"
)

func TestGetCurrTime(t *testing.T) {
	now := time.Now()

	currTime, err := current_time.GetCurrTime()
	if err != nil {
		t.Error(err)
	}
	if currTime.Before(now) || currTime.After(now.Add(6*time.Second)) {
		t.Errorf("%v !≈ %v", now, currTime)
	}
}
