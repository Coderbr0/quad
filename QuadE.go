package piscine

import "fmt"

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 { // for (1,1)
			fmt.Println("A")
		} else if x == 1 && y > 1 { // for (1,5)
			fmt.Println("A")
			for i := 0; i <= y-3; i++ {
				fmt.Println("B")
			}
			fmt.Println("C")
		} else if x > 1 && y == 1 { // for (5,1)
			fmt.Print("A")
			for i := 0; i <= x-3; i++ {
				fmt.Print("B")
			}
			fmt.Print("C")
		} else if x > 1 && y > 1 { // for (5,3)
			fmt.Print("A") // first line
			for i := 0; i <= y-1; i++ {
				fmt.Print("B")
			}
			fmt.Print("C\n")            // end of first line
			for i := 0; i <= y-3; i++ { // Middle Line
				fmt.Print("B")
				for i := 0; i <= y-1; i++ {
					fmt.Print(" ")
				}
				fmt.Println("B") // end of middle
			}
			fmt.Print("C") // Last line
			for i := 0; i <= y-1; i++ {
				fmt.Print("B")
			}
			fmt.Print("A\n") // The End
		}
	} else {
		fmt.Println('\n')
	}
}
