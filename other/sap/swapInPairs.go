package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	inputNode := parseNode(input)
	outputNode := swapPairs(inputNode)
	output := printNode(outputNode)
	fmt.Println(output)
}
type ListNode struct {
	Val int
	Next *ListNode
}

// init a string to node.
func parseNode(input string) *ListNode {
	head := &ListNode{}
	p := head
	for _,val := range strings.Split(input, " ") {
		node := &ListNode{}
		node.Val,_ = strconv.Atoi(val)
		node.Next = nil
		p.Next = node
		p = p.Next
	}
	return head.Next
}

// main algorithm
func swapPairs(head *ListNode) *ListNode {
	// abnormal input check.
	if head == nil {
		return head
	}

	// single link list.
	if head.Next == nil {
		return head
	}

	newHead := head
	node1 := head
	node2 := head.Next

	// init the return head node
	// currently node1 is ahead of node2
	node1.Next = node2.Next
	node2.Next = node1
	newHead = node2

	// swap every two nodes in pairs, until reach the odd rest or end.
	for ; (node1.Next != nil && node1.Next.Next != nil); {
		// move to Next pair, keeping node2 ahead of node1.
		temp := node1
		node1 = temp.Next
		node2 = node1.Next

		// swap the pair like before, now node1 ahead of node2.
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
	}

	return newHead
}

// print link list in string.
func printNode(head *ListNode) string {
	// TODO: write your code here
	var output = ""

	// abnormal check.
	if head == nil {
		return output
	}

	current := head

	// make the first element is not blank compared with below.
	output += strconv.Itoa(current.Val)
	current = current.Next

	for ;current != nil; {
		output += " " + strconv.Itoa(current.Val)
		current = current.Next
	}
	return output
}