package model

type Node struct {
	Key  string
	Val  string
	Prev *Node
	Next *Node
}
