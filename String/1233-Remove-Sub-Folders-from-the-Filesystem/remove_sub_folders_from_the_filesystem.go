package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

// removeSubFolder - doesn't use sorting, but rather use trie
func removeSubfolders(dirs []string) []string {

	if len(dirs) <= 1 {
		return dirs
	}

	var (
		root = NewPath("")
		head *Path
	)

	// populating directories
	for _, dir := range dirs {

		head = root

		if dir == "/" {
			root.Final = true
		}

		subDirectories := strings.Split(dir, "/")

		for idx, pathName := range subDirectories {

			// can be root, can be catch.
			if pathName == "" {
				continue
			}

			// just in case for automated root
			if head.Final {
				break
			}

			path, err := head.GetByName(pathName)
			if err != nil {

				subDirectory := NewPath(pathName)
				subDirectory.Final = idx == (len(subDirectories) - 1)

				head.AddSubDir(subDirectory)
				head = subDirectory
				continue
			}

			if path.Final {
				break
			}

			path.Final = idx == (len(subDirectories) - 1)

			// fmt.Println(dir, idx, pathName)

			head = path
			continue
		}
	}

	// i am lazy to recursivly traverse the tree
	// return strings.Split(root.String(), "\n")

	return root.Tree()
}

func (parent Path) Tree() []string {
	if parent.Final {
		return []string{parent.Name}
	}

	var results = make([]string, 0, len(parent.subDirectories))
	for _, s := range parent.subDirectories {
		for _, subdir := range s.Tree() {
			results = append(results, parent.Name+"/"+subdir)
		}
	}

	return results

}

type Path struct {
	sync.Mutex

	Name           string
	Final          bool
	subDirectories map[string]*Path
}

func (parent Path) String() string {

	if len(parent.subDirectories) == 0 {
		return parent.Name
	}

	var output bytes.Buffer

	for _, p := range parent.subDirectories {

		for _, subdir := range strings.Split(p.String(), "\n") {
			output.WriteString(parent.Name)
			output.WriteByte('/')
			output.WriteString(subdir)
			output.WriteByte('\n')
		}
	}

	return strings.TrimRight(output.String(), "\n")
}

func (parent Path) HasSubDir(path *Path) bool {
	_, exists := parent.subDirectories[path.Name]
	return exists
}

func (parent *Path) AddSubDir(path *Path) error {
	parent.Lock()
	defer parent.Unlock()
	if parent.HasSubDir(path) {
		return fmt.Errorf("Directory %s already exists in %s", path.Name, parent.Name)
	}

	parent.subDirectories[path.Name] = path

	return nil
}

func (parent Path) GetByName(name string) (*Path, error) {
	if path, exists := parent.subDirectories[name]; exists {
		return path, nil
	}
	return nil, fmt.Errorf("Path %s not found", name)
}

func NewPath(name string) *Path {
	return &Path{
		Name:           name,
		subDirectories: make(map[string]*Path),
	}
}
