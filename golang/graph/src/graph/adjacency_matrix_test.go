
package graph

import (
	"graph"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	assert.True(t, g.IsEmpty())
}

func TestIsNotEmpty(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	g.AddVertice()
	assert.False(t, g.IsEmpty())
}

func TestNotHasVertice(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	assert.False(t, g.HasVertice(0))
}

func TestHasVertice(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	vId := g.AddVertice()
	assert.True(t, g.HasVertice(vId))
}

func TestNotHasEdge(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	err, added := g.AddEdge(v1Id, v2Id)
	assert.NoError(t, err)
	assert.True(t, added)
	assert.False(t, g.HasEdge(v2Id, v1Id))
}

func TestHasEdge(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	err, added := g.AddEdge(v1Id, v2Id)
	assert.NoError(t, err)
	assert.True(t, added)
	assert.True(t, g.HasEdge(v1Id, v2Id))
}

func TestVerticesCount(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	assert.Equal(t, 0, g.VerticesCount())

	g.AddVertice()
	assert.Equal(t, 1, g.VerticesCount())
}

func TestEdgesCount(t *testing.T) {
	g := graph.MakeIncidentMatrixGraph()
	assert.Equal(t, 0, g.EdgesCount())

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()

	assert.Equal(t, 0, g.EdgesCount())

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.NoError(t, err)
	assert.True(t, ok)

	assert.Equal(t, 1, g.EdgesCount())

	err, ok = g.AddEdge(v2Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 2, g.EdgesCount())
}

func TestInEdgesCount(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()

	assert.Equal(t, 0, g.InEdgesCount(v1Id))
	assert.Equal(t, 0, g.InEdgesCount(v2Id))
	assert.Equal(t, 0, g.InEdgesCount(v3Id))

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 0, g.InEdgesCount(v1Id))
	assert.Equal(t, 1, g.InEdgesCount(v2Id))
	assert.Equal(t, 0, g.InEdgesCount(v3Id))

	err, ok = g.AddEdge(v2Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	err, ok = g.AddEdge(v1Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 2, g.InEdgesCount(v1Id))
	assert.Equal(t, 1, g.InEdgesCount(v2Id))
	assert.Equal(t, 0, g.InEdgesCount(v3Id))

	err, ok = g.AddEdge(v3Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 3, g.InEdgesCount(v1Id))
	assert.Equal(t, 1, g.InEdgesCount(v2Id))
	assert.Equal(t, 0, g.InEdgesCount(v3Id))

	ok = g.RemoveEdge(v1Id, v1Id)
	assert.True(t, ok)
	assert.Equal(t, 2, g.InEdgesCount(v1Id))
	assert.Equal(t, 1, g.InEdgesCount(v2Id))
	assert.Equal(t, 0, g.InEdgesCount(v3Id))
}

func TestOutEdgesCount(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()

	assert.Equal(t, 0, g.OutEdgesCount(v1Id))
	assert.Equal(t, 0, g.OutEdgesCount(v2Id))
	assert.Equal(t, 0, g.OutEdgesCount(v3Id))

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 1, g.OutEdgesCount(v1Id))
	assert.Equal(t, 0, g.OutEdgesCount(v2Id))
	assert.Equal(t, 0, g.OutEdgesCount(v3Id))

	err, ok = g.AddEdge(v2Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	err, ok = g.AddEdge(v1Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 2, g.OutEdgesCount(v1Id))
	assert.Equal(t, 1, g.OutEdgesCount(v2Id))
	assert.Equal(t, 0, g.OutEdgesCount(v3Id))

	err, ok = g.AddEdge(v3Id, v1Id)
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, 2, g.OutEdgesCount(v1Id))
	assert.Equal(t, 1, g.OutEdgesCount(v2Id))
	assert.Equal(t, 1, g.OutEdgesCount(v3Id))

	ok = g.RemoveEdge(v1Id, v1Id)
	assert.True(t, ok)
	assert.Equal(t, 1, g.OutEdgesCount(v1Id))
	assert.Equal(t, 1, g.OutEdgesCount(v2Id))
	assert.Equal(t, 1, g.OutEdgesCount(v3Id))
}

func TestAddVertice(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	assert.Equal(t, 0, g.AddVertice())
	assert.Equal(t, 1, g.VerticesCount())

	assert.Equal(t, 1, g.AddVertice())
	assert.Equal(t, 2, g.VerticesCount())

	assert.Equal(t, 2, g.AddVertice())
	assert.Equal(t, 3, g.VerticesCount())
}

