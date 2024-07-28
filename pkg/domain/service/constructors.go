// Code generated by tools/update-providers.go; DO NOT EDIT.
package service

import "github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"

func ProvideBlocksService(
	blocksRepo repository.BlocksRepository,
) BlocksService {
	return &blocksServiceImpl{
		blocksRepo: blocksRepo,
	}
}

func ProvideCategoriesService(
	categoriesRepository repository.CategoriesRepository,
) CategoriesService {
	return &categoriesServiceImpl{
		categoriesRepository: categoriesRepository,
	}
}

func ProvideIspService(
	repo repository.IspRepository,
	txHandler repository.TxHandler,
) IspService {
	return &ispServiceImpl{
		repo: repo,

		txHandler: txHandler,
	}
}

func ProvideSitesService(
	sitesRepository repository.SitesRepository,
	blocksRepository repository.BlocksRepository,
	txHandler repository.TxHandler,
) SitesService {
	return &sitesServiceImpl{
		sitesRepository: sitesRepository,

		blocksRepository: blocksRepository,

		txHandler: txHandler,
	}
}
