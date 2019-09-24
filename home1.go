package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Flag struct {
	NoRegister   bool
	NoRepetition bool
	Revers bool
	Stdout bool
	Number bool
}

func main() {
	flags := new(Flag)
	parseFlag(flags)

	filename := os.Args[len(os.Args)-1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	dataSlice := strings.Split(string(data), "\n")

	switch {
	case flags.NoRepetition == true: // -u
		Repetitionsort(dataSlice)
	case flags.NoRegister == true: // -f
		RegistorSort(dataSlice)
	case flags.Revers == true: // -r
		ReverseSort(dataSlice)
//	case flags.Stdout == true: // -n
	//	FileWrite(dataSlice)
	case flags.Number == true :
		NumberSort(dataSlice)
	}

	//sort.Strings(dataSlice)
	//fmt.Println(strings.Join(dataSlice, " "))

}

func parseFlag(flags *Flag) {
	flag.BoolVar(&flags.NoRepetition, "u", false, "Repetition sort")
	flag.BoolVar(&flags.NoRegister, "f", false, "Register sort")
	flag.BoolVar(&flags.Revers, "r", false, "Revers sort")
	flag.BoolVar(&flags.Stdout, "n", false, "Number sort")
	flag.Parse()
}

func RegistorSort(data []string) {
	sort.Slice(data, func(i, j int) bool {return strings.ToLower(data[i]) < strings.ToLower(data[j]) })
	//fmt.Println(data)
}

func Repetitionsort(data []string) {
var newdata string
	for i := 0; i < len(data); i++ {
		flags := 0
	for j := i + 1; j < len(data); j++ {
		if data[i] == data[j] {
			flags++
		}
}
	if flags == 0 {
		newdata = newdata + data[i] + " "
	}
}

sortdata := strings.Split(newdata, " ")
//fmt.Println(sortdata)
sort.Strings(sortdata)
fmt.Println(sortdata)
}

func ReverseSort(data []string) {
	sort.Slice(data, func(i,j int) bool {return data[i] > data[j]})
	fmt.Println(data)
}

func NumberSort (data []string)  {
	for i := 0; i < len(data); i++ {
		//if data[i] !=
			}
}