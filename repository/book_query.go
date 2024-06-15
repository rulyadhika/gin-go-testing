package repository

const (
	findOneByIdQuery = `SELECT id, title, author FROM books WHERE id=$1`
	findAllQuery     = `SELECT id, title, author FROM books`
	createQuery      = `INSERT INTO books(title, author) VALUES($1,$2) RETURNING id`
)
