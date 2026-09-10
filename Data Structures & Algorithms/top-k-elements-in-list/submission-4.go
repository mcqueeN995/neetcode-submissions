import (
	"cmp"
	"slices"
)

type Pair struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, elm := range nums {
		m[elm]++
	}

	arr := make([]Pair, 0, len(m))
	for key, val := range m {
		p := Pair{key, val}
		arr = append(arr, p)
	}

	slices.SortFunc(arr, func(a, b Pair) int {
		return cmp.Compare(b.value, a.value)
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, arr[i].key)
	}
	return res
}