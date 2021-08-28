package main

import "fmt"

func huiwen(i int) {
	var j int
	if i < 0 {
		fmt.Println("not is huiwen")
		return
	}
	tmp := i
	for tmp != 0 {
		j = j*10 + tmp%10
		tmp /= 10
	}

	if i == j {
		fmt.Println("is huiwen")
	} else {
		fmt.Println("not is huiwen")
	}
}

func main() {
	huiwen(121)
	huiwen(-121)
}
