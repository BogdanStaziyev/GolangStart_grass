package event

type Event struct {
	Id   int64 		`db:"id" json:"id"`
	Name string		`db:"name" json:"name"`
}