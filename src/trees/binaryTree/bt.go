package binaryTree

import "fmt"

type IBinaryTree interface {
	Insert(data int)
	Delete(data int)
	Search(data int) bool
	InorderTraversal()
	PreorderTraversal()
	PostorderTraversal()
	IsEmpty() bool
	Size() int
}

func GetChoices() []string {
	return []string{
		"Insert",
		"Delete",
		"Search",
		"Inorder Traversal",
		"Preorder Traversal",
		"Postorder Traversal",
		"Is empty",
		"Size",
	}
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
	size int
}

func (t *BinaryTree) Insert(data int) {
	newNode := &Node{data: data}
	if t.root == nil {
		t.root = newNode
		t.size++
		return
	}
	current := t.root
	for {
		if data < current.data {
			if current.left == nil {
				current.left = newNode
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				break
			}
			current = current.right
		}
	}
	t.size++
}

func (t *BinaryTree) Delete(data int) {
	t.root = t.deleteNode(t.root, data)
}

func (t *BinaryTree) deleteNode(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.data {
		node.left = t.deleteNode(node.left, data)
	} else if data > node.data {
		node.right = t.deleteNode(node.right, data)
	} else {
		if node.left == nil {
			t.size--
			return node.right
		} else if node.right == nil {
			t.size--
			return node.left
		}
		minNode := t.findMin(node.right)
		node.data = minNode.data
		node.right = t.deleteNode(node.right, minNode.data)
	}
	return node
}

func (t *BinaryTree) findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *BinaryTree) Search(data int) bool {
	return t.searchNode(t.root, data)
}

func (t *BinaryTree) searchNode(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if node.data == data {
		return true
	}
	if data < node.data {
		return t.searchNode(node.left, data)
	}
	return t.searchNode(node.right, data)
}

func (t *BinaryTree) InorderTraversal() {
	t.inorder(t.root)
}

func (t *BinaryTree) inorder(node *Node) {
	if node != nil {
		t.inorder(node.left)
		fmt.Printf("%d ", node.data)
		t.inorder(node.right)
	}
}

func (t *BinaryTree) PreorderTraversal() {
	t.preorder(t.root)
}

func (t *BinaryTree) preorder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.data)
		t.preorder(node.left)
		t.preorder(node.right)
	}
}

func (t *BinaryTree) PostorderTraversal() {
	t.postorder(t.root)
}

func (t *BinaryTree) postorder(node *Node) {
	if node != nil {
		t.postorder(node.left)
		t.postorder(node.right)
		fmt.Printf("%d ", node.data)
	}
}

func (t *BinaryTree) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree) Size() int {
	return t.size
}

var _ IBinaryTree = (*BinaryTree)(nil)
