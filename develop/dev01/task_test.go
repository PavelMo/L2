package main

import (
	"testing"
	"time"
)

func TestGetCurrTime(t *testing.T) {
	now := time.Now()

	currTime, err := getCurrTime()
	if err != nil {
		t.Error(err)
	}
	if currTime.Before(now) || currTime.After(now.Add(6*time.Second)) {
		t.Errorf("%v !≈ %v", now, currTime)
	}
}
