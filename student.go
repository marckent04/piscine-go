package main

import "fmt"

//Raid1a ...verifier si un nombre est negatif
func Raid1a(x, y int) {
	hch := make([]string, x)
	hcb := make([]string, x)
	vc := make([]string, x)

	//horizontal haut
	for i := 0; i < x; i++ {
		if i == 0 {
			hch[i] = "/"
		} else if i == (x - 1) {
			hch[i] = "\\"
		} else {
			hch[i] = "*"
		}
	}

	//horizontal bas
	for i := 0; i < x; i++ {
		if i == 0 {
			hcb[i] = "\\"
		} else if i == (x - 1) {
			hcb[i] = "/"
		} else {
			hcb[i] = "*"
		}
	}

	//verticale
	for i := 0; i < len(vc); i++ {
		if i == 0 {
			vc[i] = "*"
		} else if i == (len(vc) - 1) {
			vc[i] = "*"
		} else {
			vc[i] = " "
		}
	}

	for i := 0; i < x; i++ {
		if i == (x - 1) {
			fmt.Println(hch[i])
		} else {
			fmt.Print(hch[i])
		}
	}

	for a := 0; a < (y - 2); a++ {
		for i := 0; i < len(vc); i++ {
			fmt.Print(vc[i])
		}
		fmt.Print("\n")
	}

	if y > 1 {
		for i := 0; i < x; i++ {
			if i == (x - 1) {
				fmt.Println(hcb[i])
			} else {
				fmt.Print(hcb[i])
			}
		}
	}

}

func main() {
	Raid1a(1, 5)
}
