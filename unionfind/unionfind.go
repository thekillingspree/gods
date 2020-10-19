package unionfind

//UnionFind represents a Union Find Data structure
type UnionFind struct {
	ids        []int
	sizes      []int
	components int
}

// InitUnionFind will initialize a new Union Find data structure.
func InitUnionFind(n int) *UnionFind {
	uf := UnionFind{}
	uf.ids = make([]int, n)
	uf.sizes = make([]int, n)
	uf.components = n
	for i := 0; i < n; i++ {
		uf.ids[i] = i
		uf.sizes[i] = 1
	}

	return &uf
}

// Find the root of a component
func (uf *UnionFind) root(i int) int {
	for i != uf.ids[i] {
		uf.ids[i] = uf.ids[uf.ids[i]]
		i = uf.ids[i]
	}

	return i
}

// Find will return true if the two components are connected.
func (uf UnionFind) Find(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

// Union connects two components
func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.root(p)
	qRoot := uf.root(q)
	if pRoot == qRoot {
		return
	}

	if uf.sizes[pRoot] > uf.sizes[qRoot] {
		uf.ids[qRoot] = pRoot
		uf.sizes[pRoot] += uf.sizes[qRoot]
	} else {
		uf.ids[pRoot] = qRoot
		uf.sizes[qRoot] += uf.sizes[pRoot]
	}
	uf.components--
}

// Components returns the total number of connected components
func (uf UnionFind) Components() int {
	return uf.components
}
