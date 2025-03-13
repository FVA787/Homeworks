package package2

func Sum(a *int) int {
	*a += 3
	return *a
}
