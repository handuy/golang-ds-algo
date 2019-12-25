package main

import (
	"log"
)

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type DoublyLinkedList struct {
	headNode *Node
}

func (list *DoublyLinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	// Chạy vòng lặp đến node cuối cùng thì break
	for node = list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
			break
		}
	}

	return lastNode
}

func (list *DoublyLinkedList) FindNodeWithValue(n int) *Node {
	var node *Node
	var found *Node

	// Chạy vòng lặp cho đến khi tìm ra node có property = n
	for node = list.headNode; node != nil; node = node.nextNode {
		if node.property == n {
			found = node
			break
		}
	}

	return found
}

func (list *DoublyLinkedList) IterateList() {
	var node *Node

	// Ban đầu lấy ra headNode
	// Nếu headNode != nil thì in ra property, không thì thoát khỏi loop
	// In ra property xong thì chuyển sang node tiếp theo
	// Lại check node != nil
	// ...
	// Sau khi in ra property của node cuối cùng, chuyển sang node tiếp theo: node = (node cuối cùng).nextNode
	// Do là node cuối cùng nên node == nil --> thoát khỏi loop
	for node = list.headNode; node != nil; node = node.nextNode {
		log.Println(node.property)
	}
}

func (list *DoublyLinkedList) AddToHead(prop int) {
	var node Node
	node.property = prop

	// Nếu ban đầu trong Doubly Linked List đã có sẵn phần tử Head Node
	if list.headNode != nil {

		// Node mới thêm vào sẽ trở thành previousNode của Head Node hiện tại
		list.headNode.previousNode = &node

		// Head Node hiện tại sẽ trở thành nextNode của node mới thêm vào
		node.nextNode = list.headNode
	}

	// Nếu ban đầu trong Doubly Linked List chưa có node nào
	// thì set luôn node mới thêm vào thành Head Node
	list.headNode = &node
}

func (list *DoublyLinkedList) AddToEnd(prop int) {
	var node Node
	node.property = prop
	node.nextNode = nil

	var currentLastNode = list.LastNode()

	if currentLastNode != nil {
		currentLastNode.nextNode = &node
		node.previousNode = currentLastNode
	} else {
		// Không có last node, tức là list hiện đang rỗng không có node nào
		list.headNode = &node
	}
}

func (list *DoublyLinkedList) AddAfter(oldProp, newProp int) {
	var newNode = Node{}
	newNode.property = newProp

	var foundNode = list.FindNodeWithValue(oldProp)
	var afterFoundNode = foundNode.nextNode

	// Nếu node tìm thấy là node cuối cùng
	if afterFoundNode == nil {
		list.AddToEnd(newProp)
	} else {
		foundNode.nextNode = &newNode
		newNode.previousNode = foundNode

		newNode.nextNode = afterFoundNode
		afterFoundNode.previousNode = &newNode
	}
}

func main() {
	var list = DoublyLinkedList{}

	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)

	list.AddAfter(4, 0)

	list.IterateList()
}
