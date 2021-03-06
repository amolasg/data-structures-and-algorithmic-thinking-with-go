/*
# Copyright (c) July 12, 2020 CareerMonk Publications and others.
# E-Mail           		: info@careermonk.com
# Creation Date    		: 2020-07-12 06:15:46
# Last modification		: 2020-07-12
#               by		: Narasimha Karumanchi
# Book Title			: Data Structures And Algorithmic Thinking With Go
# Warranty         		: This software is provided "as is" without any
# 				   warranty; without even the implied warranty of
# 				    merchantability or fitness for a particular purpose. */

package main
import "fmt"

// print numbers 1 to n backward
func print(n int) int {
	if n == 0 {
		return 0 // this is the terminating base case
	}
	fmt.Println(n)
	return print(n - 1) // recursive call to itself again
}

func main() {
	fmt.Println(print(5))
}
