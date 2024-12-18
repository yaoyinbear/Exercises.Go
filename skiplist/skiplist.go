package skiplist

import (
	"math/rand"
)

type Key interface{}

type Value interface{}

type LessFunc func(Key, Key) bool

type skipLevel struct {
	next *skipNode
	span uint
}

type skipNode struct {
	key    Key
	value  Value
	levels []skipLevel
	prev   *skipNode
}

type SkipList struct {
	head     *skipNode
	tail     *skipNode
	level    uint
	count    uint
	lessFunc LessFunc
}

type updateNodes []*skipNode
type ranks []uint

const maxLevel = 32
const t = 25
const tBase = 100

func genLevel() uint {
	lv := 1
	for ; lv < maxLevel && rand.Intn(tBase) <= t; lv++ {
	}
	return uint(lv)
}

func newNode(level uint) *skipNode {
	return &skipNode{levels: make([]skipLevel, level)}
}

func NewSkipList(lessFunc LessFunc) *SkipList {
	skipList := SkipList{
		head:     newNode(maxLevel),
		level:    1,
		lessFunc: lessFunc,
	}

	return &skipList
}

func (skipList *SkipList) keyEqual(key1, key2 Key) bool {
	return !skipList.lessFunc(key1, key2) && !skipList.lessFunc(key2, key1)
}

func (skipList *SkipList) findLastLessThan(key Key) (*skipNode, updateNodes, ranks) {
	curNode := skipList.head
	updateNodes := make(updateNodes, skipList.level)
	ranks := make(ranks, skipList.level)
	for lv := skipList.level; lv > 0; lv-- {
		if lv < skipList.level {
			ranks[lv-1] = ranks[lv]
		}
		for nextNode := curNode.levels[lv-1].next; nextNode != nil && skipList.lessFunc(nextNode.key, key); curNode, nextNode = nextNode, curNode.levels[lv-1].next {
			ranks[lv-1] += curNode.levels[lv-1].span
		}
		updateNodes[lv-1] = curNode
	}

	return curNode, updateNodes, ranks
}

func (skiplist *SkipList) Find(key Key) *Value {
	node, _, _ := skiplist.findLastLessThan(key)
	nextNode := node.levels[0].next
	if nextNode != nil && skiplist.keyEqual(key, nextNode.key) {
		return &nextNode.value
	}

	return nil
}

func (skiplist *SkipList) Insert(key Key, value Value) bool {
	node, updateNodes, ranks := skiplist.findLastLessThan(key)
	nextNode := node.levels[0].next
	if nextNode != nil && skiplist.keyEqual(key, nextNode.key) {
		return false
	}

	newLv := genLevel()
	newNode := newNode(newLv)
	newNode.key = key
	newNode.value = value

	for i := skiplist.level + 1; i <= newLv; i++ {
		skiplist.head.levels[i-1].next = newNode
		skiplist.head.levels[i-1].span = ranks[0] + 1
		newNode.levels[i-1].span = skiplist.count - ranks[0]
	}

	for i := newLv + 1; i <= skiplist.level; i++ {
		updateNodes[i-1].levels[i-1].span++
	}

	for i := min(newLv, skiplist.level); i > 0; i-- {
		newNode.levels[i-1].next = updateNodes[i-1].levels[i-1].next
		updateNodes[i-1].levels[i-1].next = newNode

		newNode.levels[i-1].span = updateNodes[i-1].levels[i-1].span - (ranks[0] - ranks[i-1])
		updateNodes[i-1].levels[i-1].span = ranks[0] - ranks[i-1] + 1
	}

	if newNode.levels[0].next != nil {
		newNode.levels[0].next.prev = newNode
	} else {
		skiplist.tail = newNode
	}

	if node != skiplist.head {
		newNode.prev = node
	}

	skiplist.count++
	if newLv > skiplist.level {
		skiplist.level = newLv
	}

	return true
}

func (skiplist *SkipList) Remove(key Key) bool {
	node, updateNodes, _ := skiplist.findLastLessThan(key)
	removeNode := node.levels[0].next
	if removeNode == nil || !skiplist.keyEqual(key, removeNode.key) {
		return false
	}

	for i := skiplist.level; i > 0; i-- {
		if removeNode == updateNodes[i-1].levels[i-1].next {
			updateNodes[i-1].levels[i-1].next = removeNode.levels[i-1].next
			updateNodes[i-1].levels[i-1].span += removeNode.levels[i-1].span - 1
		} else {
			updateNodes[i-1].levels[i-1].span--
		}
	}

	if removeNode.levels[0].next != nil {
		removeNode.levels[0].next.prev = removeNode.prev
	} else {
		skiplist.tail = removeNode.prev
	}

	for ; skiplist.level > 1 && skiplist.head.levels[skiplist.level-1].next == nil; skiplist.level-- {
		skiplist.head.levels[skiplist.level-1].span = 0
	}

	skiplist.count--

	// removeNode.key = nil
	// removeNode.value = nil
	// clear(removeNode.levels)
	// removeNode.prev = nil

	return true
}

func (skiplist *SkipList) Range(fn func(key Key, value Value) bool) {
	for node := skiplist.head.levels[0].next; node != nil; node = node.levels[0].next {
		if !fn(node.key, node.value) {
			return
		}
	}
}
