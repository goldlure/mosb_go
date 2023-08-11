package models

import "fmt"

var InputMap = map[string]func() string{
	"/search":   search,
	"/giveaway": giveAway,
	"/notFound": notFound,
}

func search() string {
	fmt.Println("Search!")
	return "Search!"
}

func giveAway() string {
	fmt.Println("GiveAway!")
	return "GiveAway!"
}

func notFound() string {
	fmt.Println("notFound!")
	return "your request is not supported, please, choose another"
}
