func climbStairs(n int) int {
      pprev := 1
    prev := 1
    if n == 0{
        return 0
    }
    for i := 0 ; i< n-1; i++{
        sum := pprev + prev
        pprev = prev
        prev = sum
    }
    return prev
}