package main

import (
	"flag"
	"fmt"
	"strconv"
)

const printableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r"

func main() {
	regex := flag.String("g", "", "Regex string (Required)")
	limit := flag.Int("l", 20, "Max limit for range size")
	count := flag.Bool("c", false, "Count matching strings")
	maxNumber := flag.Int("m", -1, "Max number of strings")
	random := flag.Bool("r", false, "Returns a random string that matches to the regex")
	simplify := flag.Bool("s", false, "Simplifies a regular expression")
	delimiter := flag.String("d", "\n", "Delimiter - default is \\n")

	flag.Parse()

	if *regex != "" {
		fmt.Println(string(*regex))
		if *limit != 20 {
			fmt.Println(string(*limit))
		}
		if *count {
			counter := strconv.FormatBool(*count)
			fmt.Println(counter)
		}
		if *maxNumber != -1 {
			fmt.Println(string(*maxNumber))
		}
		if *random {
			randomer := strconv.FormatBool(*random)
			fmt.Println(randomer)
		}
		if *simplify {
			simplifier := strconv.FormatBool(*simplify)
			fmt.Println(simplifier)
		}
		if *delimiter != "\n" {
			fmt.Println(string(*delimiter))
		}
	} else {
		fmt.Println("exit!")
	}

}
