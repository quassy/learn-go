package main

import (
	"fmt"
	"math"
)

func coffee() {
	var cups, water, milk, beans = 1, 200, 50, 15
	var cups_max, water_max, milk_max, beans_max = 0, 0, 0, 0

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&water_max)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milk_max)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&beans_max)

	cups_max = int(math.Min(float64(water_max/water), float64(milk_max/milk)))
	cups_max = int(math.Min(float64(cups_max), float64(beans_max/beans)))

	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	switch {
	case cups > cups_max:
		fmt.Printf("No, I can make only %d cups of coffee\n", cups_max)
	case cups == cups_max:
		fmt.Printf("Yes, I can make that amount of coffee\n")
	case cups < cups_max:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", cups_max-cups)
	}

	// water *= cups  // in ml
	// milk *= cups  // in ml
	// beans *= cups  // in g
	// fmt.Printf("For %d cups of coffee you will need:\n", cups)
	// fmt.Printf("%d ml of water\n", water)
	// fmt.Printf("%d ml of milk\n", milk)
	// fmt.Printf("%d g of coffee beans\n", beans)
}
