package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numbersOfSet int
	fmt.Fscan(in, &numbersOfSet)

	var numbersOfSet2 int
	fmt.Fscan(in, &numbersOfSet2)

	counter := 0
	for i := 0; i < 1000; i++ {
		counter++
	}

	fmt.Fprintln(out, numbersOfSet2)

	inputAll(numbersOfSet)
}

func inputAll(numbersOfSet int) {

}

func testInput() {

}

func inputOne() {

}
