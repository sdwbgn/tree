package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Tree node
// Contains children nodes and node name
type DirTreeNode struct {
	elements []DirTreeNode
	name     string
}

// Builds DirTreeNode tree for given dir and name
// If dirOnly is true, then it will ignore files
func generateDirTree(dir string, name string, dirOnly bool) DirTreeNode {
	// init tree node
	val := DirTreeNode{
		elements: nil,
		name:     name,
	}

	// try to read given dir
	files, err := ioutil.ReadDir(dir)

	// if error occurred, then ignore the directory
	if err != nil {
		return val
	}

	for _, fInfo := range files {
		// if directory, then recursively generate tree,
		// else, if not set to ignore files, add list node to the tree
		if fInfo.IsDir() {
			val.elements = append(val.elements, generateDirTree(dir+"/"+fInfo.Name(), fInfo.Name(), dirOnly))
		} else if !dirOnly {
			val.elements = append(val.elements, DirTreeNode{
				elements: nil,
				name:     fInfo.Name(),
			})
		}
	}
	return val
}

// Prints branches from given branches array and uses
// lastBranch for last branch
func printBranches(branches *[]bool, lastBranch string) {
	for key, br := range *branches {
		if key+1 == len(*branches) {
			fmt.Print(lastBranch)
		} else if br {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
}

// Recursively prints tree from given DirTreeNode
func printDirTree(node DirTreeNode, branches *[]bool) {
	// print node name
	fmt.Println(node.name)

	// init branches array
	if branches == nil {
		branches = new([]bool)
	}

	// increase branch level
	*branches = append(*branches, true)

	for key, file := range node.elements {
		// if last element in node, then change the branch style
		if key+1 == len(node.elements) {
			printBranches(branches, "└── ")
			(*branches)[len(*branches)-1] = false
		} else {
			printBranches(branches, "├── ")
		}

		// continue to recursively print directory tree
		printDirTree(file, branches)
	}

	// decrease branch level
	*branches = (*branches)[:len(*branches)-1]
}

func main() {
	// init parameters
	dirOnly := false
	dir := ""

	// Parse arguments
	if len(os.Args) == 1 { // no arguments
		dir = "."
	} else if len(os.Args) == 2 { // one argument
		if os.Args[1] == "-d" {
			dirOnly = true
			dir = "."
		} else {
			dir = os.Args[1]
		}
	} else if len(os.Args) == 3 && (os.Args[1] == "-d" || os.Args[2] == "-d") { // two arguments
		dirOnly = true
		if os.Args[1] == "-d" {
			dir = os.Args[2]
		} else {
			dir = os.Args[1]
		}
	} else { // invalid arguments
		print("Usage: tree [directory] [-d]\n" +
			"  -d\tShow only directories\n" +
			"  If no directory provided, current directory will be traversed")
		return
	}

	// check if given directory exists and is it valid
	info, err := os.Stat(dir)
	if err != nil {
		print("Directory doesn't exist or not enough permissions")
		return
	}
	if !info.IsDir() {
		print("Not a directory")
		return
	}

	// print result
	printDirTree(generateDirTree(dir, dir, dirOnly), nil)
}
