//go: main.go
// go routines for congruent programming in go

package main

import (
	"fmt"
	"sync"
)

func genreateNmumber(total int, ch chan<- int, waitgroup *sync.WaitGroup) {
	//write into the channel using ch <- value
	defer waitgroup.Done() // decrements the counter by 1 of the waiting group when the goroutine completes
	defer close(ch)        // close the channel to not acculumate and allocate memories for the channel and create a memory leak

	for i := 1; i <= total; i++ {
		fmt.Print("generating number:")
		fmt.Println(i)
		ch <- i
	}
}

func printNumber(ch <-chan int, waitgroup *sync.WaitGroup) {
	//read from the channel assigining the value of <- ch
	defer waitgroup.Done()

	for i := range ch {
		fmt.Print("printing number:")
		fmt.Println(i)
	}
}

func readChannel(ch <-chan int)    {}
func writeChannel(ch chan<- int)   {}
func readWriteChannel(ch chan int) {}

func main() {
	var waitgroup sync.WaitGroup

	//-----------------------------------
	// creating  a channels for better communication between goroutines
	// var intChannel chan int = make(chan int, 10)
	// intChannel <- 10 // writing into the channel's
	// for functionn we are writing value into the channels while writing so for parameter s ch chan  <- int

	// var _ = <-intChannel // reading from the channel's
	// __ := <-intChannel   // equivalence to var __ int = <-intChannel  using short hand declaration
	// _ = __
	// writing we are reading value from the channels while assigning it to the variable  while reading so for parameter s ch <- chan int
	//-----------------------------------

	waitgroup.Add(2) // waiting for 2 goroutines to finish
	/*
	   .Add(3) deadlock as two coroutine are done and the main goroutine is waiting for the third goroutine to finish
	*/
	var numberChannel chan int = make(chan int)

	go genreateNmumber(10, numberChannel, &waitgroup) // start a new goroutine
	go printNumber(numberChannel, &waitgroup)         // start a new goroutine

	// both of  the functions are waiting for each other to write and read from the channel so the deadlock occurs

	/*
		The deadlock can happen due to the way channel communication works in Go. When part of a program is writing to a channel, it will wait until another part of the program reads from that channel before continuing on. Similarly, if a program is reading from a channel it will wait until another part of the program writes to that channel before it continues
	*/
        // in deadlock the coroutines will be in the blocked state 

	// instead of using chan <- int and <- chan int we can use chan int to read and write from the channel  to prevent deadlock

	// coroutine control is dfferent for everyone in go so we can't predict the order of the execution of the goroutines
	// if deadlock decrease wait goroutine by 1 and check if the deadlock is resolved or not
	// if -ve deadlock increase the wait goroutine by 1 and check if the deadlock is resolved or not

	println("waiting for goroutines to finish")
	waitgroup.Wait() // wait for all goroutines to finish

	println("all goroutines finished")

}
