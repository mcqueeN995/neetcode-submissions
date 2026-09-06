import uni "unicode"

func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    
    for l < r {
        for l < r && !uni.IsLetter(rune(s[l])) && !uni.IsDigit(rune(s[l])) {
            l++
        }
        
        for l < r && !uni.IsLetter(rune(s[r])) && !uni.IsDigit(rune(s[r])) {
            r--
        }
        
        if uni.ToLower(rune(s[l])) != uni.ToLower(rune(s[r])) {
            return false
        }
        
        l++
        r--
    }
    
    return true
}