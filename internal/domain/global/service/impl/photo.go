package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/model"
)

func (s *GlobalService) GetAllPhotos(ctx context.Context, user *authutil.UserClaims) (resp []*dto.PhotosRow, err error) {
	data := make([]*model.Photo, 0)

	if err = s.db.WithContext(ctx).Preload("User").Model(&model.Photo{}).Where("user_id = ?", user.ID).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.PhotosRow, 0)

	for _, v := range data {
		resp = append(resp, &dto.PhotosRow{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoURL:  v.PhotoUrl,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: &dto.UserPhotoRow{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return resp, nil
}

func (s *GlobalService) GetPhotosById(ctx context.Context, id int) (resp *dto.PhotosRow, err error) {
	v, err := s.globalRepository.FindPhotosById(ctx, id)
	if err != nil {
		log.Errorf("err get Photos paginated")
		return
	}
	resp = &dto.PhotosRow{
		ID:        v.ID,
		Title:     v.Title,
		Caption:   v.Caption,
		PhotoURL:  v.PhotoUrl,
		UserID:    v.UserID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		User: &dto.UserPhotoRow{
			Email:    v.User.Email,
			Username: v.User.Username,
		},
	}
	return
}

func (s *GlobalService) CreatePhotos(ctx context.Context, payload *dto.PayloadPhotos, user *authutil.UserClaims) (resp *dto.PhotoCreteResponse, err error) {
	data, err := s.globalRepository.CreatePhotos(ctx, &model.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoURL,
		UserID:   user.ID,
	})

	if err != nil {
		log.Errorf("err create Photos")
		return
	}

	resp = &dto.PhotoCreteResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoUrl,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdatePhotosById(ctx context.Context, id int, payload *dto.PayloadPhotos) (resp *dto.PhotoUpdateResponse, err error) {
	_, err = s.globalRepository.UpdatePhotosById(ctx, id, &model.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoURL,
	})

	if err != nil {
		log.Errorf("err update Photos %d", id)
		return
	}

	row, err := s.globalRepository.FindPhotosById(ctx, id)
	if err != nil {
		log.Errorf("err get Photos %d", id)
		return
	}

	resp = &dto.PhotoUpdateResponse{
		ID:        row.ID,
		Title:     row.Title,
		Caption:   row.Caption,
		PhotoURL:  row.PhotoUrl,
		UserID:    row.UserID,
		UpdatedAt: row.UpdatedAt,
	}
	return
}

func (s *GlobalService) DeletePhotosById(ctx context.Context, id int) (resp *int64, err error) {
	resp, err = s.globalRepository.DeletePhotosById(ctx, id)
	if err != nil {
		log.Errorf("err delete Photos %d", id)
		return
	}
	return
}
