package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	assert := assert.New(t)
	uf := InitUnionFind(10)

	uf.Union(5, 6)
	assert.Equal(9, uf.components, "should be 9")
	uf.Union(3, 4)
	uf.Union(2, 9)
	uf.Union(4, 9)
	assert.Equal(9, uf.root(3), uf.root(2), "should be equal")
	assert.Equal(6, uf.components, "should be equal")

	// Already connected, uf.components must not change
	uf.Union(3, 2)
	assert.Equal(6, uf.components, "should be equal")

	uf.Union(9, 5)
	assert.Equal(9, uf.root(6), "should be equal")
}

func TestFind(t *testing.T) {
	assert := assert.New(t)
	uf := InitUnionFind(10)
	uf.Union(5, 6)
	assert.True(uf.Find(5, 6))
	uf.Union(3, 4)
	uf.Union(2, 9)
	uf.Union(4, 9)
	assert.True(uf.Find(2, 3))
	assert.False(uf.Find(2, 5))
}

func TestComponents(t *testing.T) {
	assert := assert.New(t)
	uf := InitUnionFind(10)
	uf.Union(5, 6)
	assert.Equal(9, uf.Components())
}
