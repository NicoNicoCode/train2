package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Heap struct{
	capacity int
	size int
	elements []int
}

type pQueue *Heap

func initHeap(max int)(pQueue){
	var h Heap
	
	h.capacity = max
	h.size = 0
	h.elements = make([] int,max)

	return pQueue(&h)
}

func destroy(h pQueue){
	  h.elements = nil
	  h.capacity = 0
	  h.size = 0
	  h = nil
}

func makeEmpty(h pQueue){
	h.elements = nil
	h.size = 0
	h.capacity = 0
}

func insert(x int,h pQueue){
	var i int

	if isFull(h){
		fmt.Println("Priority queue is full")
		return 
	}

	h.size++
	for i=h.size;h.elements[i/2]>x;i/=2{
		h.elements[i]=h.elements[i/2]
	}
	h.elements[i]=x

}

func deleteMin(h pQueue)(int){
	var i,child int
	var minElement,lastElement int

	if(isEmpty(h)){
		fmt.Println("Priority queue is empty")
		return h.elements[0]
	}
	minElement = h.elements[1]
	lastElement = h.elements[h.size]
	h.size--

	for i=1;i*2<=h.size;i=child{
		child = i*2
		if(child != h.size && h.elements[child + 1] < h.elements[child]){
			child++
		}

		if(lastElement > h.elements[child]){
			h.elements[i] = h.elements[child]
		}else{
			break
		}
	}
	h.elements[i] = lastElement

	return minElement
}

func findMin(h pQueue)(int){
	return h.elements[1]
}

func isEmpty(h pQueue)(bool){
	if h.elements == nil && h != nil{
		return true
	}else{
		return false
	}
}

func isFull(h pQueue)(bool){
	if isEmpty(h)==true{
		if h.elements[h.capacity-1] != 0 {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}


func main(){
	var t pQueue
	fmt.Println("生成可存储64个数的最小堆")
	t=initHeap(64)

	r:=rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0;i<63;i++{
		insert(r.Intn(100)+5,t)
	}

	fmt.Printf("第1个数%d\n",t.elements[0])
	fmt.Printf("第32个数%d\n",t.elements[31])
	fmt.Printf("最小值为%d\n",findMin(t))
	deleteMin(t)
	fmt.Printf("最小值为%d\n",findMin(t))
 
}