package main

import "log"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

// AddToHead Gắn một node vào đầu LinkedList
func(list *LinkedList) AddToHead(n int) {
	// Tạo node mới
	var node = Node{}
	node.property = n

	// headNode hiện tại của LinkedList giờ trở thành nextNode của node mới tạo
	// Nếu ban đầu LinkedList chưa có gì --> list.headNode = nil
	node.nextNode = list.headNode

	// log.Println("next node is", node.nextNode)

	// node mới tạo giờ trở thành headNode mới của LinkedList
	list.headNode = &node
}

func(list *LinkedList) IterateList() {
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

func(list *LinkedList) LastNode() *Node {
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

func(list *LinkedList) AddToEnd(n int) {
	var node = Node{}
	node.property = n
	// node này sẽ là node cuối trong LinkedList
	node.nextNode = nil

	var lastNode = list.LastNode()
	lastNode.nextNode = &node
}

func(list *LinkedList) FindNodeWithValue(n int) *Node {
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

// AddNodeAfter Chèn một node vào Linked List
func(list *LinkedList) AddNodeAfter(oldNodeProperty, newNodeProperty int) {
	var nodeToBeAdded = Node{}
	nodeToBeAdded.property = newNodeProperty

	var foundNode = list.FindNodeWithValue(oldNodeProperty)
	var afterFoundNode = foundNode.nextNode

	nodeToBeAdded.nextNode = afterFoundNode
	foundNode.nextNode = &nodeToBeAdded
}

func main() {
	var result = LinkedList{}
	result.AddToHead(5)
	result.AddToHead(4)
	result.AddToHead(3)
	result.AddToHead(2)
	result.AddToHead(1)

	
	log.Println("linked list before")
	result.IterateList()

	result.AddNodeAfter(5, 10)

	log.Println("last node after")
	result.IterateList()

}