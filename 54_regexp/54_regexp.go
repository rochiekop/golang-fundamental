package main

import (
	"fmt"
	"regexp"
)

func main() {

	re := regexp.MustCompile(`\.(.*) `)
	match := re.FindStringSubmatch(".d 1000=11,12")
	if len(match) != 0 {
		fmt.Printf("1. %s\n", match[1])
	}

	match = re.FindStringSubmatch("e 2000=11")
	if len(match) != 0 {
		fmt.Printf("2. %s\n", match[1])
	}

	match = re.FindStringSubmatch(".e2000=11")
	if len(match) != 0 {
		fmt.Printf("3. %s\n", match[1])
	}

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)

}
