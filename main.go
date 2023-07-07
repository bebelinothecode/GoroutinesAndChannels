package main

import (
	"fmt"
	"sync"
	"time"
)



type Data struct {
	Timestamp time.Time
	Payload map[string]interface{}
}

func main() {
	dataChannel := make(chan Data)
	//Create the consumer
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataChannel {
			fmt.Printf("received data at %s:\n",data.Timestamp.Format("2006-01-03 15:04:05"))
			fmt.Printf("Payload: %v\n", data.Payload)
		}

	}()

	//Producer 1
	var producerWg sync.WaitGroup
	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		fmt.Println("Prodocer 1 has started")
		time.Sleep(2 * time.Second)
		payload := map[string]interface{}{"sensor":"Temperature", "value":25.5}
		data := Data{Timestamp: time.Now(), Payload: payload}
		dataChannel <- data
		fmt.Println("producer 1 sent data")
	}()

	//Producer 2	
	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		fmt.Println("Producer 2 started")
		time.Sleep(4 * time.Second)
		payload := map[string]interface{}{"sensor":"Pressure", "value":40}
		data := Data{Timestamp: time.Now(), Payload: payload}
		dataChannel <- data
		fmt.Println("producer 2 sent data")
	}()

	producerWg.Wait()
	close(dataChannel)
	wg.Wait()
	fmt.Println("All the data processed")
}