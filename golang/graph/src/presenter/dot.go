
package presenter 

import (
	"fmt"
	"graph"
)

// Представление графа для вывода на консоль в формате dot для Graphviz
type DotEdgeVisitor struct {
	graph graph.Graph
}

type DotPresenter struct {
	edge DotEdgeVisitor
}

func MakeDotEdgeVisitor(graph graph.Graph) DotEdgeVisitor {
	return DotEdgeVisitor{graph: graph}
}

func MakeDotPresenter(graph graph.Graph) DotPresenter {
	return DotPresenter{
		edge: MakeDotEdgeVisitor(graph),
	}
}

func (visitor DotEdgeVisitor) Visit(fromVerticeId, toVerticeId int, data graph.Data) {
	if visitor.graph != nil {
		_, fromData := visitor.graph.GetDataVertice(fromVerticeId)
		_, toData := visitor.graph.GetDataVertice(toVerticeId)
		if fromData != nil && fromData.Get() != nil {
			fmt.Printf("%s", fromData.Get())
		} else {
			fmt.Printf("vertice%d", fromVerticeId)
		}
		
		if toData != nil && toData.Get() != nil {
			fmt.Printf(" -> %s;\n", toData.Get())
		} else {
			fmt.Printf(" -> vertice%d;\n", toVerticeId)
		}

	} else {
		fmt.Printf("vertice%d -> vertice%d;\n", fromVerticeId, toVerticeId)
	}
}

func (presenter DotPresenter) Present(gr graph.Graph) {
	fmt.Print("digraph G {\n")
	gr.Each(nil, presenter.edge)
	fmt.Print("}\n")
}

func DotGraph(gr graph.Graph) {
	dot := MakeDotPresenter(gr)
	dot.Present(gr)
}
