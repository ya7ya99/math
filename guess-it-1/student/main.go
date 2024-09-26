package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num int
	var numbers []int
	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			fmt.Println(err)
			break
		}
		if num >= 1000 || num <= -1000 {
			fmt.Println(-num, num)
		} else {
			numbers = append(numbers, num)
			Guess_Range(numbers)
		}
	}
}

func Guess_Range(numbers []int) {
	avg := Average(numbers)
	if len(numbers) == 1 {
		fmt.Println(50, 900)
	} else {
		 sd := Deviation(numbers, avg)
		fmt.Println(int(avg-sd), int(avg+sd))
	}
}

func Deviation(numbers []int, avg float64)  float64 {
	var sum float64
	for _, num := range numbers {
		sum += math.Pow(float64(num)-avg, 2)
	}
	variance := sum / float64(len(numbers))
	return  math.Sqrt(variance)
}

func Average(numbers []int) float64 {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}
