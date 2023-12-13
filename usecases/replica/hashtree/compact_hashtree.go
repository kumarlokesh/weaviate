//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package hashtree

type CompactHashTree struct {
	capacity int
	hashtree *HashTree

	// derived values from capacity and hashtree height
	// kept here just to avoid recalculation
	leavesCount              int
	normalGroupSize          int
	largeGroupSize           int
	largeGroupsCount         int
	leavesCountInLargeGroups int
}

func NewCompactHashTree(capacity int, maxHeight int) *CompactHashTree {
	if capacity < 1 {
		panic("illegal capacity")
	}

	if maxHeight < 1 {
		panic("illegal max height")
	}

	height := requiredHeight(capacity)

	if height > maxHeight {
		height = maxHeight
	}

	leavesCount := LeavesCount(height)
	normalGroupSize := capacity / leavesCount
	largeGroupSize := normalGroupSize + 1
	largeGroupsCount := capacity % leavesCount
	leavesCountInLargeGroups := largeGroupsCount * largeGroupSize

	return &CompactHashTree{
		capacity: capacity,
		hashtree: NewHashTree(height),

		leavesCount:              leavesCount,
		normalGroupSize:          normalGroupSize,
		largeGroupSize:           largeGroupSize,
		largeGroupsCount:         largeGroupsCount,
		leavesCountInLargeGroups: leavesCountInLargeGroups,
	}
}

func (ht *CompactHashTree) Height() int {
	return ht.hashtree.Height()
}

// AggregateLeafWith aggregates a new value into a shared leaf
// Each compacted leaf is shared by a number of consecutive leaves
func (ht *CompactHashTree) AggregateLeafWith(i int, val []byte) *CompactHashTree {
	if i >= ht.capacity {
		panic("out of capacity")
	}

	var mappedLeaf int

	if i < ht.leavesCountInLargeGroups {
		mappedLeaf = i / ht.largeGroupSize
	} else {
		mappedLeaf = (i - ht.largeGroupsCount) / ht.normalGroupSize
	}

	ht.hashtree.AggregateLeafWith(mappedLeaf, val)

	return ht
}

func (ht *CompactHashTree) Level(level int, discriminant *Bitset, digests []Digest) (n int, err error) {
	return ht.hashtree.Level(level, discriminant, digests)
}

func (ht *CompactHashTree) Reset() *CompactHashTree {
	ht.hashtree.Reset()
	return ht
}

func requiredHeight(n int) int {
	h := 0

	for ; n > 0; h++ {
		n = n >> 1
	}

	return h
}
