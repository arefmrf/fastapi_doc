package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func printIt(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

/* func main() {
	 words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	//wg := sync.WaitGroup  why its not working? why var x struct works?
	var wg sync.WaitGroup
	wg.Add(len(words))
	for i, word := range words {
		go printIt(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait()
} */

/* ===========================================  */

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("recieved order #%d\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("making pizza #%d takes %d seconds. . .\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			msg = fmt.Sprintf("pizza #%d failed by reason 1.", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("pizza #%d failed by reason 2.", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza #%d is ready.", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzeria(pizzaJob)
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				fmt.Println("*** ", i.message, " ***")
				fmt.Printf("*** order #%d is out for delivery\n", i.pizzaNumber)
			} else {
				fmt.Println("*** ", i.message, " ***")
				fmt.Printf("*** order #%d is !@#$\n", i.pizzaNumber)
			}
		} else {
			fmt.Println("*** Done making pizza ***")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("Failed closing channel", err)
			}
		}
	}

	fmt.Printf("###  in total of %d attempt %d pizza made, %d pizza failed", total, pizzasMade, pizzasFailed)
}
