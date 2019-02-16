
package graph

import (
	"errors"
	"fmt"
)

// Представление графов с использованием матрицы иницидентности
// Этот тип графа должен быть без петель (то есть ты не сможешь создать ребро из вершины в себя же)
type IncidentMatrixGraph struct {
	matrix map[int]map[int]int
	edges map[int]bool
	edgesCounter int

	edgesData map[int]map[int]Data
	verticesData map[int]Data
}

func MakeIncidentMatrixGraph() *IncidentMatrixGraph {
	return &IncidentMatrixGraph {
		matrix: make(map[int]map[int]int, 0),
		edges: make(map[int]bool, 0),
		edgesCounter: 0,
		edgesData: make(map[int]map[int]Data, 0),
		verticesData: make(map[int]Data, 0),
	}
}

// Граф пуст если в нем нет ни одной вершины
func (gr IncidentMatrixGraph) IsEmpty() bool {
	return gr.VerticesCount() == 0
}

// Есть ли в графе вершина с таким Id
func (gr IncidentMatrixGraph) HasVertice(verticeId int) bool {
	_, has := gr.matrix[verticeId]
	return has
}

// Есть ли в графе ребро (если нет любой из вершин - будет false)
func (gr IncidentMatrixGraph) HasEdge(fromVerticeId, toVerticeId int) bool {
	if !gr.HasVertice(toVerticeId) {
		return false
	}
	if !gr.HasVertice(fromVerticeId) {
		return false
	}
	for fromEdgeId, fromStatus := range gr.matrix[fromVerticeId] {
		if fromStatus == -1 {
			for toEdgeId, toStatus := range gr.matrix[toVerticeId] {
				if toStatus == 1 && fromEdgeId == toEdgeId {
					return true
				}
			}
		}
	}
	return false
}

// Количество вершин
func (gr IncidentMatrixGraph) VerticesCount() int {
	return len(gr.matrix)
}

// Количество ребер
func (gr IncidentMatrixGraph) EdgesCount() int {
	return len(gr.edges)
}

// Количество входящих в указанную вершину ребер.
func (gr IncidentMatrixGraph) InEdgesCount(toVerticeId int) int {
	if !gr.HasVertice(toVerticeId) {
		return 0
	}
	count := 0
	for _, status := range gr.matrix[toVerticeId] {
		if status == 1 {
			count = count + 1
		}
	}
	return count
}

// Количество исходящих из указанной вершины ребер
func (gr IncidentMatrixGraph) OutEdgesCount(fromVerticeId int) int {
	if !gr.HasVertice(fromVerticeId) {
		return 0
	}
	count := 0
	for _, status := range gr.matrix[fromVerticeId] {
		if status == -1 {
			count = count + 1
		}
	}
	return count
}

// Добавить вершину
func (gr *IncidentMatrixGraph) AddVertice() int {
	maxVerticeId := UnknownVerticeId
	for verticeId, _ := range gr.matrix {
		if verticeId > maxVerticeId {
			maxVerticeId = verticeId
		}
	}
	gr.matrix[maxVerticeId + 1] = make(map[int]int, 0)
	return maxVerticeId + 1
}

// Добавить ребро
// error, _ если нет одной из вершин
// nil, false - если ребро уже есть
// nil, true - если ребро создал
func (gr IncidentMatrixGraph) AddEdge(fromVerticeId, toVerticeId int) (error, bool) {
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
	if fromVerticeId == toVerticeId {
		return errors.New("Представление графа матрицей инцидентности предполагает, что в нем не будет петель - ребер из вершины в себя же"), false
	}
	if gr.HasEdge(fromVerticeId, toVerticeId) {
		return nil, false
	}
	maxEdgeId := -1
	for edgeId, _ := range gr.edges {
		if edgeId > maxEdgeId {
			maxEdgeId = edgeId
		}
	}
	gr.edges[maxEdgeId + 1] = true
	gr.matrix[fromVerticeId][maxEdgeId + 1] = -1
	gr.matrix[toVerticeId][maxEdgeId + 1] = 1

	return nil, true
}

// Удалить ребра, которые заканчиваются в этой вершине, не удаляя ее саму
// false - вершины не существует 
func (gr IncidentMatrixGraph) RemoveEdgesEnding(toVerticeId int) bool {
	if !gr.HasVertice(toVerticeId) {
		return false
	}
	var foundEdgeList[]int = make([]int, 0)
	for edgeId, status := range gr.matrix[toVerticeId] {
		if status == 1 {
			foundEdgeList = append(foundEdgeList, edgeId)
			delete(gr.matrix[toVerticeId], edgeId)
		}
	}
	for _, edgeId := range foundEdgeList {
		for verticeId, edges := range gr.matrix {
			status, exists := edges[edgeId]
			if exists && status == -1 {
				delete(gr.matrix[verticeId], edgeId)
			}
		}
		delete(gr.edges, edgeId)
	}
	return true
}

