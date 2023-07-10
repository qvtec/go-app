package repository

import (
	"database/sql"
	"qvtec/go-app/internal/domain"
	"qvtec/go-app/pkg/db"
	"time"
)

type AlbumRepository interface {
	GetAll() ([]*domain.Album, error)
	GetByID(id int) (*domain.Album, error)
	Create(album *domain.Album) error
	Update(album *domain.Album) error
	Delete(id int) error
}

type mysqlAlbumRepository struct {
	DB *db.MySQLDB
}

func NewAlbumRepository(db *db.MySQLDB) AlbumRepository {
	return &mysqlAlbumRepository{
		DB: db,
	}
}

func (r *mysqlAlbumRepository) GetAll() ([]*domain.Album, error) {
	query := "SELECT id, title, contents FROM albums"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []*domain.Album{}
	for rows.Next() {
		album := &domain.Album{}
		err := rows.Scan(&album.ID, &album.Title, &album.Contents)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (r *mysqlAlbumRepository) Create(album *domain.Album) error {
	currentTime := time.Now().UTC()
	album.CreatedAt = currentTime
	album.UpdatedAt = currentTime
	query := "INSERT INTO albums (title, contents, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Execute(query, album.Title, album.Contents, album.CreatedAt, album.UpdatedAt)
	if err != nil {
		return err
	}

	albumID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	album.ID = int(albumID)

	return nil
}

func (r *mysqlAlbumRepository) GetByID(id int) (*domain.Album, error) {
	query := "SELECT id, title, contents FROM albums WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	album := &domain.Album{}
	err := row.Scan(&album.ID, &album.Title, &album.Contents)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAlbumNotFound
		}
		return nil, err
	}

	return album, nil
}

func (r *mysqlAlbumRepository) Update(album *domain.Album) error {
	currentTime := time.Now().UTC()
	album.UpdatedAt = currentTime
	query := "UPDATE albums SET title = ?, contents = ?, updated_at = ? WHERE id = ?"
	_, err := r.DB.Execute(query, album.Title, album.Contents, album.UpdatedAt, album.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlAlbumRepository) Delete(id int) error {
	query := "DELETE FROM albums WHERE id = ?"
	_, err := r.DB.Execute(query, id)
	if err != nil {
		return err
	}

	return nil
}
