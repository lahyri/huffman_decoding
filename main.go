package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Node - basic structure
type Node struct {
	Char       rune
	Frequency  int
	LeftChild  *Node
	RightChild *Node
}

func (t *Node) buildTree(n []Node) {
	if len(n) > 0 {
		if t.LeftChild == nil {

			t.LeftChild = &n[0]
			n = n[1:]

			t.buildTree(n)

		} else if t.RightChild == nil {

			if len(n) > 1 {

				remaining := 0

				for _, r := range n {
					remaining += r.Frequency
				}

				t.RightChild = &Node{
					Char:       0,
					Frequency:  remaining,
					LeftChild:  nil,
					RightChild: nil,
				}
				t.RightChild.buildTree(n)

			} else {

				t.RightChild = &n[0]
				n = n[1:]
				t.RightChild.buildTree(n)
			}

		} else {
			t.RightChild.buildTree(n)
		}
	}
}

func (t *Node) PrintTree() {
	fmt.Println(t)

	fmt.Println("left child")
	fmt.Println(t.LeftChild)
	if t.RightChild != nil {
		fmt.Println("right child")
		t.RightChild.PrintTree()
	}
}

//based on a refference string, builds the corresponding Huffman coding tree
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert base codeword to the algorithm")
	fmt.Print("-> ")
	base, _ := reader.ReadString('\n')
	base = strings.Replace(base, "\n", "", -1)

	runedstring := []rune(base)
	unordered := []Node{Node{
		Char:      runedstring[0],
		Frequency: 1,
	}}
	runedstring = runedstring[1:]

	//checks if a rune already has been detected in the unoredered array
	for _, currentRune := range runedstring {

		for j, node := range unordered {

			if currentRune == node.Char {
				node.Frequency++
				unordered[j] = node

				break

			} else if j+1 == len(unordered) {

				newChar := Node{
					Char:       currentRune,
					Frequency:  1,
					LeftChild:  nil,
					RightChild: nil,
				}
				unordered = append(unordered, newChar)
			}
		}
	}

	orderedRunes := []Node{}
	for len(unordered) > 0 {
		majRune := Node{Frequency: 0}
		pos := 0
		for i, currRune := range unordered {
			if currRune.Frequency > majRune.Frequency {
				majRune = currRune
				pos = i
			}
		}
		orderedRunes = append(orderedRunes, unordered[pos])
		unordered = append(unordered[:pos], unordered[pos+1:]...)
	}
	root := Node{
		Char:       0,
		Frequency:  len(base),
		LeftChild:  nil,
		RightChild: nil,
	}
	root.buildTree(orderedRunes)
	root.PrintTree()
	fmt.Println("Insert coded word")
	fmt.Print("-> ")
	directions, _ := reader.ReadString('\n')
	directions = strings.Replace(directions, "\n", "", -1)
	fmt.Println(BuildWord(directions, &root))
}

func BuildWord(codedWord string, treeRoot *Node) string {
	if _, err := strconv.Atoi(codedWord); err != nil {
		return "ERROR: INCORRECT CODED WORD"
	}
	var directions []int
	for _, digit := range codedWord {
		directions = append(directions, int(digit)-int('0'))
	}
	var word, letter string
	for len(directions) > 0 && directions != nil {
		letter, directions = treeRoot.ReachLetter(directions)
		word += letter
	}
	return word
}

func (t *Node) ReachLetter(directions []int) (string, []int) {
	switch directions[0] {
	case 0:
		node := *t.LeftChild

		return string(node.Char), directions[1:]
	case 1:
		node := *t.RightChild
		if node.Char == 0 {
			directions = directions[1:]
			return node.ReachLetter(directions)
		}
		return string(node.Char), directions[1:]
	}
	return "ERROR", nil
}
