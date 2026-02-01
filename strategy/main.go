package strategy

import "fmt"

type Strategy interface {
	Execute() error
}

type Node struct {
	strategy Strategy
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) SetStrategy(strategy Strategy) {
	n.strategy = strategy
}

func (n *Node) Execute() error {
	return n.strategy.Execute()
}

type FilterNode struct{}

func (f *FilterNode) Execute() error {
	fmt.Println("Filtering data")
	return nil
}

type SortNode struct{}

func (s *SortNode) Execute() error {
	fmt.Println("Sorting data")
	return nil
}

func Main() {
	filterNode := &FilterNode{}
	sortNode := &SortNode{}

	node := NewNode()
	node.SetStrategy(filterNode)
	node.Execute()

	node.SetStrategy(sortNode)
	node.Execute()
}
