package headver

import (
	"testing"
	"time"
)

func TestBuildWithTime(t *testing.T) {
	date := time.Date(2023, time.June, 8, 0, 0, 0, 0, time.UTC)
	match(t, BuildWithTime(date, 1, 0, ""), "1.2323.0")
	match(t, BuildWithTime(date, 1, 0, "patch"), "1.2323.0-patch")
}

func match(t *testing.T, v Version, expected string) {
	t.Helper()
	if v.String() != expected {
		t.Errorf("Expected %s, got %s", expected, v.String())
	}
}
