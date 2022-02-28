package sqlhelper

const (
	FindAll = "select id, title, description, checked, created_at, updated_at from notes"
	FindOne = "select id, title, description, checked, created_at, updated_at from notes where id ="
)
