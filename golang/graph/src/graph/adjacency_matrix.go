
package graph

import (
	"errors"
	"fmt"
)

// Представление графов с использованием матрицы смежности
type AdjacencyMatrixGraph struct {
	matrix map[int]map[int]bool

	edgesData map[int]map[int]Data
	verticesData map[int]Data
}

func MakeAdjacencyMatrixGraph() *AdjacencyMatrixGraph {
	return &AdjacencyMatrixGraph{
		matrix: make(map[int]map[int]bool, 0),
		edgesData: make(map[int]map[int]Data, 0),
		verticesData: make(map[int]Data, 0),
	}
}

// Граф пуст если в нем нет ни одной вершины
func (gr AdjacencyMatrixGraph) IsEmpty() bool {
	return gr.VerticesCount() == 0
}

// Есть ли в графе вершина с таким Id
func (gr AdjacencyMatrixGraph) HasVertice(verticeId int) bool {
	_, has := gr.matrix[verticeId]
	return has
}

// Есть ли в графе ребро (если нет любой из вершин - будет false)
func (gr AdjacencyMatrixGraph) HasEdge(fromVerticeId, toVerticeId int) bool {
	list, hasFromVertice := gr.matrix[fromVerticeId]
	if hasFromVertice && list != nil {
		has, exists := list[toVerticeId]
		return exists && has
	}
	return false
}

// Количество вершин
func (gr AdjacencyMatrixGraph) VerticesCount() int {
	return len(gr.matrix)
}

// Количество ребер
func (gr AdjacencyMatrixGraph) EdgesCount() int {
	count := 0
	for _, list := range gr.matrix {
		for _, hasEdge := range list {
			if hasEdge {
				count = count + 1
			}
		}
	}
	return count
}

// Добавить вершину
func (gr *AdjacencyMatrixGraph) AddVertice() int {
	maxVerticeId := -1
	for verticeId, _ := range gr.matrix {
		if verticeId > maxVerticeId {
			maxVerticeId = verticeId
		}
	}
	gr.matrix[maxVerticeId + 1] = make(map[int]bool, 0)
	return maxVerticeId + 1
}

// Добавить ребро
// error, _ если нет одной из вершин
// nil, false - если ребро уже есть
// nil, true - если ребро создал
func (gr AdjacencyMatrixGraph) AddEdge(fromVerticeId, toVerticeId int) (error, bool) {
	if !gr.HasVertice(toVerticeId) {
		message := fmt.Sprintf(
			"Не существует вершины %d, поэтому не могу добавить ребро %d -> %d", 
			toVerticeId, 
			fromVerticeId, 
			toVerticeId,
		)
		return errors.New(message), false
	}
	if !gr.HasVertice(fromVerticeId) {
		message := fmt.Sprintf(
			"Не существует вершины %d, поэтому не могу добавить ребро %d -> %d", 
			fromVerticeId, 
			fromVerticeId, 
			toVerticeId,
		)
		return errors.New(message), false
	}
	has, exists := gr.matrix[fromVerticeId][toVerticeId]
	if exists && has {
		return nil, false
	} else {
		gr.matrix[fromVerticeId][toVerticeId] = true
		return nil, true
	}
}

// Удалить ребра, которые заканчиваются в этой вершине, не удаляя ее саму
// false - вершины не существует 
func (gr AdjacencyMatrixGraph) RemoveEdgesEnding(toVerticeId int) bool {
	if !gr.HasVertice(toVerticeId) {
		return false
	}

	for fromVerticeId, list := range gr.matrix {
		_, exists := list[toVerticeId]
		if exists {
			gr.matrix[fromVerticeId][toVerticeId] = false
		}
	}

	return true
}

// Удалить вершину вместе с ребрами
// false - вершины не было
func (gr AdjacencyMatrixGraph) RemoveVertice(verticeId int) bool {
	// Дропаем все ребра, входящие в эту вершину
	if !gr.RemoveEdgesEnding(verticeId) {
		return false
	}

	// Дропаем все ребра исходящие из этой вершины и вершину
	delete(gr.matrix, verticeId)

	return true
}

// Удалить ребро
// false - ребра или одной из вершин не существует
func (gr AdjacencyMatrixGraph) RemoveEdge(fromVerticeId, toVerticeId int) bool {
	if !gr.HasVertice(fromVerticeId) {
		return false
	}
	if !gr.HasVertice(toVerticeId) {
		return false
	}

	has, exists := gr.matrix[fromVerticeId][toVerticeId]
	if exists && has {
		gr.matrix[fromVerticeId][toVerticeId] = false
	}
	return exists && has
}

// Привязать к вершине какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если вершины нет
func (gr AdjacencyMatrixGraph) BindDataVertice(data Data, verticeId int) bool {
	if !gr.HasVertice(verticeId) {
		return false
	}
	gr.verticesData[verticeId] = data
	return true
}

// Получить привязанные к вершине данные
// false - вершина не существует
// bool, nil - данные к вершине не привязаны
func (gr AdjacencyMatrixGraph) GetDataVertice(verticeId int) (bool, Data) {
	if !gr.HasVertice(verticeId) {
		return false, nil
	}
	return true, gr.verticesData[verticeId]
}

// Привязать к ребру какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если ребра или вершин нет
func (gr AdjacencyMatrixGraph) BindDataEdge(data Data, fromVerticeId, toVerticeId int) bool {
	if !gr.HasEdge(fromVerticeId, toVerticeId) {
		return false
	}
	if gr.edgesData[fromVerticeId] == nil {
		gr.edgesData[fromVerticeId] = make(map[int]Data, 0)
	}
	gr.edgesData[fromVerticeId][toVerticeId] = data
	return true
}

// Получить привязанные к ребру данные
// false - вершины или ребро не существует
// bool, nil - данные к ребру не привязаны
func (gr AdjacencyMatrixGraph) GetDataEdge(fromVerticeId, toVerticeId int) (bool, Data) {
	if !gr.HasEdge(fromVerticeId, toVerticeId) {
		return false, nil
	}
	return true, gr.edgesData[fromVerticeId][toVerticeId]
}

// Перебор всех ребер и узлов графа (это не поиск в ширину или глубину)
func (gr AdjacencyMatrixGraph) Each(verticeVisitor VerticeVisitor, edgeVisitor EdgeVisitor) {
	for fromVerticeId, list := range gr.matrix {
		if verticeVisitor != nil {
			_, data := gr.GetDataVertice(fromVerticeId)
			verticeVisitor.Visit(fromVerticeId, data)
		}
		if edgeVisitor != nil {
			for toVerticeId, has := range list {
				if has {
					_, data := gr.GetDataEdge(fromVerticeId, toVerticeId)
					edgeVisitor.Visit(fromVerticeId, toVerticeId, data)
				}
			}
		}
	}
}
