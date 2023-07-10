package main

func st() *int {
	a := 1
	return &a
}

func main() {
	var _ *int = st()

}
