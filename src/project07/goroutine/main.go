package main
import (
	"fmt"
	_ "time"
)

func writeData(intChan chan int) {
	for i:=1;i<=50;i++ {
		intChan<- i
		fmt.Println("writeData=", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData=%v\n", v)
	}
	exitChan<- true
	close(exitChan)
}


func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	// time.Sleep(time.Second * 10)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}