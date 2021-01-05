package calculation

// function global, bisa diakses di luar package
func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

// function local, hanya bisa diakses dalam package sama
func add(number int, numberTwo int) int {
	return number + numberTwo
}