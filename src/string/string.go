package string

func Reverse(s string) string {
    b := []byte(s)
    for i := 0; i< len(b)/ 2; i++ {
        j := len(b) - i - i
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}
