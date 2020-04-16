package main

func Sum(numbers [5]int) int{
	//Code I wrote
	var sum int
	var number int
	for _, number = range numbers{
		sum += number
	}

	//Chris James' code example
	//sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	return sum

}
