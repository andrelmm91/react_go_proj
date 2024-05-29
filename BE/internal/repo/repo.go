package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	// get
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	AllMovies() ([]*models.Movie, error)
	AllGenres() ([]*models.Genre, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
	// create
	InsertMovie(movie models.Movie) (int, error)
	// edit
	UpdateMovie(movie models.Movie) error
	UpdateMovieGenres(id int, genreIDs []int) error
	// delete
	DeleteMovie(id int) error 
}
