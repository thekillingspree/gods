package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := InitTrie()
	assert.NotNil(t, trie.root, "should not be nil")
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	trie := InitTrie()
	toInsert := "abc"
	trie.Insert(toInsert)
	curNode := trie.root
	for i := range toInsert {
		assert.NotNil(curNode.children[toInsert[i]-'a'], "it should not be nil")
		curNode = curNode.children[toInsert[i]-'a']
	}
	assert.True(curNode.isEnd, "it should be the end of the word")
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	trie := InitTrie()
	toInsert := "abc"
	trie.Insert(toInsert)
	assert.True(trie.Search(toInsert), "it should be present in the trie")
	assert.False(trie.Search("hello"), "it should not be present in the trie")
}