func TestAddEdge(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	assert.Equal(t, 0, g.EdgesCount())

	err, ok := g.AddEdge(v2Id, v1Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 1, g.EdgesCount())

	err, ok = g.AddEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 2, g.EdgesCount())
}

func TestAddEdgeComplex(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	err, ok = g.AddEdge(v1Id, v3Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	err, ok = g.AddEdge(v2Id, v1Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 3, g.EdgesCount())
}

func TestAddEdgeToVerticeNotExists(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	err, ok := g.AddEdge(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "Не существует вершины 0, поэтому не могу добавить ребро 1 -> 0", err.Error())
}

func TestAddEdgeFromVerticeNotExists(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	g.AddVertice()
	err, ok := g.AddEdge(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "Не существует вершины 1, поэтому не могу добавить ребро 1 -> 0", err.Error())
}

func TestAddEdgeAlreadyExists(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	err, ok = g.AddEdge(v1Id, v2Id)
	assert.False(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, 1, g.EdgesCount())
}

func TestRemoveEdgesEnding(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()
	v4Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v3Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v4Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v2Id, v4Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v3Id, v4Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 5, g.EdgesCount())

	ok = g.RemoveEdgesEnding(1)
	assert.True(t, ok)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())

	assert.True(t, g.HasEdge(v2Id, v4Id))
	assert.True(t, g.HasEdge(v3Id, v4Id))
	assert.False(t, g.HasEdge(v1Id, v2Id))
	assert.False(t, g.HasEdge(v2Id, v2Id))
	assert.False(t, g.HasEdge(v3Id, v2Id))
	assert.False(t, g.HasEdge(v4Id, v2Id))
}

func TestRemoveVertice(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()
	v4Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v3Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v4Id, v2Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v2Id, v4Id) // will be removed
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v3Id, v4Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 4, g.VerticesCount())
	assert.Equal(t, 5, g.EdgesCount())

	ok = g.RemoveVertice(1)
	assert.True(t, ok)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 1, g.EdgesCount())

	assert.True(t, g.HasEdge(v3Id, v4Id))
	assert.False(t, g.HasEdge(v1Id, v2Id))
	assert.False(t, g.HasEdge(v2Id, v2Id))
	assert.False(t, g.HasEdge(v3Id, v2Id))
	assert.False(t, g.HasEdge(v4Id, v2Id))
	assert.False(t, g.HasEdge(v2Id, v4Id))
}

func TestRemoveEdgeNeverExist(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	ok := g.RemoveEdge(0, 1)
	assert.False(t, ok)
}

func TestRemoveRemovedEdge(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v1Id, v3Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v2Id, v1Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 3, g.EdgesCount())

	ok = g.RemoveEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.False(t, g.HasEdge(v1Id, v2Id))

	ok = g.RemoveEdge(v1Id, v2Id)
	assert.False(t, ok)
	assert.False(t, g.HasEdge(v1Id, v2Id))

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())
	assert.True(t, g.HasEdge(v1Id, v3Id))
	assert.True(t, g.HasEdge(v2Id, v1Id))
}

func TestRemoveEdge(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()

	v1Id := g.AddVertice()
	v2Id := g.AddVertice()
	v3Id := g.AddVertice()

	err, ok := g.AddEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v1Id, v3Id)
	assert.True(t, ok)
	assert.NoError(t, err)
	err, ok = g.AddEdge(v2Id, v1Id)
	assert.True(t, ok)
	assert.NoError(t, err)

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 3, g.EdgesCount())

	ok = g.RemoveEdge(v1Id, v2Id)
	assert.True(t, ok)
	assert.False(t, g.HasEdge(v1Id, v2Id))

	assert.Equal(t, 3, g.VerticesCount())
	assert.Equal(t, 2, g.EdgesCount())
	assert.True(t, g.HasEdge(v1Id, v3Id))
	assert.True(t, g.HasEdge(v2Id, v1Id))
}

func TestBindDataVerticeNotExist(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
	l := graph.MakeLabelData("bind")

	ok := g.BindDataVertice(l, 1)
	assert.False(t, ok)

	ok, data := g.GetDataVertice(1)
	assert.False(t, ok)
	assert.Nil(t, data)
}

func TestBindDataVertice(t *testing.T) {
	g := graph.MakeAdjacencyMatrixGraph()
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
	g := graph.MakeAdjacencyMatrixGraph()
	l := graph.MakeLabelData("bind")

	ok := g.BindDataEdge(l, 1, 1)
	assert.False(t, ok)

	ok, data := g.GetDataEdge(1, 1)
	assert.False(t, ok)
	assert.Nil(t, data)
}
