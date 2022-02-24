package main

import (
	"fmt"
	"time"
)

func main() {

	// for {
	// 	conn, err := ldap.Dial("ldap://140.124.181.91")

	// }

	// if err != nil {
	// 	panic("error")
	// }

	ch := make(chan int)

	go func() {
		fmt.Println("[T1]thread 1")

		time.Sleep(time.Second * 3)

		ch <- 1
	}()

	// go func() {
	// 	fmt.Println("[T2]thread 2")
	// 	fmt.Println("[T2]waiting for reading data from channel")

	// 	data := <-ch

	// 	fmt.Println("[T2]data got!", data)
	// }()

	fmt.Println("waiting for data...")
	data := <-ch

	fmt.Println("data got!", data)
}
