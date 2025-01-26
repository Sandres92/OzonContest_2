package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

type ArrayInput struct {
	name  string
	price string
}

type ArrayOutput struct {
	name  string
	price string
}

func main() {

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numbersOfSet int
	fmt.Fscan(in, &numbersOfSet)
	in.ReadString('\n')

	inputAll(numbersOfSet)

	//str := "abcdef"
	//str = strings.TrimSpace(str)
	//arr := strings.Split(str, ",")
	//fmt.Fprintln(out, len(arr))
	//
	//pattern := `^([a-z]{1,10}:\d+)(?:,([a-z]{1,10}:\d+))*$`
	////pattern := "^([a-zA-Z0-9]+):([0-9]+)"
	//re := regexp.MustCompile(pattern)
	//
	//if re.MatchString(str) {
	//	fmt.Fprintln(out, "yes")
	//} else {
	//	fmt.Fprintln(out, "no")
	//}
}

func inputAll(numbersOfSet int) {
	result := make([]bool, numbersOfSet)
	resultName := make([]string, numbersOfSet)
	for i := 0; i < numbersOfSet; i++ {
		resultName[i], result[i] = inputOne()
	}

	for i := 0; i < numbersOfSet; i++ {
		if result[i] {
			//fmt.Fprintf(out, "%d ||%s  || %s\n", i+1, resultName[i], "YES")
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func inputOne() (string, bool) {
	var countRow int
	fmt.Fscanf(in, "%d\n", &countRow)

	arrayInput := make([]ArrayInput, countRow)

	priceKey := make(map[string]string)
	nameKey := make(map[string]string)

	priceKeyOut := make(map[string]string)

	for i := 0; i < countRow; i++ {
		arrayStr := readStringArray(" ")

		arrayInput[i].name = arrayStr[0]
		arrayInput[i].price = arrayStr[1]

		priceKey[arrayStr[1]] = arrayStr[0]
		nameKey[arrayStr[0]] = arrayStr[1]
	}

	resultStr := readLine()
	pattern := `^([a-z]{1,10}:\d+)(?:,([a-z]{1,10}:\d+))*$`
	//regexp.MustCompile(`^([a-z]{1,10}:\d+)(?:,([a-z]{1,10}:\d+))*$`)
	//pattern := "^([a-zA-Z0-9]+):([0-9]+)"
	re := regexp.MustCompile(pattern)

	if !re.MatchString(resultStr) {
		//fmt.Fprintf(out, "%s , %s \n", arrayInput[0].name, "NO")

		return arrayInput[0].name, false
	} else {
		resultStrArr := strings.Split(resultStr, ",")
		arrayOutput := make([]ArrayOutput, len(resultStrArr))

		isCorrect := true
		for i := 0; i < len(resultStrArr); i++ {
			subResultStrArr := strings.Split(resultStrArr[i], ":")
			arrayOutput[i].name = subResultStrArr[0]
			arrayOutput[i].price = subResultStrArr[1]
			_, ok := priceKeyOut[arrayOutput[i].price]
			if !ok {
				priceKeyOut[arrayOutput[i].price] = arrayOutput[i].name
			} else {
				isCorrect = false
				break
			}

			_, ok = nameKey[arrayOutput[i].name]
			if !ok {
				isCorrect = false
				break
			}
		}

		if isCorrect {
			//fmt.Fprintf(out, "===============\n")
			for k, _ := range priceKey {
				_, ok := priceKeyOut[k]

				//fmt.Fprintf(out, "%s , %t \n", k, ok)

				if !ok {
					isCorrect = false
					break
				}
			}
			//fmt.Fprintf(out, "===============\n")
			//if !isNo {
			//	//fmt.Fprintf(out, "YES")
			//	//fmt.Fprintf(out, "%s , %s \n", arrayInput[0].name, "YES")
			//} else {
			//	//fmt.Fprintf(out, "NO")
			//	//fmt.Fprintf(out, "%s , %s \n", arrayInput[0].name, "NO")
			//}
		}

		return arrayInput[0].name, isCorrect
	}
}

func checkPrices(arrayInput []ArrayInput, arrayOutput []ArrayOutput) bool {

	sort.Slice(arrayInput, func(i, j int) bool {
		return arrayInput[i].price < arrayInput[j].price
	})

	sort.Slice(arrayOutput, func(i, j int) bool {
		return arrayOutput[i].price < arrayOutput[j].price
	})

	indIn := 0
	priceIndOut := 0
	nameIndOut := 0

	checkNames := make([]bool, len(arrayOutput))

	if arrayOutput[indIn].price != arrayInput[priceIndOut].price {
		return false
	}

	if arrayOutput[indIn].name != arrayInput[nameIndOut].name {
		checkNames[nameIndOut] = true
		nameIndOut++
	}

	for i := 1; i < len(arrayInput); i++ {
		if arrayInput[indIn].price == arrayInput[i].price {

			if arrayInput[i].name == arrayOutput[nameIndOut].name {
				checkNames[nameIndOut] = true
				nameIndOut++
			}

			continue
		}

		for j := priceIndOut; j < len(arrayOutput); j++ {
			if arrayOutput[priceIndOut].price == arrayOutput[j].price {
				priceIndOut++
				continue
			}

			if arrayInput[indIn].price != arrayOutput[priceIndOut].price {
				return false
			}
		}
	}
	return true
}

func checkNames(arrayInput []ArrayInput, arrayOutput []ArrayOutput) bool {
	sort.Slice(arrayInput, func(i, j int) bool {
		return arrayInput[i].price < arrayInput[j].price
	})

	sort.Slice(arrayOutput, func(i, j int) bool {
		return arrayOutput[i].price < arrayOutput[j].price
	})

	indIn := 0
	indOut := 0

	if arrayOutput[indIn].price != arrayInput[indOut].price {
		return false
	}

	for i := 1; i < len(arrayInput); i++ {
		if arrayInput[indIn].price == arrayInput[i].price {
			continue
		}

		for j := indOut; j < len(arrayOutput); j++ {
			if arrayOutput[indOut].price == arrayOutput[j].price {
				indOut++
				continue
			}

			if arrayInput[indIn].price != arrayOutput[indOut].price {
				return false
			}
		}
	}
	return true
}

func testInput() {

}

func readNumber() (int, error) {
	line, _ := in.ReadString('\n')
	return strconv.Atoi(strings.TrimSpace(line))
}

func readLine() string {
	line, _ := in.ReadString('\n')
	return strings.TrimSpace(line)
}

func readStringArray(split string) []string {
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	return strings.Split(line, split)
}

func readNumberArray() []int {
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	lineByElem := strings.Split(line, " ")

	numbers := make([]int, len(lineByElem))

	for i := 0; i < len(lineByElem); i++ {
		numbers[i], _ = strconv.Atoi(lineByElem[i])
	}

	return numbers
}
