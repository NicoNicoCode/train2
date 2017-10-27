package main

import (
	"fmt"
)

type Stack struct{
	element int
	next *Stack
}

type pNode *Stack

func isEmpty(s pNode)(int){
	if s.next == nil{
		return 1
	}else{
		return 0
	}
}

func push(x int,s pNode){
	var tmpCell Stack
	tmpCell.element = x
	if s !=nil{
		tmpCell.next = s.next
	}
	s.next = &tmpCell
}

func top(s pNode)(int){
	if isEmpty(s) == 0 {
		return (s.next).element
	}else{
		fmt.Println("这是个空栈")
		return 0
	}
}

func pop(s pNode){

	if isEmpty(s)==1 {
		fmt.Println("这是个空栈，不能弹出元素")
	}else{
		s.next = s.next.next
		s=s.next
	}
	
}

func makeEmpty(s pNode){
	for isEmpty(s)==0{
		pop(s)
	}
}

func createStack()(pNode){
	var s Stack
	s.next = nil
	return &s
}



func main(){
	fmt.Println("新建一个栈")
	p := createStack()
	fmt.Println("入栈10个数")
	for i:=0;i<10;i++{
		push(i,p)
	}
	fmt.Printf("栈顶元素%d\n",top(p))
	fmt.Println("出栈5个数")
	for i:=0;i<5;i++{
		pop(p)
	}
	fmt.Printf("栈顶元素%d\n",top(p))
	fmt.Println("清空栈。。。")
	makeEmpty(p)
	fmt.Printf("栈顶元素%d\n",top(p))
}