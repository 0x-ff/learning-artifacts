
package graph

import (
	"graph"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	assert.True(t, g.IsEmpty())
}

func TestIsNotEmpty(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	g.AddVertice()
	assert.False(t, g.IsEmpty())
}

func TestNotHasVertice(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	assert.False(t, g.HasVertice(0))
}

func TestHasVertice(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	g.AddVertice()
	assert.True(t, g.HasVertice(0))
}

func TestNotHasEdge(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	g.AddVertice()
	g.AddVertice()
	g.AddEdge(1, 0)
	assert.False(t, g.HasEdge(0, 1))
}

func TestHasEdge(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	g.AddVertice()
	g.AddVertice()
	g.AddEdge(0, 1)
	assert.True(t, g.HasEdge(0, 1))
}

func TestVerticesCount(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	assert.Equal(t, 0, g.VerticesCount())

	g.AddVertice()
	assert.Equal(t, 1, g.VerticesCount())
}

func TestEdgesCount(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	assert.Equal(t, 0, g.EdgesCount())

	g.AddVertice()
	g.AddVertice()

	assert.Equal(t, 0, g.EdgesCount())

	g.AddEdge(0, 1)
	assert.Equal(t, 1, g.EdgesCount())

	g.AddEdge(1, 0)
	assert.Equal(t, 2, g.EdgesCount())
}

func TestAddVertice(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	assert.Equal(t, 0, g.AddVertice())
	assert.Equal(t, 1, g.VerticesCount())

	assert.Equal(t, 1, g.AddVertice())
	assert.Equal(t, 2, g.VerticesCount())

	assert.Equal(t, 2, g.AddVertice())
	assert.Equal(t, 3, g.VerticesCount())
}

func TestAddEdge(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	assert.Equal(t, 0, g.EdgesCount())

	err, ok := g.AddEdge(1, 0)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 1, g.EdgesCount())

	err, ok = g.AddEdge(0, 1)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 2, g.EdgesCount())
}

func TestAddEdgeComplex(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	g.AddVertice()

	err, ok := g.AddEdge(0, 1)
	assert.True(t, ok)
	assert.NoError(t, err)

	err, ok = g.AddEdge(0, 2)
	assert.True(t, ok)
	assert.NoError(t, err)

	err, ok = g.AddEdge(1, 0)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 3, g.EdgesCount())
}

func TestAddEdgeToVerticeNotExists(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	err, ok := g.AddEdge(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "Не существует вершины 0, поэтому не могу добавить ребро 1 -> 0", err.Error())
}

func TestAddEdgeFromVerticeNotExists(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	err, ok := g.AddEdge(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "Не существует вершины 1, поэтому не могу добавить ребро 1 -> 0", err.Error())
}

func TestAddEdgeAlreadyExists(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()

	g.AddEdge(0, 1)
	err, ok := g.AddEdge(0, 1)

	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 1, g.EdgesCount())
}

func TestRemoveEdgesEnding(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	g.AddVertice()
	g.AddVertice()

	g.AddEdge(0, 1) // will be removed
	g.AddEdge(1, 1) // will be removed
	g.AddEdge(2, 1) // will be removed
	g.AddEdge(3, 1) // will be removed
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 6, g.EdgesCount())

	ok := g.RemoveEdgesEnding(1)
	assert.True(t, ok)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())

	assert.True(t, g.HasEdge(1, 3))
	assert.True(t, g.HasEdge(2, 3))
	assert.False(t, g.HasEdge(0, 1))
	assert.False(t, g.HasEdge(1, 1))
	assert.False(t, g.HasEdge(2, 1))
	assert.False(t, g.HasEdge(3, 1))
}

func TestRemoveVertice(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	g.AddVertice()
	g.AddVertice()

	g.AddEdge(0, 1) // will be removed
	g.AddEdge(1, 1) // will be removed
	g.AddEdge(2, 1) // will be removed
	g.AddEdge(3, 1) // will be removed
	g.AddEdge(1, 3) // will be removed
	g.AddEdge(2, 3)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 6, g.EdgesCount())

	ok := g.RemoveVertice(1)
	assert.True(t, ok)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 1, g.EdgesCount())

	assert.True(t, g.HasEdge(2, 3))
	assert.False(t, g.HasEdge(0, 1))
	assert.False(t, g.HasEdge(1, 1))
	assert.False(t, g.HasEdge(2, 1))
	assert.False(t, g.HasEdge(3, 1))
	assert.False(t, g.HasEdge(1, 3))
}

func TestRemoveEdgeNeverExist(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	ok := g.RemoveEdge(0, 1)
	assert.False(t, ok)
}

func TestRemoveRemovedEdge(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	g.AddVertice()

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 0)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 3, g.EdgesCount())

	ok := g.RemoveEdge(0, 1)
	assert.True(t, ok)
	assert.False(t, g.HasEdge(0, 1))

	ok = g.RemoveEdge(0, 1)
	assert.False(t, ok)
	assert.False(t, g.HasEdge(0, 1))

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())
	assert.True(t, g.HasEdge(0, 2))
	assert.True(t, g.HasEdge(1, 0))
}

func TestRemoveEdge(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()

	g.AddVertice()
	g.AddVertice()
	g.AddVertice()

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 0)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 3, g.EdgesCount())

	ok := g.RemoveEdge(0, 1)
	assert.True(t, ok)
	assert.False(t, g.HasEdge(0, 1))

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())
	assert.True(t, g.HasEdge(0, 2))
	assert.True(t, g.HasEdge(1, 0))
}

func TestBindDataVerticeNotExist(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	l := graph.MakeLabelData("bind")

	ok := g.BindDataVertice(l, 1)
	assert.False(t, ok)

	ok, data := g.GetDataVertice(1)
	assert.False(t, ok)
	assert.Nil(t, data)
}

func TestBindDataVertice(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	vId := g.AddVertice()
	l := graph.MakeLabelData("bind")

	ok := g.BindDataVertice(l, vId)
	assert.True(t, ok)

	ok, data := g.GetDataVertice(vId)
	assert.True(t, ok)
	assert.NotNil(t, data)
	assert.Equal(t, "bind", data.Get())
}

func TestBindDataEdgeNotExist(t *testing.T) {
	g := graph.MakeAdjacencyListGraph()
	l := graph.MakeLabelData("bind")

	ok := g.BindDataEdge(l, 1, 1)
	assert.False(t, ok)

	ok, data := g.GetDataEdge(1, 1)
	assert.False(t, ok)
	assert.Nil(t, data)
}
