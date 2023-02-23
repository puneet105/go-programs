//encrypt and decrypt using go routines and two channels
package main

import (
	"fmt"
)

func encrypt(val string, enc chan string){
	val = val + "SHARMA"
	enc<-val
}

func decrypt(encodedValue string, dec chan string){
	encodedValue = "PUNEET"
	dec <- encodedValue
}

func main(){
	val := "PUNEET"
	enc := make(chan string)
	dec := make(chan string)
	go encrypt(val, enc)
	encodedValue := <-enc
	fmt.Println("Encoded value is", encodedValue)
	go decrypt(encodedValue, dec)
	decodedValue := <-dec
	fmt.Println("Decoded Value is", decodedValue)
}

/*
//using single channel
package main

import (
	"fmt"
)

func encrypt(val string, enc chan string){
	val = val + "SHARMA"
	enc<-val
}

func decrypt(encodedValue string, dec chan string){
	encodedValue = "PUNEET"
	dec <- encodedValue
}

func main(){
	val := "PUNEET"
	enc := make(chan string)
	dec := make(chan string)
	go encrypt(val, enc)
	encodedValue := <-enc
	fmt.Println("Encoded value is", encodedValue)
	go decrypt(encodedValue, dec)
	decodedValue := <-dec
	fmt.Println("Decoded Value is", decodedValue)
}
*/