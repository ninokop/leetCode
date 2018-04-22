package tree

func numTrees(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1

	for i := 2; i <= n; i++ {
		for k := 1; k <= i; k++ {
			f[i] += f[i-k] * f[k-1]
		}
	}
	return f[n]
}
