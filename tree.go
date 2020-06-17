package main

import (
	"flag"
	"io/ioutil"
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
			print(lastChars)
		} else if br {
			print("│   ")
		} else {
			print("    ")
		}
	}
}

func printDirTree(node DirTreeNode, branches *[]bool) {
	println(node.name)
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
	dirOnly := flag.Bool("d", false, "Show only directories")
	flag.Parse()
	dir := ""
	if flag.NArg() == 0 {
		dir = "."
	} else if flag.NArg() == 1 {
		dir = flag.Arg(0)
	} else {
		flag.PrintDefaults()
		return
	}
	printDirTree(generateDirTree(dir, dir, *dirOnly), nil)
}
