package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/model"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/util/pagination"
)

func (r *GlobalRepository) FindPhotosById(ctx context.Context, id int) (resp *model.Photo, err error) {
	tx := r.db.WithContext(ctx).First(&resp, id)
	return resp, tx.Error
}

func (r *GlobalRepository) FindAllPhotos(ctx context.Context) (resp []*model.Photo, err error) {
	resp = make([]*model.Photo, 0)
	tx := r.db.WithContext(ctx).Model(&model.Photo{}).Find(&resp)
	return resp, tx.Error
}

func (r *GlobalRepository) FindPhotosPaginated(ctx context.Context, payload *pagination.DefaultPaginationPayload) (resp pagination.DefaultPagination, err error) {
	var Photos []*model.Photo = make([]*model.Photo, 0)
	sql := r.db.WithContext(ctx)
	if payload.Search != nil && *payload.Search != "" {
		search := fmt.Sprintf("%%%s%%", *payload.Search)
		sql = sql.Where("name LIKE ?", search)
	}
	sql.Scopes(payload.Pagination(&Photos, &resp.Paginator, sql)).Find(&Photos)
	resp.Items = Photos
	return
}

func (r *GlobalRepository) CreatePhotos(ctx context.Context, entity *model.Photo) (resp *model.Photo, err error) {
	tx := r.db.WithContext(ctx).Model(&model.Photo{}).Create(&entity)
	return entity, tx.Error
}

func (r *GlobalRepository) UpdatePhotosById(ctx context.Context, id int, entity *model.Photo) (resp *model.Photo, err error) {
	entity.ID = uint(id)
	tx := r.db.WithContext(ctx).Updates(entity)
	return entity, tx.Error
}

func (r *GlobalRepository) DeletePhotosById(ctx context.Context, id int) (resp *int64, err error) {
	tx := r.db.WithContext(ctx).Delete(&model.Photo{}, id)
	return &tx.RowsAffected, tx.Error
}
