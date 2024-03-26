// This package allows one to build a 1-1 mapping between ranges of two integer
// domains e.g. [a:a+size] in domain 1 -> [x:x+size] in domain 2
// Note that this package does not handle all edge cases including adding
// multiple ranges with the same start value.
package rangedmap

import "slices"

type Range struct {
	start  int
	offset int
}

type RangedMap struct {
	ranges   []Range
	isSorted bool
}

func InitRangedMap() *RangedMap {
	return &RangedMap{
		ranges:   []Range{},
		isSorted: false,
	}
}

func (rM *RangedMap) AddRange(r Range) {
	rM.ranges = append(rM.ranges, r)
	rM.isSorted = false
}

// Sort ranges in ascending order by start value
func (rM *RangedMap) sortAsc() {
	slices.SortFunc(rM.ranges, func(a, b Range) int {
		return a.start - b.start
	})
	rM.isSorted = true
}

// Translate from one domain to the next using the RangedMap.
// This using lazy sorting. So we only sort at the moment we need to tranverse
// through the RangedMap Ranges to do translation.
func (rM *RangedMap) Translate(a int) int {
	if !rM.isSorted {
		rM.sortAsc()
	}
	for i, r := range rM.ranges {
		// If we are at the end of the list or we are between the current start
		// value and the next start value, then we return b in the translated
		// domain equal to a + offset
		if (i == len(rM.ranges)-1) ||
			((r.start <= a) && (a < rM.ranges[i+1].start)) {
			return a + r.offset
		}
	}
	return a
}
