package trie

// Node represents a node in the Trie
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *Node
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &Node{children: make(map[rune]*Node), isEnd: false}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
			isEnd:    false,
		},
	}
}
