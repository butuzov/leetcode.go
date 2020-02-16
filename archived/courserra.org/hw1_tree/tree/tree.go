package tree

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

var (
	charDefault = '├'
	charEnd     = '└'
	patternDir  = "%[3]s%[1]c───%[2]s\n"
	patternFile = "%[3]s%[1]c───%[2]s (%[4]s)\n"
)

// Structure return nested structure (including files)
func Structure(fileSystemEntry string) (*Node, error) {

	fi, err := os.Stat(fileSystemEntry)
	if err != nil {
		return nil, err
	}

	var node = &Node{
		Name: path.Base(fileSystemEntry),
	}

	if !fi.IsDir() {
		node.Type = File
		node.size = Size(fi.Size())
		return node, nil
	}

	directoryContents, e := filepath.Glob(fileSystemEntry + "/*")
	if e != nil {
		return nil, e
	}

	node.Type = Directory
	for i := range directoryContents {
		item, e := Structure(directoryContents[i])
		if e != nil {
			continue
		}
		node.Add(item)
	}

	return node, nil
}

// filterOutFiles will return nodes (excuding files)
// so, only directories allowed.
func filterOutFiles(childrens []*Node) []*Node {

	var nodes []*Node
	for _, n := range childrens {
		if n.IsDir() {
			nodes = append(nodes, n)
		}
	}

	return nodes
}

// Tree is printer for nodes structure.
func Tree(out io.Writer, node *Node, printFiles bool, offset string) {

	var charStart = charDefault
	var nodes = node.Nodes

	if !printFiles {
		nodes = filterOutFiles(nodes)
	}

	for i, n := range nodes {

		if i == len(nodes)-1 {
			charStart = charEnd
		}

		if n.IsDir() {

			fmt.Fprintf(out, patternDir, charStart, n.Name, offset)

			var charSep = ""
			if i != len(nodes)-1 {
				charSep = "│"
			}

			Tree(out, n, printFiles, fmt.Sprintf("%s%s\t", offset, charSep))
			continue
		}

		fmt.Fprintf(out, patternFile, charStart, n.Name, offset, n.GetSize())
	}
}
