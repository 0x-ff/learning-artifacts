
package presenter

import "graph"

// Интерфейс представления данных графа (например для вывода на экран или в файл в формате dot)
type DataPresenter interface {
	Present(data graph.Data)
}

// Интерфейс представления графа (например для вывода на экран или в файл в формате dot)
type Presenter interface {
	Present(
		gr graph.Graph,
		verticeDataPresenter DataPresenter,
		edgeDataPresenter DataPresenter,
	)
}
