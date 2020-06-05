package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var Zones = [...]string{
	"Africa",
	"America",
	"Antartica",
	"Asia",
	"Atlantic",
	"Australia",
	"Europe",
	"Indian",
	"Pacific",
}

func main() {
	city := strings.Join(os.Args[1:], " ")

	for _, zone := range Zones {
		locationName := fmt.Sprintf("%s/%s", zone, ToPascalSnakeCase(city))

		location, err := time.LoadLocation(locationName)

		if err == nil {
			t := time.Now().In(location)
			fmt.Printf("%02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
			return
		}
	}

	fmt.Printf("timein: The location %s does not exist\n", city)
}
