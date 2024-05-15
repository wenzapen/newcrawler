package persist

import "log"

func ItemSaver() chan interface{} {
	itemChan := make(chan interface{})
	itemCount := 0
	go func() {
		for {
			item := <-itemChan
			log.Printf("got item #%d: %v", itemCount, item)
			itemCount++
		}

	}()

	return itemChan
}
