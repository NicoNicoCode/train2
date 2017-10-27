package main

import (
	"fmt"
	"time"
	"math/rand"
)

type HashTb1 struct{
	tableSize int
	theLists []ListNode
}

type ListNode struct{
	element int
	next *ListNode
}

type Position *ListNode
type HashTable *HashTb1

type List Position

func hash(key int,tableSize int)(int){
	return key%tableSize
}

func initTable(tableSize int)(HashTable){
	var h HashTb1

	h.tableSize = tableSize
	h.theLists = make([]ListNode,tableSize)

	return &h
}

func find(key int,h HashTable)(Position){
	var p Position
	var l List

	l = &h.theLists[hash(key,h.tableSize)]
	p = l.next
	for p!=nil && p.element != key{
		p = p.next
	}

	return p
}

func insert(key int,h HashTable){
	var pos Position
	var newCell ListNode
	var l List

	pos = find(key,h)
	if pos == nil{
		l = &h.theLists[hash(key,h.tableSize)]
		newCell.next = l.next
		newCell.element = key
		//回想链表的插入操作
		l.next = &newCell
	}

}

func main(){
	  fmt.Println("新建一个Hash表,表长为64")
	  var test HashTable
	  var fun Position
	  test=initTable(64)

	  r:=rand.New(rand.NewSource(time.Now().UnixNano()))

	  fmt.Println("插入随机的100个数")
	  for i:=0;i<100;i++{
		  insert(r.Intn(1000),test)
	  }

	  fmt.Println("打印Hash表")
	  for i:=0;i<test.tableSize;i++{
		  fmt.Printf("%d:",i)
		  fun = &test.theLists[i]
		  for fun.next!=nil{
		      if fun.next.next!=nil{
				  fmt.Printf("%d->",fun.next.element)
			  }else{
				  fmt.Printf("%d",fun.next.element)
			  }
			  fun=fun.next
		  }
		  fmt.Println()
	  }
	  
}