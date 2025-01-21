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
	PrintTree()
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
		"Print Tree",
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

const MaxTreeSize = 20

func (t *BinaryTree) Insert(data int) {
	if t.size >= MaxTreeSize {
		fmt.Println("Tree is full")
		return
	}

	if data < 0 {
		fmt.Println("Invalid data")
		return
	}

	newNode := &Node{data: data}

	if t.root == nil {
		t.root = newNode
		t.size++
		return
	}

	current := t.root
	for {
		if data == current.data {
			fmt.Println("Duplicate data entered")
			return
		}

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
	return t.balance(node)
}

func (t *BinaryTree) balance(node *Node) *Node {
	if node == nil {
		return nil
	}

	leftHeight := t.getHeight(node.left)
	rightHeight := t.getHeight(node.right)

	if abs(leftHeight-rightHeight) <= 1 {
		return node
	}

	if leftHeight > rightHeight {
		if t.getHeight(node.left.left) >= t.getHeight(node.left.right) {
			return t.rotateRight(node)
		}
		node.left = t.rotateLeft(node.left)
		return t.rotateRight(node)
	}

	if t.getHeight(node.right.right) >= t.getHeight(node.right.left) {
		return t.rotateLeft(node)
	}
	node.right = t.rotateRight(node.right)
	return t.rotateLeft(node)
}

func (t *BinaryTree) getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := t.getHeight(node.left)
	rightHeight := t.getHeight(node.right)
	return max(leftHeight, rightHeight) + 1
}

func (t *BinaryTree) rotateLeft(node *Node) *Node {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node
	return newRoot
}

func (t *BinaryTree) rotateRight(node *Node) *Node {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node
	return newRoot
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func (t *BinaryTree) PrintTree() {
	if t.root == nil {
		return
	}

	queue := []*Node{t.root}
	level := 0
	nodesInCurrentLevel := 1
	nodesInNextLevel := 0

	for len(queue) > 0 {
		fmt.Printf("Level %d: ", level)
		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == t.root {
				fmt.Printf("%d(root)[L:", node.data)
				if node.left != nil {
					fmt.Printf("%d", node.left.data)
				} else {
					fmt.Printf("nil")
				}
				fmt.Printf(",R:")
				if node.right != nil {
					fmt.Printf("%d", node.right.data)
				} else {
					fmt.Printf("nil")
				}
				fmt.Printf("] ")
			} else {
				fmt.Printf("%d[L:", node.data)
				if node.left != nil {
					fmt.Printf("%d", node.left.data)
				} else {
					fmt.Printf("nil")
				}
				fmt.Printf(",R:")
				if node.right != nil {
					fmt.Printf("%d", node.right.data)
				} else {
					fmt.Printf("nil")
				}
				fmt.Printf("] ")
			}

			if node.left != nil {
				queue = append(queue, node.left)
				nodesInNextLevel++
			}
			if node.right != nil {
				queue = append(queue, node.right)
				nodesInNextLevel++
			}
		}
		fmt.Println()

		level++
		nodesInCurrentLevel = nodesInNextLevel
		nodesInNextLevel = 0
	}
}

var _ IBinaryTree = (*BinaryTree)(nil)
