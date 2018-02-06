package main

import (
	"fmt"
	"strings"
)

func main() {
	n0001 := TNode{name: "NODE 1"}
	n0002 := TNode{name: "NODE 2"}
	n0003 := TNode{name: "NODE 3"}
	n0004 := TNode{name: "NODE 4"}
	n0005 := TNode{name: "NODE 5"}
	n0006 := TNode{name: "NODE 6"}
	n0007 := TNode{name: "NODE 7"}
	n0008 := TNode{name: "NODE 8"}
	n0009 := TNode{name: "NODE 9"}
	n0010 := TNode{name: "NODE 10"}
	n0011 := TNode{name: "NODE 11"}
	n0012 := TNode{name: "NODE 12"}
	n0013 := TNode{name: "NODE 13"}
	n0014 := TNode{name: "NODE 14"}
	n0015 := TNode{name: "NODE 15"}

	/*
		0001
			0002
			0003
				0006
				0007
				0008
					0009
			0004
			0005
				0010
				0011
					0013
					0014
					0015
				0012
	*/

	n0001.AddChild(&n0002)
	n0001.AddChild(&n0003)
	n0001.AddChild(&n0004)
	n0001.AddChild(&n0005)

	n0003.AddChild(&n0006)
	n0003.AddChild(&n0007)
	n0003.AddChild(&n0008)

	n0008.AddChild(&n0009)

	n0005.AddChild(&n0010)
	n0005.AddChild(&n0011)
	n0005.AddChild(&n0012)

	n0011.AddChild(&n0013)
	n0011.AddChild(&n0014)
	n0011.AddChild(&n0015)

	t := TTree{root: n0001}

	t.show()
}

type TTree struct {
	root TNode
	// TODO Check that root has no parents
}

type TNode struct {
	name     string
	children []*TNode
}

func (n *TNode) showIndent(depth int) {
	fmt.Print(strings.Repeat("	", depth))
	fmt.Println(n.name)
	for _, v := range n.children {
		v.showIndent(depth + 1)
	}
}

func (n *TNode) show() {
	n.showIndent(0)
}

func (n *TNode) AddChild(c *TNode) {
	n.children = append(n.children, c)
}

func (n *TNode) AddParent(p TNode) {

}

func (t *TTree) show() {
	t.root.show()
}

type TComparable struct {
	pops []string
}

func Compare(tc1, tc2 TComparable) {

}