// Удалить вершину вместе с ребрами
// false - вершины не было
func (gr IncidentMatrixGraph) RemoveVertice(verticeId int) bool {
	if !gr.HasVertice(verticeId) {
		return false
	}
	var foundEdgeList[]int = make([]int, 0)
	for edgeId, _ := range gr.matrix[verticeId] {
		foundEdgeList = append(foundEdgeList, edgeId)
	}
	for _, edgeId := range foundEdgeList {
		for vId, edges := range gr.matrix {
			_, exists := edges[edgeId]
			if exists {
				delete(gr.matrix[vId], edgeId)
			}
		}
		delete(gr.edges, edgeId)
	}
	delete(gr.matrix, verticeId)
	return true
}

// Удалить ребро
// false - ребра или одной из вершин не существует
func (gr IncidentMatrixGraph) RemoveEdge(fromVerticeId, toVerticeId int) bool {
	if !gr.HasVertice(toVerticeId) {
		return false
	}
	if !gr.HasVertice(fromVerticeId) {
		return false
	}

	for fromEdgeId, fromStatus := range gr.matrix[fromVerticeId] {
		if fromStatus == -1 {
			for toEdgeId, toStatus := range gr.matrix[toVerticeId] {
				if toStatus == 1 && toEdgeId == fromEdgeId {
					delete(gr.matrix[fromVerticeId], toEdgeId)
					delete(gr.matrix[toVerticeId], toEdgeId)
					delete(gr.edges, toEdgeId)
					return true
				}
			}
		}
	}
	return false
}

// Привязать к вершине какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если вершины нет
func (gr IncidentMatrixGraph) BindDataVertice(data Data, verticeId int) bool {
	if !gr.HasVertice(verticeId) {
		return false
	}
	gr.verticesData[verticeId] = data
	return true
}

// Получить привязанные к вершине данные
// false - вершина не существует
// bool, nil - данные к вершине не привязаны
func (gr IncidentMatrixGraph) GetDataVertice(verticeId int) (bool, Data) {
	if !gr.HasVertice(verticeId) {
		return false, nil
	}
	return true, gr.verticesData[verticeId]
}

// Привязать к ребру какие-то данные (при удалении ребер и вершин данные останутся в памяти)
// если уже что-то привязано - перепривязать
// false если ребра или вершин нет
func (gr IncidentMatrixGraph) BindDataEdge(data Data, fromVerticeId, toVerticeId int) bool {
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
func (gr IncidentMatrixGraph) GetDataEdge(fromVerticeId, toVerticeId int) (bool, Data) {
	if !gr.HasEdge(fromVerticeId, toVerticeId) {
		return false, nil
	}
	return true, gr.edgesData[fromVerticeId][toVerticeId]
}

// Перебор всех ребер и узлов графа (это не поиск в ширину или глубину)
func (gr IncidentMatrixGraph) Each(verticeVisitor VerticeVisitor, edgeVisitor EdgeVisitor) {
	if verticeVisitor != nil {
		for verticeId, _ := range gr.matrix {
			_, data := gr.GetDataVertice(verticeId)
			verticeVisitor.Visit(verticeId, data)
		}
	}	
	if edgeVisitor != nil {
		for edgeId, _ := range gr.edges {
			fromVerticeId, toVerticeId := gr.getVerticePath(edgeId)
			if toVerticeId > UnknownVerticeId && fromVerticeId > UnknownVerticeId {
				_, data := gr.GetDataEdge(fromVerticeId, toVerticeId)
				edgeVisitor.Visit(fromVerticeId, toVerticeId, data)
			}
		}
	}
}

func (gr IncidentMatrixGraph) getVerticePath(edgeId int) (int, int) {
	fromVerticeId := UnknownVerticeId
	toVerticeId := UnknownVerticeId

	for verticeId, edges := range gr.matrix {
		if toVerticeId > UnknownVerticeId && fromVerticeId > UnknownVerticeId {
			return fromVerticeId, toVerticeId
		}
		for eId, status := range edges {
			if eId == edgeId {
				if status == 1 {
					toVerticeId = verticeId
				}
				if status == -1 {
					fromVerticeId = verticeId
				}
			}
		}
	}
	return fromVerticeId, toVerticeId
}
