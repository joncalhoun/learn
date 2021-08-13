package main

import (
	"fmt"
	"sync"
	"time"
)

type Inventory struct {
	stock int
	mutex sync.RWMutex
}

func (i *Inventory) Stock() int {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	return i.stock
}

func (i *Inventory) DecreaseStock(amount int) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.stock-amount < 0 {
		return fmt.Errorf("inventory too low")
	}
	i.stock = i.stock - amount
	return nil
}

func main() {
	count := 10000
	var inventory Inventory
	inventory.DecreaseStock(-count)
	fmt.Println("Initial stock:", inventory.Stock())

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(who string) {
			BuyPhone(&inventory, who)
			wg.Done()
		}(fmt.Sprintf("Person %d", i))
	}

	wg.Wait()
	fmt.Println("Final inventory:", inventory.Stock())
}

func BuyPhone(i *Inventory, who string) {
	err := i.DecreaseStock(1)
	if err != nil {
		fmt.Println("Inventory low")
		return
	}

	// check stock
	// do a credit check
	time.Sleep(1 * time.Second)
	// if credit check fails
	creditCheckFail := false
	if creditCheckFail {
		i.DecreaseStock(-1)
		return
	}
	// finaly, buy phone
}

func contrived() {
	balance := 100

	// You visit the ATM and withdraw 60
	yourNewBalance := balance - 60

	// partner also checks out at the grocery store
	partnerNewBalance := balance - 44

	// You at the ATM
	if yourNewBalance > 0 {
		// spits out money
		// and then...
		balance = yourNewBalance
	}

	// At the grocery store
	if partnerNewBalance > 0 {
		// approve transaction
		balance = partnerNewBalance
	}

	fmt.Println("Your balance is:", balance)
}
