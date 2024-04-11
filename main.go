package main

import (
	"fmt"
	"sort"
)

type Laptop struct {
	Gain  float64
	Price float64
}

func genLaptops() []Laptop {
	laptops := []Laptop{
		{3, 4},
		{2, 5},
		{1, 6},
		{5, 4},
		{7, 3},
		{7, 4},
		{8, 2},
		{12, 6},
		{14, 8},
	}
	return laptops
}

func getMaxCapital(N int, C float64, laptops []Laptop) float64 {
	if len(laptops) == 0 || N == 0 || C == 0 {
		return C
	}
	
	// max <-
	sort.Slice(laptops, func(i, j int) bool {
		return ((laptops[i].Gain - laptops[i].Price) / laptops[i].Price) > ((laptops[j].Gain - laptops[j].Price) / laptops[j].Price)
	})

	for N != 0 {
		for i, v := range laptops {
			if v.Price <= C {
				profit := v.Gain - v.Price
				C += profit
				laptops = append(laptops[:i], laptops[i+1:]...)
				break
			} else if laptops[len(laptops)-1].Price > C && i == len(laptops)-1 {
				return C
			}
		}
		N--
	}
	return C
}

func main() {
	N := 3
	C := 5.0
	laptops := genLaptops()
	fmt.Println(laptops)
	newLaptops := getMaxCapital(N, C, laptops)
	fmt.Println(newLaptops)
}
