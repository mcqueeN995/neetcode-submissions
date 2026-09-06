func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        if num, err := strconv.Atoi(token); err == nil {
            stack = append(stack, num)
        } else {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2] 

            var result int
            switch token {
            case "+":
                result = a + b
            case "-":
                result = a - b
            case "*":
                result = a * b
            case "/":
                result = a / b 
            }
            stack = append(stack, result)
        }
    }
    return stack[0] 
}