package gotil

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	exp := time.Now().UnixNano() / 1000000
	got := GetCurrentTime()
	if got != exp {
		t.Errorf("GetCurrentTime() expected %q but got %q", exp, got)
	}
}
