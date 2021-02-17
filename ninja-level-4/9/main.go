package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`fernando_shan`] = []string{`stirred not shaken`, `shaken not shaken`, `not stirred not`}

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}
}
