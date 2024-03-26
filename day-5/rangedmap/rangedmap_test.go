package rangedmap

import (
	"testing"
)

func TestRangedMap(t *testing.T) {
	rM := InitRangedMap()
	rM.AddRange(Range{
		start:  5,
		offset: 10,
	})
	rM.AddRange(Range{
		start:  3,
		offset: -1,
	})
	if rM.ranges[0].start != 5 {
		t.Error("First range should be at start: 5")
	}

	b := rM.Translate(3)
	if b != 2 {
		t.Error("Translated value should be 2")
	}
	if rM.ranges[0].start != 3 {
		t.Error("First range should be at start: 3 after sorting")
	}

	b = rM.Translate(16)
	if b != 26 {
		t.Error("Translated value should be 26")
	}
	t.Logf("%#v\n", rM)
}
