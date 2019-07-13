package main

import "fmt"

func main() {
	n := 64

	fmt.Printf("%d\n %b\n %X \n\n", n, n, n)

	i := 42

	fmt.Printf("%d\n %b\n %X \n\n", i, i, i)

	f := i << 1

	fmt.Printf("%d\n %b\n %X \n\n", f, f, f)

	a := `Mike said, "Hello"`

	fmt.Println(a)

}
