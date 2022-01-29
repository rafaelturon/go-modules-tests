package arithmetic

func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func CubeVolume(n int) int, errpr {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("n is zero")
	}
}
