package main

import (
	"fmt"
)

type stack struct{
	top int
	capacity int
	value []interface{}
}

type queue struct{
	top int
	capacity int
	value []interface{}
}

func (s *stack)push(value int){
	if s.top == (s.capacity) - 1{
		fmt.Println("Stack Overloaded")
	}else{
		fmt.Println("Pushing value....!!")
		s.top++
		s.value[s.top] = value
		fmt.Printf("Resultant Stack is : \n%+v\n",s)
	}
}
func (s *stack)pop(){
	if s.top == -1{
		fmt.Println("Stack empty")
	}else{
		fmt.Println("Popping out :", s.value[s.top])
		s.value = s.value[:s.top]
		s.top--
		fmt.Printf("Resultant stack is : \n%+v",s)
	}
}

func (q *queue)enqueue(value int){
	if q.top == (q.capacity) -1{
		fmt.Println("Queue Overloaded")
	}else{
		fmt.Println("Enqueue value...!!")
		q.top++
		q.value[q.top] = value
		fmt.Printf("Resultant Queue is : \n%+v\n", q)
	}
}

func (q *queue)dequeue(){
	if q.top == -1{
		fmt.Println("Queue Empty")
	}else{
		fmt.Println("Dequeue value :", q.value[0])
		q.value = q.value[1:]
		q.top--
		fmt.Printf("Resultant Queue is : \n%+v\n",q)

	}
}
func main(){
	fmt.Println("Stack Implementation Using Go")
	s := stack{
		top: -1,
		capacity: 5,
		value:	make([]interface{}, 5),
	}
	s.push(5)
	s.push(6)
	s.push(7)
	s.pop()

	fmt.Println("Queue Implementation Using Go")
	q := queue{
		top: -1,
		capacity: 5,
		value: make([]interface{},5),
	}
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.dequeue()


}