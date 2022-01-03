package main

import (
	"fmt"
)

type drink struct {
	name  string
	water int
	milk  int
	beans int
	price int
	cups  int
}

func coffee2() {
	espresso := drink{
		name:  "Espresso",
		water: 250,
		beans: 16,
		price: 400,
		cups:  1,
	}
	latte := drink{
		name:  "Latte",
		water: 350,
		milk:  75,
		beans: 20,
		price: 700,
		cups:  1,
	}
	cappucino := drink{
		name:  "Cappuciono",
		water: 200,
		milk:  100,
		beans: 12,
		price: 600,
		cups:  1,
	}
	selection_drink := drink{}
	fill_drink := drink{}
	inventory := drink{
		name:  "Inventory",
		water: 400,
		milk:  540,
		beans: 120,
		price: 55000,
		cups:  9,
	}
	var command, selection string

	for command != "exit" {
		fmt.Print("Write action (buy, fill, take, remaining, exit): ")
		fmt.Scanf("%s", &command)
		switch command {
		case "buy":
			fmt.Print("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: ")
			fmt.Scanf("%s", &selection)
			switch selection {
			case "espresso", "1":
				selection_drink = espresso
			case "latte", "2":
				selection_drink = latte
			case "cappucino", "3":
				selection_drink = cappucino
			default:
				fmt.Println("Invalid selection")
			}

			if selection_drink.water > inventory.water {
				fmt.Println("Sorry not enough water!")
				break
			} else if selection_drink.milk > inventory.milk {
				fmt.Println("Sorry not enough milk!")
				break
			} else if selection_drink.beans > inventory.beans {
				fmt.Println("Sorry not enough beans!")
				break
			} else if selection_drink.cups > inventory.cups {
				fmt.Println("Sorry not enough cups!")
				break
			} else {
				fmt.Println("I have enough resources, making you a coffee!")
			}
			inventory.water -= selection_drink.water
			inventory.milk -= selection_drink.milk
			inventory.beans -= selection_drink.beans
			inventory.cups -= selection_drink.cups
			inventory.price += selection_drink.price
		case "fill":
			fmt.Print("Write how many ml of water you want to add: ")
			fmt.Scanf("%d", &fill_drink.water)
			fmt.Print("Write how many ml of milk you want to add: ")
			fmt.Scanf("%d", &fill_drink.milk)
			fmt.Print("Write how many grams of coffee beans you want to add: ")
			fmt.Scanf("%d", &fill_drink.beans)
			fmt.Print("Write how many disposable coffee cups you want to add: ")
			fmt.Scanf("%d", &fill_drink.cups)
			inventory.water += fill_drink.water
			inventory.milk += fill_drink.milk
			inventory.beans += fill_drink.beans
			inventory.cups += fill_drink.cups
		case "take":
			fmt.Printf("I gave you %.2f$\n", float64(inventory.price/100))
			inventory.price = 0
		case "remaining":
			fmt.Printf("The coffee machine has:\n")
			fmt.Printf("%d of water\n", inventory.water)
			fmt.Printf("%d of milk\n", inventory.milk)
			fmt.Printf("%d of coffee beans\n", inventory.beans)
			fmt.Printf("%d of disposable cups\n", inventory.cups)
			fmt.Printf("$%d of money\n", inventory.price/100)
			fmt.Println()
		case "exit":
			fmt.Println("Exiting.")
		default:
			fmt.Println("Not implemented.")
		}
	}
}
