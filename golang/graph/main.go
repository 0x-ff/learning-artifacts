
package main

import (
	"graph"
	"presenter"
)

func main() {
	adjListGraph := graph.MakeAdjacencyListGraph()
	vId1 := adjListGraph.AddVertice()
	vId2 := adjListGraph.AddVertice()
	vId3 := adjListGraph.AddVertice()
	vId4 := adjListGraph.AddVertice()
	adjListGraph.AddEdge(vId1, vId2)
	adjListGraph.AddEdge(vId3, vId4)
	adjListGraph.AddEdge(vId4, vId3)
	adjListGraph.AddEdge(vId4, vId4)
	adjListGraph.AddEdge(vId4, vId2)

	adjListGraph.BindDataVertice(graph.MakeLabelData("вершина0"), vId1)

	presenter.DotGraph(adjListGraph)
}