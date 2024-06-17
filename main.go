package main

import (
	"encoding/json"
	"fmt"
)

var numbers = map[int]string{
	0: "zéro", 1: "un", 2: "deux", 3: "trois", 4: "quatre", 5: "cinq",
	6: "six", 7: "sept", 8: "huit", 9: "neuf", 10: "dix", 11: "onze", 12: "douze",
	13: "treize", 14: "quatorze", 15: "quinze", 16: "seize",
}

func units(n int) string {
	return numbers[n]
}

type numberState interface {
	convert(n int) string
}

type unitsState struct{}

func (u unitsState) convert(n int) string {
	if n > 16 {
		return fmt.Sprintf("dix-%s", units(n-10))
	}
	return numbers[n]
}

type tensState struct{}

func (t tensState) convert(n int) string {
	tens := []string{"dix", "vingt", "trente", "quarante", "cinquante", "soixante"}
	// off by one tens slice index
	d := n/10 - 1
	isDivisibleByTen := n % 10
	if isDivisibleByTen == 0 {
		return tens[d]
	}
	return fmt.Sprintf("%s-%s", tens[d], units(n%10))
}

type specialSeventiesState struct{}

func (s specialSeventiesState) convert(n int) string {
	if n == 70 {
		return "soixante-dix"
	} else if n == 71 {
		return "soixante-et-onze"
	}
	return fmt.Sprintf("soixante-%s", convert(n%60))
}

type eightiesState struct{}

func (e eightiesState) convert(n int) string {
	if n == 80 {
		return "quatre-vingts"
	}
	return fmt.Sprintf("quatre-vingt-%s", convert(n%80))
}

type hundredsState struct{}

func (h hundredsState) convert(n int) string {
	hundred := "cent"
	if n == 100 {
		return hundred
	}

	factor := n / 100
	if factor > 1 {
		hundred = fmt.Sprintf("%s-cent", convert(factor))
	}

	modHundred := n % 100
	if modHundred > 0 {
		return fmt.Sprintf("%s-%s", hundred, convert(modHundred))
	}

	return hundred
}

type thousandsState struct{}

func (t thousandsState) convert(n int) string {
	thousand := "mille"
	if n == 1000 {
		return thousand
	}

	factor := n / 1000
	if factor > 1 {
		thousand = fmt.Sprintf("%s-mille", convert(factor))
	}

	modThousand := n % 1000
	if modThousand > 0 {
		return fmt.Sprintf("%s-%s", thousand, convert(modThousand))
	}

	return thousand
}

func convert(n int) string {
	if n == 0 {
		return numbers[n]
	} else if n < 0 {
		return fmt.Sprintf("moins %s", convert(-n))
	}

	var state numberState
	if n < 20 {
		state = unitsState{}
	} else if n < 70 {
		state = tensState{}
	} else if n < 80 {
		state = specialSeventiesState{}
	} else if n < 100 {
		state = eightiesState{}
	} else if n < 1000 {
		state = hundredsState{}
	} else {
		state = thousandsState{}
	}

	return state.convert(n)
}

func main() {
	input := []int{0, 1, 5, 10, 11, 15, 20, 21, 30, 35, 50, 51, 68, 70, 75, 99, 100, 101, 105, 111, 123, 168, 171, 175, 199, 200, 201, 555, 999, 1000, 1001, 1111, 1199, 1234, 1999, 2000, 2001, 2020, 2021, 2345, 9999, 10000, 11111, 12345, 123456, 654321, 999999}
	output := make([]string, len(input))
	for i, num := range input {
		output[i] = convert(num)
		fmt.Printf("num %d: %s \n", num, output[i])
	}
	b, _ := json.Marshal(output)
	fmt.Printf("%v", string(b))
}
