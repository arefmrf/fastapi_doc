package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		fmt.Printf("** %s goes to the waiting room to check for clients. **\n", barber)
		for {
			if len(shop.ClientsChan) == 0 {
				fmt.Printf("** No client so %s takes a nap. **\n", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {

					fmt.Printf("** %s wakes %s up. **\n", client, barber)
					isSleeping = false
				}
				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}
func (shop *BarberShop) cutHair(barber, client string) {
	fmt.Printf("** %s is cutting %s's hair. **\n", barber, client)
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("** %s finished cutting %s's hair. **\n", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("** %s is going home. **\n", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Printf("** Closing shop for the day **\n")
	close(shop.ClientsChan)
	shop.Open = false
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	fmt.Printf("** The barber shop is now closed **\n")
}

func (shop *BarberShop) addClient(client string) {
	fmt.Printf("** %s arrives. **\n", client)
	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			fmt.Printf("** %s takes a seat in waiting room. **\n", client)
		default:
			fmt.Printf("** The waiting room is full so %s leaves. **\n", client)
		}
	} else {
		fmt.Printf("** The shop is already closed so %s leaves. **\n", client)
	}

}

func main() {
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}
	fmt.Println("The shop is Open . . .")
	shop.addBarber("Frank")
	shop.addBarber("XFrank")
	shop.addBarber("XXFrank")

	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1
	go func() {
		for {
			randomMillSecond := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillSecond)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()
	<-closed
}
