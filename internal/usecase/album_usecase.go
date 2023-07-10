package usecase

import (
	"qvtec/go-app/internal/domain"
	"qvtec/go-app/internal/repository"
)

type AlbumUseCase interface {
	GetAll() ([]*domain.Album, error)
	Create(album *domain.Album) error
	GetByID(id int) (*domain.Album, error)
	Update(album *domain.Album) error
	Delete(id int) error
}

type albumUseCase struct {
	albumRepository repository.AlbumRepository
}

func NewAlbumUseCase(albumRepository repository.AlbumRepository) AlbumUseCase {
	return &albumUseCase{
		albumRepository: albumRepository,
	}
}

func (uc *albumUseCase) GetAll() ([]*domain.Album, error) {
	// @todo validation
	albums, err := uc.albumRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (uc *albumUseCase) Create(album *domain.Album) error {
	// @todo validation
	err := uc.albumRepository.Create(album)
	if err != nil {
		return err
	}
	return nil
}

func (uc *albumUseCase) GetByID(id int) (*domain.Album, error) {
	// @todo validation
	album, err := uc.albumRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (uc *albumUseCase) Update(album *domain.Album) error {
	// @todo validation
	err := uc.albumRepository.Update(album)
	if err != nil {
		return err
	}
	return nil
}

func (uc *albumUseCase) Delete(id int) error {
	// @todo validation
	err := uc.albumRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
