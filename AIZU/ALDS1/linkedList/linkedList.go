package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	key  int
	next *Node
	prev *Node
}

var nilNode = &Node{-1, nil, nil}

func (l *Node) PrintList() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	isf := 1
	for cur := l.prev; cur != nilNode; cur = cur.prev {
		if isf == 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, cur.key)
		isf = 0
	}
	fmt.Fprintln(w, "")
}

func (l *Node) Insert(key int) {
	newNode := &Node{key, nilNode, nilNode}
	if l.prev == nil {
		l.next = newNode
		l.prev = newNode
		return
	}
	newNode.next = l
	newNode.prev = l.prev
	l.prev.next = newNode
	l.prev = newNode
}

func (l *Node) ListSearch(key int) (*Node, bool) {
	for p := l.prev; p != nilNode; p = p.prev {
		if p.key == key {
			return p, true
		}
	}
	return nil, false
}

func (l *Node) Delete(node *Node) {

	if node.prev != nilNode {
		node.prev.next = node.next
	} else {
		l.next = node.next
	}

	if node.next != nilNode {
		node.next.prev = node.prev
	} else {
		l.prev = node.prev
	}
}

func (l *Node) DeleteFirst() {
	l.Delete(l.prev)
}

func (l *Node) DeleteLast() {
	l.Delete(l.next)
}

func (l *Node) RemoveItem(key int) {
	node, ok := l.ListSearch(key)
	if !ok {
		return
	}
	l.Delete(node)
}

const (
	initialBufSize = 3e7
	maxBufSize     = 3e7
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func readInt() (read int) {
	scanner.Scan()
	read, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return
}

func readRunes() (read []rune) {
	scanner.Scan()
	for _, v := range scanner.Text() {
		read = append(read, v)
	}
	return
}

func main() {
	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)
	scanner.Split(bufio.ScanWords)

	n := readInt()

	for i := 0; i < n; i++ {
		com := readRunes()
		if com[0] == 'i' {
			key := readInt()
			nilNode.Insert(key)
		} else if com[0] == 'd' {
			if len(com) > 6 {
				if com[6] == 'F' {
					nilNode.DeleteFirst()
				} else if com[6] == 'L' {
					nilNode.DeleteLast()
				}
			} else {
				key := readInt()
				nilNode.RemoveItem(key)
			}
		}
	}

	nilNode.PrintList()
}
