package main

import "io/ioutil"

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
		} else {
			val.elements = append(val.elements, DirTreeNode{
				elements: nil,
				name:     fInfo.Name(),
			})
		}
	}
	return val
}

func printBranches(branches []bool, lastChars string) {
	for key, i := range branches {
		if key+1 == len(branches) {
			print(lastChars)
		} else if i {
			print("|   ")
		} else {
			print("    ")
		}
	}
}

func printDirTree(node DirTreeNode, branches []bool) {

}

func main() {

}
