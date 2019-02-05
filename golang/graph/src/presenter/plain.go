
package presenter 

import (
	"fmt"
	"graph"
)

// Представление графа для вывода на консоль обычным текстом
type PlainVerticeVisitor struct {
	presenter DataPresenter
}

type PlainEdgeVisitor struct {
	presenter DataPresenter
}

type PlainPresenter struct {
	vertice PlainVerticeVisitor
	edge PlainEdgeVisitor
}

type LabelDataPlainPresenter struct {
}

func MakeLabelDataPlainPresenter() LabelDataPlainPresenter {
	return LabelDataPlainPresenter{}
}

func (presenter LabelDataPlainPresenter) Present(data graph.Data) {
	if data != nil && data.Get() != nil {
		fmt.Print(data.Get())
	}
}

func MakePlainVerticeVisitor(dataPresenter DataPresenter) PlainVerticeVisitor {
	return PlainVerticeVisitor{presenter: dataPresenter}
}

func MakePlainEdgeVisitor(dataPresenter DataPresenter) PlainEdgeVisitor {
	return PlainEdgeVisitor{presenter: dataPresenter}
}

func MakePlainPresenter(verticeDataPresenter, edgeDataPresenter DataPresenter) PlainPresenter {
	return PlainPresenter{
		vertice: MakePlainVerticeVisitor(verticeDataPresenter), 
		edge: MakePlainEdgeVisitor(edgeDataPresenter),
	}
}

func (visitor PlainVerticeVisitor) Visit(verticeId int, data graph.Data) {
	fmt.Printf("vertice %d ", verticeId)
	visitor.presenter.Present(data)
	fmt.Print("\n")
}

func (visitor PlainEdgeVisitor) Visit(fromVerticeId, toVerticeId int, data graph.Data) {
	fmt.Printf("edge %d -> %d ", fromVerticeId, toVerticeId)
	visitor.presenter.Present(data)
	fmt.Print("\n")
}

func (presenter PlainPresenter) Present(gr graph.Graph) {
	gr.Each(presenter.vertice, presenter.edge)
}

func PrintGraph(gr graph.Graph) {
	dataPresenter := MakeLabelDataPlainPresenter()
	plain := MakePlainPresenter(dataPresenter, dataPresenter)
	plain.Present(gr)
}
