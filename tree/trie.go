package tree

import (
	"strings"
)

type trieNode struct {
	children    [26]*trieNode
	isEndOfWord bool
}

type Trie struct {
	Root *trieNode
}

func (t *Trie) Insert(key string) {
	if t.Root == nil {
		t.Root = &trieNode{}
	}
	thisNode := t.Root
	chars := wordToCharIndex(key)
	for i := 0; i < len(chars); i++ {
		if thisNode.children[chars[i]] == nil {
			newNode := &trieNode{}
			thisNode.children[chars[i]] = newNode
			thisNode = newNode
		} else {
			thisNode = thisNode.children[chars[i]]
		}
	}
	thisNode.isEndOfWord = true
}

func (t *Trie) Search(key string) bool {
	if t.Root == nil {
		return false
	}
	chars := wordToCharIndex(key)
	cntr := 0
	thisNode := t.Root.children[chars[cntr]]
	for thisNode != nil {
		if cntr == len(chars)-1 {
			if !thisNode.isEndOfWord {
				break
			}
			return true
		}
		cntr++
		thisNode = thisNode.children[chars[cntr]]
	}
	return false
}

func (t *Trie) Delete(key string) {
	if t.Root == nil {
		return
	}
	thisNode := t.Root
	chars := wordToCharIndex(key)
	for _, char := range chars {
		if thisNode.children[char] != nil {
			thisNode = thisNode.children[char]
		}
	}
	thisNode.isEndOfWord = false
}

func (t *Trie) Keys() []string {
	if t.Root == nil {
		return nil
	}
	words := &[]string{}
	collectWords(t.Root, []string{}, words)
	return *words
}

func collectWords(node *trieNode, word []string, words *[]string) {
	for i, v := range node.children {
		if v != nil {
			local := []string{}
			local = append(local, word...)
			local = append(local, indexToChar(i))
			if v.isEndOfWord {
				s := strings.Join(local, "")
				*words = append(*words, s)
			}
			collectWords(v, local, words)
		}
	}
}

func wordToCharIndex(word string) (indices []int) {
	word = strings.ToLower(word)
	for _, ch := range word {
		indices = append(indices, (int(ch) - 97))
	}
	return
}

func indexToChar(i int) string {
	return string(rune(i + 97))
}
