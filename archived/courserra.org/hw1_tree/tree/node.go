package tree

import "fmt"

const (
	File      = 1
	Directory = 2
)

// Node represents a directory node.
type Node struct {
	Name  string
	Type  int
	size  Size
	Nodes []*Node
}

type Size int64

func (s Size) String() string {
	if s == 0 {
		return "empty"
	}
	return fmt.Sprintf("%db", s)
}

func (n Node) IsDir() bool {
	return n.Type == Directory
}

func (n Node) IsFile() bool {
	return n.Type == File
}

func (n Node) GetSize() Size {
	if n.IsFile() {
		return n.size
	}
	return Size(0)
}

func (n *Node) Add(child *Node) {
	n.Nodes = append(n.Nodes, child)
}
