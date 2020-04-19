package efficientgo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type person struct {
	name, profession string
	year             uint32
}

func printAsValue() {
	p := person{"Martin", "Senior Backend Go Developer", 1980}

	fmt.Printf("%v\n", p)
}

func printAsIndividualFields() {
	p := person{"Martin", "Senior Backend Go Developer", 1980}

	fmt.Printf("%s, %s, %d\n", p.name, p.profession, p.year)
}

func printByteSlice() {
	p := person{"Martin", "Senior Backend Go Developer", 1980}

	w := bufio.NewWriter(os.Stdout)
	if _, err := w.WriteString(p.name + ", " + p.profession + ", " + strconv.Itoa(int(p.year))); err != nil {
		panic(err)
	}
}
