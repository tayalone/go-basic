package realworld

type Langeuage struct {
	ID   uint
	Name string
}

var RawData = map[uint]Langeuage{
	1: {ID: 1, Name: "GO"},
	2: {ID: 2, Name: "Type Script"},
}
