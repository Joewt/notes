package main

import (
	"container/list"
	"fmt"
)

/*
链表节点的结构
type Element struct {
	next, prev *Element
	list *list
	Value interface{}
}
*/

func main() {
	l := list.New()
	l.Init() // 清空链表
	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")
	fmt.Println(l.Len())
	v := l.Back() //返回最后一个节点
	fmt.Println(v.Value)
	v = l.Front() //第一个节点
	fmt.Println(v.Value)

}
