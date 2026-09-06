import "slices"

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	boats := 0
	l, r := 0, len(people) - 1
	for l <= r {
		if people[l] + people[r] <= limit {
			l++
			r--
		} else {
			r--
		}
		boats++
	}


	return boats
}
