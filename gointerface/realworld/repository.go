package realworld

type Repository interface {
	QueryLang(uint) Langeuage
}

type repo struct {
	data map[uint]Langeuage
}

func (r repo) QueryLang(i uint) Langeuage {
	return r.data[i]
}

func NewStaticRepository() Repository {
	return repo{data: RawData}
}
