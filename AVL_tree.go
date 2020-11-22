package main

import (
	"fmt"
)

func main(){

	tree := &AVLNode{key: 100, Value: 400}
	fmt.Println(tree)
	//tree.Insert(200)
	//tree.Insert(300)
	//fmt.Println(tree)

	tree.add(110, 400)
	tree.add(96, 300)
	tree.add(19, 400)
	tree.add(76, 200)
	tree.add(150, 700)
	tree.add(310, 500)
	tree.add(7, 400)
	tree.add(24, 390)
	tree.add(88, 410)
	tree.add(276, 440)



	fmt.Println(tree)

	fmt.Println(tree.search(76))

	

}

	
}



// AVLTree structure. Public methods are Add, Remove, Update, Search, DisplayTreeInOrder.
type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(key int, value int) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree) Remove(key int) {
	t.root.remove(key)
}

func (t *AVLTree) Update(oldKey int, newKey int, newValue int) {
	t.root.remove(oldKey)
	t.root = t.root.add(newKey, newValue)
}

func (t *AVLTree) Search(key int) (node *AVLNode) {
	return t.root.search(key)
}

func (t *AVLTree) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

// AVLNode structure
type AVLNode struct {
	key   int
	Value int

	// height counts nodes (not edges)
	height int
	left   *AVLNode
	right  *AVLNode
}

// Adds a new node
func (n *AVLNode) add(key int, value int) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		// if same key exists update value
		n.Value = value
	}
	return n.rebalanceTree()
}

// Removes a node
func (n *AVLNode) remove(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			// delete smallest node that we replaced
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

// Searches for a node
func (n *AVLNode) search(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

// Displays nodes left-depth first (used for debugging)
func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// Checks if node is balanced and rebalance
func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Finds the smallest child (based on the key) for the current node
func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

// Returns max number - TODO: std lib seemed to only have a method for floats!
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
