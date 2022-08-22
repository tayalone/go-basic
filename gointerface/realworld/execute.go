package realworld

import "fmt"

// type Langeuage struct {
// 	ID   uint
// 	Name string
// }

// type Repository interface {
// 	QueryLang(uint) Langeuage
// }

// type Handler struct {
// 	repo Repository
// }

// func NewHandler(repo Repository) *Handler {
// 	return &Handler{repo: repo}
// }

// func (h *Handler) Do(id uint) string {
// 	lang := h.repo.QueryLang(id)
// 	return fmt.Sprintf("%s language", lang.Name)
// }

// var data = map[uint]Langeuage{
// 	1: {ID: 1, Name: "GO"},
// 	2: {ID: 2, Name: "Type Script"},
// }

// type repo struct {
// 	data map[uint]Langeuage
// }

// func (r repo) QueryLang(i uint) Langeuage {
// 	return r.data[i]
// }

// func NewStaticRepository() Repository {
// 	return repo{data: data}
// }

/*
Execute is Represent How Gophers use Inter face in real-world
*/
func Execute() {
	/* Repository Just Like Db Connector */
	repo := NewStaticRepository()
	/* Like a Controller ใน Mvc */
	h := NewHandler(repo)
	s := h.Do(2)
	fmt.Println(s)
}
