package piscine

import "fmt"

//IsNegative ...verifier si un nombre est negatif
func IsNegative(nb int) {
	result := "F"
	if nb < 0 {
		result = "T"
	}

	fmt.Println(result)
}
