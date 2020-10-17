package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	hashTable := InitHashTable()
	cases := map[string]interface{}{
		"abc":         []int{1, 2, 3},
		"Hello World": []int{10, 20, 30},
		"1":           "Hello 👋",
	}
	for key, val := range cases {
		hashTable.Insert(key, val)
	}
	hashTable.Insert("1", "Hello") //Replacing
	assert.Equal(len(cases), hashTable.Len(), "all cases must be added.")
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	hashTable := InitHashTable()

	//Empty hashtable
	present := hashTable.Search("abc")
	assert.False(present, "should be false")

	hashTable.Insert("abc", 123)
	hashTable.Insert("abcd", 123)

	present = hashTable.Search("abc")
	assert.True(present, "should be true")
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	hashTable := InitHashTable()
	cases := map[string]interface{}{
		"abc":         []int{1, 2, 3},
		"Hello World": []int{10, 20, 30},
		"1":           "Hello 👋",
	}

	//Empty hashtable
	gotVal, err := hashTable.Get("abc")
	assert.NotNil(err, "should not be nil")
	assert.Nil(gotVal, "should be nil")

	for key, val := range cases {
		hashTable.Insert(key, val)
	}

	for key, wantVal := range cases {
		gotVal, err = hashTable.Get(key)
		assert.Nil(err, "should be nil")
		assert.Equal(wantVal, gotVal, "should be equal")
	}
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	hashTable := InitHashTable()

	//Empty hashtable
	err := hashTable.Delete("abc")
	assert.NotNil(err, "should not be nil")

	hashTable.Insert("abc", 123)

	//Invalid key
	err = hashTable.Delete("abcd")
	assert.NotNil(err, "should not be nil")

	hashTable.Insert("abcd", 123)
	hashTable.Insert("abce", 123)
	hashTable.Insert("abcef", 123)

	err = hashTable.Delete("abc")
	assert.Nil(err, "should be nil")
	assert.Equal(3, hashTable.Len(), "should be 2")
}
