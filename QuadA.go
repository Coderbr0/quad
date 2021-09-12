package piscine

import "fmt"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 { // for (1,1)
			fmt.Println("o")
		} else if x == 1 && y > 1 { // for (1,5)
			fmt.Println("o")
			for i := 0; i <= y-3; i++ {
				fmt.Println("|")
			}
			fmt.Println("o")
		} else if x > 1 && y == 1 { // for (5,1)
			fmt.Print("o")
			for i := 0; i <= x-3; i++ {
				fmt.Print("-")
			}
			fmt.Print("o")
		} else if x > 1 && y > 1 { // for (5,3)
			fmt.Print("o") // first line
			for i := 0; i <= y-1; i++ {
				fmt.Print("-")
			}
			fmt.Print("o\n")            // end of first line
			for i := 0; i <= y-3; i++ { // Middle Line
				fmt.Print("|")
				for i := 0; i <= y-1; i++ {
					fmt.Print(" ")
				}
				fmt.Println("|") // end of middle
			}
			fmt.Print("o") // Last line
			for i := 0; i <= y-1; i++ {
				fmt.Print("-")
			}
			fmt.Print("o\n") // The End
		}
	} else {
		fmt.Println('\n')
	}
}
