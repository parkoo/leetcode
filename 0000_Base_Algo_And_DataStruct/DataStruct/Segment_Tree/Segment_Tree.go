package main

import (
	"fmt"
	"strings"
)

/**
	线段树:
	1. 对于一个固定区间 data[0, len(data)-1], 可以构建一棵对应的 SegmentTree，SegmentTree[index] 存储一段区间data[l, r]的统计量 (区间最大值、区间和...);
	2. 对于任意一个区间 data[l, r] 的统计量, 可以用 logn 的时间复杂度找到对应的 SegmentTree[index];
	3. 线段树解决的问题: 区间固定，区间内的数据会发生变化，对区间内的数据进行更新和查询的操作.

	数组实现:
	1. 对于满二叉树，有一个特性：最后一层的节点数大致等于前面所有层节点之和
	2. 线段树不是完全二叉树，但它是一种平衡二叉树，将线段树填补成满二叉树 (不存在的节点存储“nil”), 则可以用数组存储 （完全二叉树可以用数组存储，满二叉树一定是完全二叉树）
	3. 线段树通常不考虑添加元素，如果区间中有n个元素，则相应的存储线段树的数组最多需要4n个静态存储空间（最坏情况：n=2^k+1, 此时，浪费了近一半的空间(存储“nil”)）
**/

// MergeFn 融合函数，用于融合统计数据
type MergeFn func(a, b interface{}) interface{}

type SegmentTree struct {
	merger MergeFn
	data   []interface{}
	tree   []interface{}
}

func NewSegmentTree(data []interface{}, merger MergeFn) *SegmentTree {
	segmentTree := &SegmentTree{
		merger: merger,
		data:   data,
		tree:   make([]interface{}, 4*len(data)),
	}

	buildSegmentTree(segmentTree, 0, 0, len(data)-1)
	return segmentTree
}

// buildSegmentTree 递归构建线段树，在线段树数组中索引为index处，存放数据数组中区间为[l, r]的数据的统计量
func buildSegmentTree(st *SegmentTree, index, l, r int) {
	if l == r {
		st.tree[index] = st.data[l]
		return
	}

	mid := l + (r-l)/2
	buildSegmentTree(st, getLeftChild(index), l, mid)
	buildSegmentTree(st, getRightChild(index), mid+1, r)

	// 根据具体业务，将左右节点的统计量融合成根节点的统计量
	st.tree[index] = st.merger(st.tree[getLeftChild(index)], st.tree[getRightChild(index)])
}

// getLeftChild 获取左子节点索引
func getLeftChild(index int) int {
	return 2*index + 1
}

// getRightChild 获取右子节点索引
func getRightChild(index int) int {
	return 2 * (index + 1)
}

// Query 查询区间统计量
func (st *SegmentTree) Query(ql, qr int) interface{} {
	if ql < 0 || ql >= len(st.data) || qr < 0 || qr >= len(st.data) || ql > qr {
		panic("index is illegal")
	}

	return query(st, 0, 0, len(st.data)-1, ql, qr)
}

// query 递归查询区间 data[ql, qr] 的统计值
// 从以index为根的节点（此节点存储着[l,r]数据的统计量）搜索存储着区间[ql, qr]的统计量的节点
func query(st *SegmentTree, index, l, r, ql, qr int) interface{} {
	if l == qr && r == qr {
		return st.tree[index]
	}

	mid := l + (r-l)/2
	if ql >= mid+1 {
		return query(st, getRightChild(index), mid+1, r, ql, qr)
	} else if qr <= mid {
		return query(st, getLeftChild(index), l, mid, ql, qr)
	} else {
		leftRes := query(st, getLeftChild(index), l, mid, ql, qr)
		rightRes := query(st, getRightChild(index), mid+1, r, ql, qr)
		return st.merger(leftRes, rightRes)
	}
}

// Set 更新区间位置为i的数据，同步更新区间树相应节点的统计量
func (st *SegmentTree) Set(i int, v interface{}) {
	if i < 0 || i >= len(st.data) {
		panic("index is illegal")
	}

	set(st, 0, 0, len(st.data)-1, i, v)
}

// set 递归更新线段树节点的统计量
func set(st *SegmentTree, index, l, r, i int, v interface{}) {
	if l == r {
		st.tree[index] = v
		return
	}

	mid := l + (r-l)/2
	if i <= mid {
		set(st, getLeftChild(index), l, mid, i, v)
	} else {
		set(st, getRightChild(index), mid+1, r, i, v)
	}
	st.tree[index] = st.merger(st.tree[getLeftChild(index)], st.tree[getRightChild(index)])
}

// Get 获取数据
func (st *SegmentTree) Get(i int) interface{} {
	if i < 0 || i >= len(st.data) {
		panic("index is illegal")
	}

	return st.data[i]
}

// PrintST 打印区间树
func PrintST(st *SegmentTree) string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for i, v := range st.tree {
		if v != nil {
			sb.WriteString(v.(string))
		} else {
			sb.WriteString("nil")
		}
		if i < len(st.tree)-1 {
			sb.WriteString("; ")
		}
	}
	sb.WriteString("]")

	return sb.String()
}

// 验证
func main() {
	data := []interface{}{"1", "2", "3", "4", "5"}

	st := NewSegmentTree(data, func(a, b interface{}) interface{} {
		return a.(string) + "-" + b.(string)
	})

	fmt.Println(PrintST(st))
}
