package trie

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

//Trie Represents a trie
type Trie struct {
	root *Node
}

// InitTrie creates a new trie
func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}
	return result
}

// Insert inserts a word into the trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, val := range word {
		charIndex := val - 'a'
		if current.children[charIndex] == nil {
			node := &Node{}
			current.children[charIndex] = node
		}
		current = current.children[charIndex]
	}
	current.isEnd = true
}

// Search searches a word in the trie
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, val := range word {
		charIndex := val - 'a'
		if current.children[charIndex] == nil {
			return false
		}
		current = current.children[charIndex]
	}

	return current.isEnd
}
