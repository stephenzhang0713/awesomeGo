package gomonkey

func networkCompute(a, b int) (int, error) {
	// do something in remote computer
	c := a + b

	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := networkCompute(a, b)
	return sum, err
}
