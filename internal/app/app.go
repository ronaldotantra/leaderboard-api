package app

import (
	"context"
)

type Application struct {
	Storages     *Storages
	Repositories *Repositories
	Services     *Services
}

type ApplicationParams struct {
	StorageParams StoragesParams
}

func SetupApp(ctx context.Context, params ApplicationParams) (*Application, error) {
	storages, err := SetupStorages(params.StorageParams)
	if err != nil {
		return nil, err
	}

	repositories := SetupRepositories(storages)
	services := SetupServices(ctx, storages, repositories)

	return &Application{
		Storages:     storages,
		Repositories: repositories,
		Services:     services,
	}, nil
}
