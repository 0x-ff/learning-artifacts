
package graph

// Представление данных в вершинах и ребрах графа в виде строковых меток
type LabelData struct {
	Label string
}

func (l LabelData) Get() interface{} {
	return l.Label
}

func MakeLabelData(label string) LabelData {
	return LabelData{Label: label}
}
