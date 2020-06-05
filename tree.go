//pending refactor to select statements instead of if/else statements
package main

//structs
type Tree struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

//methods
func (n *Node) Insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(node)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

func (n *Node) exists(value int) bool {
	if n.value != value {
		return false
	}
	if n.value == value {
		return true
	}
	if n.value <= value {
		return n.value.exists(value)
	} else {
		return n.right.exists(value)
	}
}

func (n *Node) printNode() {
	if n == 0 {
		return
	}
	println(t.node)

	printNode(n.left)
	printNode(n.right)
}

func main() {
	var t tree

	t.insert(1).
		insert(2).
		insert(3).
		insert(4).
		insert(5).
		insert(6)

	println(t.exists.node(5))
}
