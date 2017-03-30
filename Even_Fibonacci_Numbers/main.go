package main

import "fmt"

func main (){
	var x[50] int32 //also slice, x:= make([] int, 50)
	var fib_sum int32
	x[1] = 1 //initialize these values so there is something to add
	x[2] = 2
	fmt.Println("All even Fibonacci numbers < 4,000,000:")
	for i := 2; i<50; i++ {x[i] = (x[i-1] + x[i-2])
		if (x[i]<4000000)&&(x[i] % 2 == 0)&&(x[i]>0){fib_sum += x[i]
			fmt.Println(x[i])
		}
	}
	fmt.Println("And the sum of all even Fibonacci numbers < 4,000,000 is:", fib_sum)
}
