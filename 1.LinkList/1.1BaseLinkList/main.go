package main

import (
	"fmt"
	"strconv"
)

// 链表结构
type LinkList struct {
	head *LinkListNode
}

func NewLinkList() *LinkList {
	return &LinkList{
		head: &LinkListNode{
			data: "head",
			next: nil,
		},
	}
}

// 添加链表节点
// 1.先循环到最末尾节点
// 2.链接节点
// 注意 最末尾一个节点的next为nil
func (l *LinkList) AddNode(item *LinkListNode) {
	nowNode := l.head
	for {
		if nowNode.next != nil {
			nowNode = nowNode.next
		} else {
			nowNode.next = item
			break
		}
	}

}

// 删除链表节点
func (l *LinkList) DeleteNode(data string) {
	nowNode := l.head
	var prevNode *LinkListNode
	prevNode = nil
	for {
		if nowNode != nil {
			if prevNode == nil {
				// 如果第一次循环 不能删除头节点
				prevNode = nowNode
				// 删除头节点
				if nowNode.data == data {
					l.head = nowNode.next
					continue
				}
				nowNode = nowNode.next
				continue
			}

			// 中间节点尾节点判断
			if nowNode.data == data {
				prevNode.next = nowNode.next
				nowNode = nil
				break
			}
			prevNode = nowNode
			nowNode = nowNode.next
			continue
		}
		break
	}
}

func (l *LinkList) Print() {
	nowNode := l.head
	for {
		if nowNode != nil {
			fmt.Println(nowNode.data)
			nowNode = nowNode.next
			continue
		}
		break
	}
}

// 链表节点
type LinkListNode struct {
	next *LinkListNode
	data string
}

func NewLinkListNode(data string) *LinkListNode {
	return &LinkListNode{
		next: nil,
		data: data,
	}
}

func main() {

	linkList := NewLinkList()

	// 添加数据
	for i := 0; i < 10; i++ {
		newData := NewLinkListNode(strconv.Itoa(i))
		linkList.AddNode(newData)
	}
	// linkList.DeleteNode("3")
	linkList.DeleteNode("head")
	linkList.DeleteNode("9")
	linkList.Print()
}
