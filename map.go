package main

import "fmt"

func main() {
	var capital map[string]string
	capital = make(map[string]string)
	capital["Italia"] = "Roma"
	capital["França"] = "Paris"
	for capitalmap := range capital {
		fmt.Println("A capital do ", capitalmap, "é ", capital[capitalmap])
	}

}
