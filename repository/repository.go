package repository

import "github.com/google/uuid"
import "GoRepository/model"

func NewRepository() Repository {
	return CacheRepository{make(map[uuid.UUID]model.Model)}
}

type Repository interface {

	Find(uuid uuid.UUID) model.Model
	Delete(uuid uuid.UUID)
	Save(model model.Model)

}

type CacheRepository struct {
	Cache map[uuid.UUID]model.Model
}

func (repository CacheRepository) Find(uuid uuid.UUID) (model model.Model) {
	model = repository.Cache[uuid]
	return model
}

func (repository CacheRepository) Delete(uuid uuid.UUID) {
	delete(repository.Cache, uuid)
}

func (repository CacheRepository) Save(model model.Model) {
	repository.Cache[model.Id()] = model
}