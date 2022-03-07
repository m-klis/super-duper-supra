package sqlhelper

const (
	FindAll = `SELECT 
			   id, title, description, checked, created_at, updated_at 
			   FROM 
			   notes
			   ORDER BY
			   id`

	FindOne = `SELECT 
			   id, title, description, checked, created_at, updated_at 
			   FROM 
			   notes 
			   WHERE 
			   id = `

	CreateNote = `INSERT INTO notes 
				  (title, description, checked) 
				  VALUES 
				  (:title, :description, :checked) 
				  RETURNING 
				  id, created_at, updated_at`

	UpdateNote = `UPDATE 
				  notes 
				  SET 
				  title=$1, description=$2, checked=$3, updated_at=now() 
				  WHERE 
				  id=$4 
				  RETURNING 
				  created_at, updated_at`

	DeleteNote = `DELETE FROM 
				  notes 
				  WHERE 
				  id = $1`
)

// `UPDATE user SET first_name=:first, last_name=:last WHERE first_name = 'Bin'
