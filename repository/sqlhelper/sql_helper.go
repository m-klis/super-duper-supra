package sqlhelper

const (
	FindAll    = `select id, title, description, checked, created_at, updated_at from notes`
	FindOne    = `select id, title, description, checked, created_at, updated_at from notes where id =`
	CreateNote = `insert into notes (title, description, checked) values (:title, :description, :checked) returning id`
)
