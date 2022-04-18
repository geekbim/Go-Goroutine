package go_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// send data to channel
	channel <- "test channel"

	// receive data from channel
	data := <-channel
	fmt.Println(<-channel)
	fmt.Println(data)

	// close channel
	defer close(channel)
}

func TestCreateChannelV2(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Abim"
		fmt.Println("Finish Send Data to Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Abim"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel in and out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Abim"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// buffer channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// example 1
	// channel <- "Abim"
	// channel <- "Dhanu"
	// channel <- "Ejas"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// example 2
	go func() {
		channel <- "Abim"
		channel <- "Dhanu"
		channel <- "Ejas"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Selesai")
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data", data)
	}

	fmt.Println("Finish")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// select default
func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channel 2", data)
			counter++
		default:
			fmt.Println("Waiting Data...")
		}
		if counter == 2 {
			break
		}
	}
}
