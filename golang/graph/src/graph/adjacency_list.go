
package graph

import (
	"errors"
	"fmt"
)

// Представление графов с использованием списков смежности
type AdjacencyListGraph struct {
	// Списки смежности
	// verticeId => [verticeId1, .... verticeIdN]
	edges map[int]map[int]int

	// Отсортированные по Id вершины
	vertices map[int]int

	edgesData map[int]map[int]Data
	verticesData map[int]Data
}

func MakeAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		edges: make(map[int]map[int]int, 0),
		vertices: make(map[int]int, 0),
		edgesData: make(map[int]map[int]Data, 0),
		verticesData: make(map[int]Data, 0),
	}
}

// Граф пуст если в нем нет ни одной вершины
func (gr AdjacencyListGraph) IsEmpty() bool {
	return gr.VerticesCount() == 0
}

// Есть ли в графе вершина с таким Id
func (gr AdjacencyListGraph) HasVertice(verticeId int) bool {
	for _, vId := range gr.vertices {
		if verticeId == vId {
			return true
		}
	}
	return false
}

// Есть ли в графе ребро (если нет любой из вершин - будет false)
func (gr AdjacencyListGraph) HasEdge(fromVerticeId, toVerticeId int) bool {
	for _, vId := range gr.vertices {
		if fromVerticeId == vId {
			for _, vvId := range gr.edges[vId] {
				if toVerticeId == vvId {
					return true
				}
			}
		}
	}
	return false
}

// Количество вершин
func (gr AdjacencyListGraph) VerticesCount() int {
	return len(gr.vertices)
}

// Количество ребер
func (gr AdjacencyListGraph) EdgesCount() int {
	count := 0
	for _, list := range gr.edges {
		count = count + len(list)
	}
	return count
}

// Количество входящих в указанную вершину ребер.
func (gr AdjacencyListGraph) InEdgesCount(verticeId int) int {
	count := 0
	for _, list := range gr.edges {
		for _, toVerticeId := range list {
			if toVerticeId == verticeId {
				count = count + 1
				break
			}
		}
	}
	return count
}

// Количество исходящих из указанной вершины ребер
func (gr AdjacencyListGraph) OutEdgesCount(verticeId int) int {
	count := 0
	for fromVerticeId, list := range gr.edges {
		if fromVerticeId == verticeId {
			count = len(list)
			break
		}
	}
	return count
}

// Добавить вершину
func (gr *AdjacencyListGraph) AddVertice() int {
	l := len(gr.vertices)
	id := 0
	if l != 0 {
		id = gr.vertices[l - 1] + 1
	}
	
	gr.vertices[id] = id
	gr.edges[id] = make(map[int]int, 0)
	return id
}

// Добавить ребро
// error, _ если нет одной из вершин
// nil, false - если ребро уже есть
// nil, true - если ребро создал
func (gr AdjacencyListGraph) AddEdge(fromVerticeId, toVerticeId int) (error, bool) {
	if !gr.HasVertice(toVerticeId) {
		message := fmt.Sprintf(
			"Не существует вершины %d, поэтому не могу добавить ребро %d -> %d", 
			toVerticeId, 
			fromVerticeId, 
			toVerticeId,
		)
		return errors.New(message), false
	}
	for verticeId, verticeIdList := range gr.edges {
		if fromVerticeId == verticeId {
			for _, vId := range verticeIdList {
				if toVerticeId == vId {
					return nil, false
				}
			}
			gr.edges[fromVerticeId][toVerticeId] = toVerticeId
			return nil, true
		}
	}
	message := fmt.Sprintf(
		"Не существует вершины %d, поэтому не могу добавить ребро %d -> %d", 
		fromVerticeId, 
		fromVerticeId, 
		toVerticeId,
	)
	return errors.New(message), false
}

// Удалить ребра, которые заканчиваются в этой вершине, не удаляя ее саму
// false - вершины не существует 
func (gr AdjacencyListGraph) RemoveEdgesEnding(toVerticeId int) bool {
	if !gr.HasVertice(toVerticeId) {
		return false
	}

	for fromVerticeId, list := range gr.edges {
		found := false
		foundIdx := 0
		for idx, vId := range list {
			if vId == toVerticeId {
				found = true
				foundIdx = idx
				break
			}
		}
		if found {
			delete(gr.edges[fromVerticeId], foundIdx)
		}
	}

	return true
}

// Удалить вершину вместе с ребрами
// false - вершины не было
func (gr AdjacencyListGraph) RemoveVertice(verticeId int) bool {
	// Дропаем все ребра, входящие в эту вершину
	if !gr.RemoveEdgesEnding(verticeId) {
		return false
	}

	// Дропаем все ребра исходящие из этой вершины
	delete(gr.edges, verticeId)

	// Дропаем вершину
	found := false
	foundIdx := 0
	for idx, vId := range gr.vertices {
		if vId == verticeId {
			found = true
			foundIdx = idx
			break
		}
	}

	if found {
		delete(gr.vertices, foundIdx)
	}

	return true
}

// Удалить ребро
// false - ребра или одной из вершин не существует
func (gr AdjacencyListGraph) RemoveEdge(fromVerticeId, toVerticeId int) bool {
	if !gr.HasVertice(fromVerticeId) {
		return false
	}
	if !gr.HasVertice(toVerticeId) {
		return false
	}

	found := false
	foundIdx := 0
	for idx, vId := range gr.edges[fromVerticeId] {
		if vId == toVerticeId {
			found = true
			foundIdx = idx
			break
		}
	}
	if found {
		delete(gr.edges[fromVerticeId], foundIdx)
	}

	return found
}

// Привязать к вершине какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если вершины нет
func (gr AdjacencyListGraph) BindDataVertice(data Data, verticeId int) bool {
	if !gr.HasVertice(verticeId) {
		return false
	}
	gr.verticesData[verticeId] = data
	return true
}

// Получить привязанные к вершине данные
// false - вершина не существует
// bool, nil - данные к вершине не привязаны
func (gr AdjacencyListGraph) GetDataVertice(verticeId int) (bool, Data) {
	if !gr.HasVertice(verticeId) {
		return false, nil
	}
	return true, gr.verticesData[verticeId]
}

// Привязать к ребру какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если ребра или вершин нет
func (gr AdjacencyListGraph) BindDataEdge(data Data, fromVerticeId, toVerticeId int) bool {
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
func (gr AdjacencyListGraph) GetDataEdge(fromVerticeId, toVerticeId int) (bool, Data) {
	if !gr.HasEdge(fromVerticeId, toVerticeId) {
		return false, nil
	}
	return true, gr.edgesData[fromVerticeId][toVerticeId]
}

// Перебор всех ребер и узлов графа (это не поиск в ширину или глубину)
func (gr AdjacencyListGraph) Each(verticeVisitor VerticeVisitor, edgeVisitor EdgeVisitor) {
	if verticeVisitor != nil {
		for _, verticeId := range gr.vertices {
			_, data := gr.GetDataVertice(verticeId)
			verticeVisitor.Visit(verticeId, data)
		}
	}
	
	if edgeVisitor != nil {
		for fromVerticeId, list := range gr.edges {
			for _, toVerticeId := range list {
				_, data := gr.GetDataEdge(fromVerticeId, toVerticeId)
				edgeVisitor.Visit(fromVerticeId, toVerticeId, data)
			}
		}
	}
}
