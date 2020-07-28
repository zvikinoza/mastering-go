package bst

import "sync"

type Set interface {
	Add(item Item) *BST
	Delete(item Item) bool
	Has(item Item) bool
	Clear()
	Items() []Item
	Size() int
}

type Type interface{}

type Item Type

type BST struct {
	root       *node
	lock       sync.RWMutex
	size       int
	comparator Comparator
}

//
// negative , if a < b
// zero     , if a == b
// positive , if a > b
//
type Comparator func(a, b Item) int

type node struct {
	value Item
	left  *node
	right *node
}

func (bst *BST) Add(item Item) *BST {
	var add func(root *node, item Item) *node

	add = func(root *node, item Item) *node {
		if root == nil {
			bst.size++
			return &node{value: item}
		}

		if bst.comparator(item, root.value) < 0 {
			root.left = add(root.left, item)
		} else if bst.comparator(item, root.value) > 0 {
			root.right = add(root.right, item)
		}

		return root
	}

	bst.lock.Lock()
	defer bst.lock.Unlock()

	bst.root = add(bst.root, item)
	return bst
}

func (bst *BST) Delete(item Item) bool {
	found := false

	var min func(n *node) Item
	min = func(n *node) Item {
		if n.left == nil {
			return n.value
		}
		return min(n.left)
	}

	var delete func(root *node, item Item) *node
	delete = func(root *node, item Item) *node {
		if root == nil {
			return nil
		}

		if bst.comparator(item, root.value) < 0 {
			root.left = delete(root.left, item)
		} else if bst.comparator(item, root.value) > 0 {
			root.right = delete(root.right, item)
		} else {

			found = true
			bst.size--

			if root.left == nil {
				return root.right
			} else if root.right == nil {
				return root.left
			} else {
				minItem := min(root.right)
				root.value = minItem
				root.right = delete(root.right, minItem)
			}
		}
		return root
	}

	bst.lock.Lock()
	defer bst.lock.Unlock()

	bst.root = delete(bst.root, item)
	return found
}

func (bst *BST) Has(item Item) bool {
	var has func(root *node, item Item) bool
	has = func(root *node, item Item) bool {
		if root == nil {
			return false
		}

		if bst.comparator(item, root.value) < 0 {
			return has(root.left, item)
		} else if bst.comparator(item, root.value) > 0 {
			return has(root.right, item)
		}
		return true
	}

	bst.lock.RLock()
	defer bst.lock.RUnlock()

	return has(bst.root, item)
}

func (bst *BST) Clear() {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	bst.root = nil
	bst.size = 0
}

// inorder traversal
func (bst *BST) Items() []Item {
	items := []Item{}

	var inorder func(root *node)
	inorder = func(root *node) {
		if root != nil {
			inorder(root.left)
			items = append(items, root.value)
			inorder(root.right)
		}
	}

	bst.lock.RLock()
	defer bst.lock.RUnlock()

	inorder(bst.root)
	return items
}

func (bst *BST) Size() int {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	return bst.size
}
