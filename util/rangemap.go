package util

import "sort"

// https://gist.github.com/d-schmidt/cd85136990de250df38956432517b253

type Range struct {
	L int
	U int
}

type RangeMap struct {
	Keys   []Range
	Values []int
}

func (rm *RangeMap) Set(keyL, keyR, value int) {
	rm.Keys = append(rm.Keys, Range{keyL, keyR})
	rm.Values = append(rm.Values, value)
}

func (rm *RangeMap) Get(key int) (int, int, int, bool) {
	i := sort.Search(len(rm.Keys), func(i int) bool {
		return key < rm.Keys[i].L
	})

	i -= 1
	if i >= 0 && i < len(rm.Keys) && key <= rm.Keys[i].U {
		return rm.Values[i], rm.Keys[i].L, rm.Keys[i].U, true
	}

	return -1, -1, -1, false
}
