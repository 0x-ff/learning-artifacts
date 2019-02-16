
package graph

const UnknownVerticeId = -1

// Данные привязанные к вершине или ребру
type Data interface {
	Get() interface{}
}

// Интерфейс посетителя вершины 
type VerticeVisitor interface {
	Visit(verticeId int, data Data)
}

// Интерфейс посетителя ребра
type EdgeVisitor interface {
	Visit(fromVerticeId, toVerticeId int, data Data)
}

// Общий интерфейс для представления любых ориентированных графов
// Вершины нумеруются айдишниками от 0 и далее
// Следующий айди получается автоинкрементом самого большого айдишника
type Graph interface {

	// Граф пуст если в нем нет ни одной вершины
	IsEmpty() bool

	// Есть ли в графе вершина с таким Id
	HasVertice(verticeId int) bool

	// Есть ли в графе ребро (если нет любой из вершин - будет false)
	HasEdge(fromVerticeId, toVerticeId int) bool

	// Количество вершин
	VerticesCount() int

	// Количество ребер
	EdgesCount() int

	// Количество входящих в указанную вершину ребер.
	InEdgesCount(verticeId int) int

	// Количество исходящих из этой вершины ребер.
	OutEdgesCount(verticeId int) int

	// Добавить вершину
	AddVertice() int

	// Добавить ребро 
	// error, _ если нет одной из вершин
	// nil, false - если ребро уже есть
	// nil, true - если ребро создал
	AddEdge(fromVerticeId, toVerticeId int) (error, bool)

	// Удалить вершину вместе с исходящими и входящими ребрами
	// false - вершины не было
	RemoveVertice(verticeId int) bool

	// Удалить ребра, которые заканчиваются в этой вершине, не удаляя ее саму
	// false - вершины не было
	RemoveEdgesEnding(toVerticeId int) bool

	// Удалить ребро
	// false - ребра не было
	RemoveEdge(fromVerticeId, toVerticeId int) bool

	// Привязать к вершине какие-то данные
	// если уже что-то привязано - перепривязать
	// false если вершины нет
	BindDataVertice(data Data, verticeId int) bool 

	// Получить привязанные к вершине данные
	// false - вершина не существует
	// bool, nil - данные к вершине не привязаны
	GetDataVertice(verticeId int) (bool, Data)

	// Привязать к ребру какие-то данные
	// если уже что-то привязано - перепривязать
	// false если ребра или вершин нет
	BindDataEdge(data Data, fromVerticeId, toVerticeId int) bool 

	// Получить привязанные к ребру данные
	// false - вершины или ребро не существует
	// bool, nil - данные к ребру не привязаны
	GetDataEdge(fromVerticeId, toVerticeId int) (bool, Data)

	// Перебор всех ребер и узлов графа (это не поиск в ширину или глубину)
	Each(verticeVisitor VerticeVisitor, edgeVisitor EdgeVisitor)
}
