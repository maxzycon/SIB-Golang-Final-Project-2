package repository

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/model"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/util/pagination"
)

type GlobalRepository interface {
	// ---- Photos
	FindPhotosById(ctx context.Context, id int) (resp *model.Photo, err error)
	FindAllPhotos(ctx context.Context) (resp []*model.Photo, err error)
	FindPhotosPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error)
	CreatePhotos(ctx context.Context, entity *model.Photo) (resp *model.Photo, err error)
	UpdatePhotosById(ctx context.Context, id int, entity *model.Photo) (resp *model.Photo, err error)
	DeletePhotosById(ctx context.Context, id int) (resp *int64, err error)
}
