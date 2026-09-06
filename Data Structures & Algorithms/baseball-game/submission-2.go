func calPoints(operations []string) int {
    stack := []int{}

    for i := 0; i < len(operations); i++ {
        op := operations[i]

        if val, err := strconv.Atoi(op); err == nil {
            stack = append(stack, val)
        } else {
            switch op {
            case "+":
                sum := stack[len(stack)-1] + stack[len(stack)-2]
                stack = append(stack, sum)
            case "D":
                stack = append(stack, stack[len(stack)-1]*2)
            case "C":
                stack = stack[:len(stack)-1]
            }
        }
    }

    result := 0
    for _, val := range stack {
        result += val
    }
    return result
}