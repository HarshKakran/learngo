package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// var c = make(chan int, 6)
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walamrt.com", "blonkit.com", "swiggy.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price < MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}

}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// select statement is if statment for channels
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found an deal on Chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail sent: Found a dela on tofu at %v", website)
	}
}
