package main

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := a[0:10]
	println(len(x))
	println(cap(x))
}
