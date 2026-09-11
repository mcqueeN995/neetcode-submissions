func firstMissingPositive(nums []int) int {
    m := make(map[int]bool)

    for _, elm := range nums {
        m[elm] = true
    }

    // Проверяем все положительные числа начиная с 1
    for i := 1; ; i++ {
        if !m[i] {
            return i
        }
    }
}