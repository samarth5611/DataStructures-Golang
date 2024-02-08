package main
import "fmt"

// Structure
type Node struct{
	child [26]*Node
	isEnd bool
}

type Trie struct{
	root *Node
}

func BuildTrie() Trie{
	return Trie{root: &Node{isEnd: false}}
}
// Insert
func (T *Trie)Insert(s string){
	curNode := T.root
	for i :=0 ; i<len(s); i++{
		charIndex := s[i] - 'a'
		if curNode.child[charIndex] == nil{
			curNode.child[charIndex] = &Node{}
		}
		curNode = curNode.child[charIndex]
	}
	curNode.isEnd = true
}
// Search
func (T *Trie)Search(s string)bool{
	curNode := T.root
	for i:=0; i<len(s); i++{
		charIndex := s[i] - 'a'
		if curNode.child[charIndex] == nil{
			return false
		}
		curNode = curNode.child[charIndex]
	}
	if curNode.isEnd == true{
		return true
	}
	return false
}

func main(){
	trie := BuildTrie()
	fmt.Println(trie.root)
	trie.Insert("sammy")
	trie.Insert("shark")
	trie.Insert("satisfaction")

	fmt.Println(trie.Search("sammy"))
	fmt.Println(trie.Search("xyc"))
	fmt.Println(trie.Search(""))
	
}