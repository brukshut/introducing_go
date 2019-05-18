package main

import ("fmt"; "time")

func main() {
	c1 := make(chan string, 5)
	c2 := make(chan string, 5)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <- c2:
				fmt.Println("Message 2", msg2)
			case time := <- time.After(time.Second * 2):
				fmt.Println("Timeout", time)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
