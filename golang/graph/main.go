
package main

import (
	"graph"
	"presenter"
)

func main() {
	g := graph.MakeAdjacencyMatrixGraph()
	vId1 := g.AddVertice()
	vId2 := g.AddVertice()
	vId3 := g.AddVertice()
	vId4 := g.AddVertice()
	g.AddEdge(vId1, vId2)
	g.AddEdge(vId3, vId4)
	g.AddEdge(vId4, vId3)
	g.AddEdge(vId4, vId4)
	g.AddEdge(vId4, vId2)

	g.BindDataVertice(graph.MakeLabelData("вершина0"), vId1)

	presenter.DotGraph(g)
}