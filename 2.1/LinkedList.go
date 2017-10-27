package main

import (
	"fmt"
)

type Node struct{
	num int
	next *Node
}

type pNode *Node
type list pNode
type pos pNode


func isEmpty(l list)(bool){
	return l.next == nil
}

func isLast(p pos,l list)(bool){
	return p.next == nil
}

func findPosition(x int,l list)(int){
	var p pos
	num := 0

	p=l.next
	for p!=nil && p.num != x{
		p = p.next
		num++
	}

	return num
}

func findPrevious(x int,l list)(pos){
	var p pos

	p=pos(l)
	for p.next !=nil && (p.next).num != x{
		p=p.next
	}

	return p
}

func delete(x int,l list){
	var p,tmpCell pos
    p = findPrevious(x,l)
	if(!isLast(p,l)){
		tmpCell = p.next
		p.next = tmpCell.next
	}

}

func append(x int,l list){
	var n list
	var tmp pos
	n=l
	for n.next!=nil{
		n=n.next
	}

	tmp.num = x
	tmp.next = nil

	n.next = tmp

}

func getLength(l list)(int){
	var i int = 0
	n:=l
	for n.next!=nil{
		i++
		n=n.next
	}
	return i
}

func insert(x int,l list,p int)(){
	var tmp Node
	tmp.num=x

	if l.next==nil{
		tmp.next=nil
		l.next=&tmp
		return 
	}

	if(getLength(l)<p){
		fmt.Println("插入失败-指定长度超过链表")
		return 
	}

	var i int
	var n list
	
	n=l

	for i>0{
		n=n.next
		i--
	}

	tmp.next=n.next
	n.next=&tmp

}



func deleteList(l list){
	var p,tmp pos

	if isEmpty(l){
		return 
	}

	p = pos(l.next)
	l.next = nil
	for p!=nil{
		tmp = p.next
		p = nil
		p = tmp
	}

}

func main(){
	var test Node
	fun := &test
	j := 0

	for i:=10;i>0;i--{
		insert(i,&test,j)
		j++
	}

	fmt.Printf("打印链表组成: \n")
	for i:=0;i<10;i++{
		if i!=9{
			fmt.Printf("%d",fun.num)
			fmt.Printf("->")
		}else{
			fmt.Printf("%d\n",fun.num)
		}
		fun=fun.next
	}

	fmt.Println("链表的长度为", getLength(list(&test)))
	fmt.Println("数字6的位置在",findPosition(6,list(&test)))
	fmt.Println("删除链表...")
	deleteList(list(&test))
}