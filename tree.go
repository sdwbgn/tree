package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type DirTreeNode struct {
	elements []DirTreeNode
	name     string
}

func generateDirTree(dir string, name string, dirOnly bool) DirTreeNode {
	val := DirTreeNode{
		elements: nil,
		name:     name,
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return val
	}

	for _, fInfo := range files {
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

func printBranches(branches *[]bool, lastChars string) {
	for key, br := range *branches {
		if key+1 == len(*branches) {
			fmt.Print(lastChars)
		} else if br {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
}

func printDirTree(node DirTreeNode, branches *[]bool) {
	fmt.Println(node.name)
	if branches == nil {
		branches = new([]bool)
	}
	*branches = append(*branches, true)
	for key, file := range node.elements {
		if key+1 == len(node.elements) {
			printBranches(branches, "└── ")
			(*branches)[len(*branches)-1] = false
		} else {
			printBranches(branches, "├── ")
		}
		printDirTree(file, branches)
	}
	*branches = (*branches)[:len(*branches)-1]
}

func main() {
	dirOnly := false
	dir := ""
	if len(os.Args) == 1 {
		dir = "."
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-d" {
			dirOnly = true
			dir = "."
		} else {
			dir = os.Args[1]
		}
	} else if len(os.Args) == 3 && (os.Args[1] == "-d" || os.Args[2] == "-d") {
		dirOnly = true
		if os.Args[1] == "-d" {
			dir = os.Args[2]
		} else {
			dir = os.Args[1]
		}
	} else {
		fmt.Print("Usage: tree [directory] [-d]\n" +
			"  -d\tShow only directories\n" +
			"  If no directory provided, current directory will be traversed")
		return
	}
	printDirTree(generateDirTree(dir, dir, dirOnly), nil)
}
