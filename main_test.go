package main

import (
	"fmt"
	"testing"
)

var seedNumbers = map[int]string{
	20:  "vingt",
	30:  "trente",
	40:  "quarante",
	50:  "cinquante",
	60:  "soixante",
	70:  "soixante-dix",
	80:  "quatre-vingts",
	81:  "quatre-vingt-un",
	82:  "quatre-vingt-deux",
	83:  "quatre-vingt-trois",
	84:  "quatre-vingt-quatre",
	85:  "quatre-vingt-cinq",
	86:  "quatre-vingt-six",
	87:  "quatre-vingt-sept",
	88:  "quatre-vingt-huit",
	89:  "quatre-vingt-neuf",
	90:  "quatre-vingt-dix",
	91:  "quatre-vingt-onze",
	92:  "quatre-vingt-douze",
	93:  "quatre-vingt-treize",
	94:  "quatre-vingt-quatorze",
	95:  "quatre-vingt-quinze",
	96:  "quatre-vingt-seize",
	97:  "quatre-vingt-dix-sept",
	98:  "quatre-vingt-dix-huit",
	99:  "quatre-vingt-dix-neuf",
	100: "cent",
}

func convertTestSeed(n int) string {
	tensDigit := n / 10
	if result, ok := seedNumbers[tensDigit*10]; ok {
		unitsDigit := n % 10
		if unitsDigit > 0 {
			return fmt.Sprintf("%s-%s", result, units(unitsDigit))
		}
		return result
	}
	return ""
}

func TestConvertZero(t *testing.T) {
	expected := "zéro"
	actual := convert(0)
	if actual != expected {
		t.Errorf("convert(0) expected %s, got %s", expected, actual)
	}
}

func TestConvertNegativeNumber(t *testing.T) {
	expected := "moins un"
	actual := convert(-1)
	if actual != expected {
		t.Errorf("convert(-1) expected %s, got %s", expected, actual)
	}
}

func TestConvertSingleDigit(t *testing.T) {
	for i := 1; i <= 9; i++ {
		expected := numbers[i]
		actual := convert(i)
		if actual != expected {
			t.Errorf("convert(%d) expected %s, got %s", i, expected, actual)
		}
	}
}

func TestConvertTens(t *testing.T) {
	for i := 20; i < 70; i++ {
		expected := convertTestSeed(i)
		actual := convert(i)
		if actual != expected {
			t.Errorf("convert(%d) expected %s, got %s", i, expected, actual)
		}
	}
}

func TestConvertSpecialSeventies(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{70, "soixante-dix"},
		{71, "soixante-et-onze"},
		{75, "soixante-quinze"},
		{79, "soixante-dix-neuf"},
	}
	for _, tc := range tests {
		actual := convert(tc.input)
		if actual != tc.output {
			t.Errorf("convert(%d) expected %s, got %s", tc.input, tc.output, actual)
		}
	}
}

func TestConvertEighties(t *testing.T) {
	for i := 80; i < 100; i++ {
		expected := seedNumbers[i]
		actual := convert(i)
		if actual != expected {
			t.Errorf("convert(%d) expected %s, got %s", i, expected, actual)
		}
	}
}

func TestConvertHundreds(t *testing.T) {
	numbers := map[int]string{
		101:  "cent-un",
		202:  "deux-cent-deux",
		303:  "trois-cent-trois",
		404:  "quatre-cent-quatre",
		505:  "cinq-cent-cinq",
		606:  "six-cent-six",
		707:  "sept-cent-sept",
		808:  "huit-cent-huit",
		909:  "neuf-cent-neuf",
		1000: "mille",
	}

	for k, v := range numbers {
		actual := convert(k)
		if actual != v {
			t.Errorf("convert(%d) expected %s, got %s", k, v, actual)
		}
	}
}

func TestConvertThousands(t *testing.T) {
	numbers := map[int]string{
		1000:   "mille",
		10000:  "dix-mille",
		50000:  "cinquante-mille",
		100000: "cent-mille",
		200000: "deux-cent-mille",
		300000: "trois-cent-mille",
		400000: "quatre-cent-mille",
		500000: "cinq-cent-mille",
		750000: "sept-cent-cinquante-mille",
		999999: "neuf-cent-quatre-vingt-dix-neuf-mille-neuf-cent-quatre-vingt-dix-neuf",
	}

	for k, v := range numbers {
		actual := convert(k)
		if actual != v {
			t.Errorf("convert(%d) expected %s, got %s", k, v, actual)
		}
	}
}
