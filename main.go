package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type House struct {
	numberOfRooms string
	city          string
	address       string
	price         int
}

func main() {
	newHouse := []House{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter number of rooms in the house:")
		user_num_rooms, _ := reader.ReadString('\n')
		fmt.Println("Enter city:")
		user_city, _ := reader.ReadString('\n')
		fmt.Println("Enter address:")
		user_address, _ := reader.ReadString('\n')
		fmt.Println("Enter price:")
		user_price, _ := reader.ReadString('\n')
		int_user_price, _ := strconv.ParseInt(strings.TrimSpace(user_price), 10, 32)

		fmt.Println("Continue? (y/n)")
		continue_loop, _ := reader.ReadString('\n')

		if strings.TrimSpace(continue_loop) == "y" {
			newHouse = append(newHouse, House{user_num_rooms, user_city, user_address, int(int_user_price)})
		} else {
			newHouse = append(newHouse, House{user_num_rooms, user_city, user_address, int(int_user_price)})
			fmt.Println("Exiting...")
			for _, house := range newHouse {
				fmt.Println("\n ~~~~~~~Houses for Sale:~~~~~~~")
				fmt.Printf("%-5s %-5s %-5s rooms %d \n", house.address, house.city, house.numberOfRooms, house.price)
			}
			os.Exit(0)
		}

		// newHouse = append(newHouse, House{user_num_rooms, user_city, user_address, int(int_user_price)})

		for _, house := range newHouse {
			fmt.Println("\n ~~~~~~~Houses for Sale:~~~~~~~")
			fmt.Printf("%-5s %-5s %-5s rooms %d \n", house.address, house.city, house.numberOfRooms, house.price)
		}
	}
}
