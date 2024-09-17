package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	f "math-skills/staticalcul"
)

func main()  {
	if len(os.Args) < 2{
		fmt.Println("USAGE: go run main.go 'filename' ")
		return
	}
	
	filename := os.Args[1]
	
	//Open and ReadFile line by line, add the value in data to print the results of statical calcul
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s:  no such file or directory\n", os.Args[1])
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan(){
			num, err := strconv.Atoi(scanner.Text())
			if err != nil{
				fmt.Println("Error: The file must contain only numbers")
				return
			}
			data = append(data, num)
	}
	av, v, stDev:= f.WelfordAlgo((data))
	m := f.MedianCalcul(data)

	fmt.Printf("\033[0;32m==== Calculation of statistical measures ====\033[0m\n")
	fmt.Printf("Average: %d\n", av)
	fmt.Printf("Median: %d\n", m)
	fmt.Printf("Variance: %d\n", v)
	fmt.Printf("Standard Deviation: %d\n", stDev)
}
