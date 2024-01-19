package geohash

import (
	"errors"
	"sync"
)

type (
	// Trie is a geohash coding prefix tree with a height fixed to geohashLen + 1.
	// Leaf nodes store longitude and latitude data, and non-leaf nodes store geohash coded indexes.
	Trie struct {
		root *node

		sync.RWMutex
	}

	node struct {
		children  [32]*node // base32
		passCount uint32    // the number of Box pass the node (leafNode.passCount = 0)

		isLeaf bool
		*Box   // only belongs to leaf node
	}
)

func NewTrie() *Trie {
	return &Trie{root: &node{}}
}

func (t *Trie) Get(geohash Geohash) (*Box, bool) {
	if t == nil || t.root == nil || !geohash.valid() {
		return nil, false
	}

	t.RLock()
	defer t.RUnlock()

	n := t.search(string(geohash))
	if n == nil || !n.isLeaf {
		return nil, false
	}
	return n.Box, true
}

func (t *Trie) GetBoxesByPrefix(prefix string) []*Box {
	if t == nil || t.root == nil || len(prefix) == 0 {
		return nil
	}

	t.RLock()
	defer t.RUnlock()

	n := t.search(prefix)
	if n == nil {
		return nil
	}
	if n.isLeaf {
		return []*Box{n.Box}
	}

	return n.dfs()
}

func (t *Trie) Put(point *Point) {
	if t == nil || t.root == nil || point == nil {
		return
	}

	geohash := point.Geohash()

	t.Lock()
	defer t.Unlock()

	if n := t.search(string(geohash)); n != nil && n.isLeaf {
		n.add(point)
		return
	}

	move := t.root
	for i := 0; i < geohashLen; i++ {
		childIndex := decode(geohash[i])
		if move.children[childIndex] == nil {
			move.children[childIndex] = &node{}
		}
		move.passCount++
		move = move.children[childIndex]
	}
	move.isLeaf = true
	move.Box = NewBox(geohash, map[string]*Point{point.key(): point})
}

func (t *Trie) Delete(geohash Geohash) bool {
	if t == nil || t.root == nil || !geohash.valid() {
		return false
	}

	t.Lock()
	defer t.Unlock()

	n := t.search(string(geohash))
	if n == nil || !n.isLeaf {
		return false
	}

	move := t.root
	for i := 0; i < geohashLen; i++ {
		index := decode(geohash[i])
		move.passCount--
		if move.passCount == 0 {
			move.children[index] = nil
			return true
		}
		move = move.children[index]
	}

	return false
}

func (t *Trie) GetPointsByCircle(center *Point, radius uint32) ([]*Point, error) {
	if t == nil || t.root == nil || center == nil || radius == 0 {
		return nil, errors.New("invalid param")
	}

	l, err := getGeohashLenByDiameter(radius << 1)
	if err != nil {
		return nil, err
	}

	points := center.circumscribedSquarePointsByCircle(radius)

	t.RLock()
	defer t.RUnlock()

	res := make([]*Point, 0)
	duplicateBox := map[Geohash]struct{}{}
	for _, p := range points {
		for _, box := range t.GetBoxesByPrefix(string(p.Geohash()[:l])) {
			if _, ok := duplicateBox[box.GetGeohash()]; !ok {
				for _, v := range box.GetAllPoints() {
					if center.Distance(v) <= radius {
						res = append(res, v)
					}
				}
				duplicateBox[box.GetGeohash()] = struct{}{}
			}
		}
	}

	return res, nil
}

func (t *Trie) Count() uint32 {
	if t == nil || t.root == nil {
		return 0
	}

	return t.root.passCount
}

func (t *Trie) search(prefix string) *node {
	if t == nil || t.root == nil || len(prefix) == 0 {
		return nil
	}

	move := t.root
	for i := 0; i < len(prefix); i++ {
		childIndex := decode(prefix[i])
		if childIndex == invalidCode || move.children[childIndex] == nil {
			return nil
		}
		move = move.children[childIndex]
	}
	return move
}

// dfs returns []*Box through the node
func (n *node) dfs() []*Box {
	if n == nil {
		return nil
	}

	if n.isLeaf {
		return []*Box{n.Box}
	}
	if n.passCount == 0 {
		return []*Box{}
	}

	res := make([]*Box, 0, n.passCount)
	for i := 0; i < len(n.children); i++ {
		if n.children[i] != nil {
			res = append(res, n.children[i].dfs()...)
		}
	}

	return res
}
