func firstMissingPositive(nums []int) int {
    m := make(map[int]bool)

    for _, elm := range nums {
        m[elm] = true
    }

    for i := 1; ; i++ {
        if !m[i] {
            return i
        }
    }
}
