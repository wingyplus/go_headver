package headver

import (
	"fmt"
	"time"
)

// Version is a headver version structure.
type Version struct {
	Head   int
	Year   int
	Week   int
	Build  int
	Suffix string
}

// String convert Version into string.
func (v Version) String() string {
	version := fmt.Sprintf("%d.%02d%02d.%d", v.Head, v.Year%100, v.Week%100, v.Build)
	if len(v.Suffix) != 0 {
		version = version + fmt.Sprintf("-%s", v.Suffix)
	}
	return version
}

// Build the headver version.
func Build(head, build int, suffix string) Version {
	return BuildWithTime(time.Now().UTC(), head, build, suffix)
}

// BuildWithTime similar to `Build` but it can specify time.
func BuildWithTime(t time.Time, head, build int, suffix string) Version {
	year, week := t.ISOWeek()
	return Version{Head: head, Year: year, Week: week, Build: build, Suffix: suffix}
}
