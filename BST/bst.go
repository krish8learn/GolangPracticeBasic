package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := &node{20, nil, nil}
	n.left = &node{10, nil, nil}
	n.right = &node{23, nil, nil}
	b := &bst{
		root: n,
		len:  3,
	}
	b.add(45)
	b.add(21)
	b.add(47)
	b.add(66)
	fmt.Println(b.search(44))
	fmt.Println(b.search(21))
	b.remove(5)
	b.remove(47)
	fmt.Println(*b)

}

//creating node (element of BST)
type node struct {
	value int
	left  *node
	right *node
}

//it help us to access starting point of BST and length
type bst struct {
	root *node
	len  int
}

//get the value node
func (n node) toString() string {
	return strconv.Itoa(n.value)
}

//get the value from starting node of BST
func (b bst) startString() string {
	sb := strings.Builder{}
	b.inoderTraversal(&sb)
	return sb.String()
}

//inorder traversal
func (b bst) inoderTraversal(sb *strings.Builder) {
	b.inorderTraversalByNode(sb, b.root)
}

func (b bst) inorderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inorderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inorderTraversalByNode(sb, root.right)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if root.value > value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.right
			for temp.right != nil {
				temp = temp.right
			}

			root.value = temp.value

			root.left = b.removeByNode(root.left, value)
		}
	}

	return nil
}
