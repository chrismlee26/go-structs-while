package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type House struct {
	numberOfRooms int
	city          string
	address       string
	price         int
}

func main() {
	newHouse := []House{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter number of rooms in the house:")
	user_num_rooms, _ := reader.ReadString('\n')
	int_user_num_rooms, _ := strconv.ParseInt(strings.TrimSpace(user_num_rooms), 10, 32)
	fmt.Println("Enter city:")
	user_city, _ := reader.ReadString('\n')
	fmt.Println("Enter address:")
	user_address, _ := reader.ReadString('\n')
	fmt.Println("Enter price:")
	user_price, _ := reader.ReadString('\n')
	int_user_price, _ := strconv.ParseInt(strings.TrimSpace(user_price), 10, 32)

	newHouse = append(newHouse, House{int(int_user_num_rooms), user_city, user_address, int(int_user_price)})

	for _, house := range newHouse {
		fmt.Printf("\r%v\t %s\t %s\t %v", house.numberOfRooms, house.city, house.address, house.price)
	}

	// fmt.Println(newHouse)

}
