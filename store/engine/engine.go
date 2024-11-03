package engine

import (
	"context"
	"database/sql"

	"github.com/briannkhata/car-zone/models"
)

type EngineStore struct {
	db *sql.DB
}

func new(db *sql.DB) EngineStore {
	return EngineStore{db: db}
}

func (s EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (s EngineStore) GetEngineByBrand(ctx context.Context, brand string, isEngine bool) {

}

func (s EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (s EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (s EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {

}
