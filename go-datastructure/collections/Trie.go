package collections

type trieNode struct {
	EndOfWord bool
	Children  map[byte]*trieNode
}

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{root: &trieNode{
		EndOfWord: false,
		Children:  make(map[byte]*trieNode),
	}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root

	for i := 0; i < len(word); i++ {
		dict := node.Children

		if v, ok := dict[word[i]]; ok {
			node = v
		} else {
			newNode := &trieNode{
				EndOfWord: false,
				Children:  make(map[byte]*trieNode),
			}
			dict[word[i]] = newNode
			node = newNode
		}
	}

	node.EndOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.search(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.search(prefix, false)
}

func (this *Trie) search (word string, findComplete bool) bool{
	node := this.root

	for i := 0; i < len(word); i++ {
		dict := node.Children

		if v, ok := dict[word[i]]; ok {
			node = v
		} else {
			return false
		}
	}

	return !findComplete || node.EndOfWord
}
