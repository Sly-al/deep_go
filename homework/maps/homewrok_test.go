package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Node struct {
	val         int
	left, right *Node
}

func insertBST(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			val: val,
		}
	}
	if root.val > val {
		root.left = insertBST(root.left, val)
	} else {
		root.right = insertBST(root.right, val)
	}
	return root
}

func deleteBST(root *Node, val int) *Node {
	if root == nil {
		return root
	}
	if val < root.val {
		root.left = deleteBST(root.left, val)
	} else if val > root.val {
		root.right = deleteBST(root.right, val)
	} else if root.left != nil && root.right != nil {
		root.val = minimumBST(root.right)
		root.right = deleteBST(root.right, root.val)
	} else {
		if root.left != nil {
			root = root.left
		} else if root.right != nil {
			root = root.right
		} else {
			root = nil
		}
	}
	return root
}

func minimumBST(root *Node) int {
	if root.left == nil {
		return root.val
	}
	return minimumBST(root.left)
}

type OrderedMap struct {
	dict map[int]int
	root *Node
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		dict: make(map[int]int),
	}
}

func (m *OrderedMap) Insert(key, value int) {
	if _, ok := m.dict[key]; ok {
		return
	}
	m.dict[key] = value

	m.root = insertBST(m.root, key)
}

func (m *OrderedMap) Erase(key int) {
	if _, ok := m.dict[key]; !ok {
		return
	}

	delete(m.dict, key)
	m.root = deleteBST(m.root, key)
}

func (m *OrderedMap) Contains(key int) bool {
	_, ok := m.dict[key]
	return ok
}

func (m *OrderedMap) Size() int {
	return len(m.dict)
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.forEach(m.root, action)
}

func (m *OrderedMap) forEach(root *Node, action func(int, int)) {
	if root != nil {
		m.forEach(root.left, action)
		action(root.val, m.dict[root.val])
		m.forEach(root.right, action)
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
