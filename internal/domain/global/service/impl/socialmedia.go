package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllSocialMedia(ctx context.Context, user *authutil.UserClaims) (resp *dto.SocialMediaWrappper, err error) {
	resp = &dto.SocialMediaWrappper{
		SocialMedias: make([]*dto.SocialMediaRow, 0),
	}
	data := make([]*model.SocialMedia, 0)

	if err = s.db.WithContext(ctx).Preload("User").Model(&model.SocialMedia{}).Where("user_id = ?", user.ID).Find(&data).Error; err != nil {
		return
	}

	for _, v := range data {
		resp.SocialMedias = append(resp.SocialMedias, &dto.SocialMediaRow{
			ID:             v.ID,
			UserID:         v.UserID,
			Name:           v.Name,
			SocialMediaURL: v.SocialMediaUrl,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User: &dto.UserSocialMediaRow{
				ID:       v.UserID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return resp, nil
}

func (s *GlobalService) CreateSocialMedia(ctx context.Context, payload *dto.PayloadSocialMedia, user *authutil.UserClaims) (resp *dto.SocialMediaCreteResponse, err error) {
	entity := &model.SocialMedia{
		UserID:         user.ID,
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaURL,
	}
	if err = s.db.Create(&entity).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err create SocialMedia")
		return
	}

	resp = &dto.SocialMediaCreteResponse{
		ID:             entity.ID,
		Name:           entity.Name,
		SocialMediaURL: entity.SocialMediaUrl,
		UserID:         entity.UserID,
		CreatedAt:      entity.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdateSocialMediaById(ctx context.Context, id int, payload *dto.PayloadUpdateSocialMedia) (resp *dto.SocialMediaUpdateResponse, err error) {
	entity := &model.SocialMedia{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:           payload.Name,
		SocialMediaUrl: payload.SocialMediaURL,
	}
	if err = s.db.Updates(&entity).Error; err != nil {
		return
	}

	data := &model.SocialMedia{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.SocialMediaUpdateResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaUrl,
		UpdatedAt:      data.UpdatedAt,
		UserID:         data.UserID,
	}
	return
}

func (s *GlobalService) DeleteSocialMediaById(ctx context.Context, id int) (resp *int64, err error) {
	if err = s.db.Delete(&model.SocialMedia{}, id).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err delete SocialMedia %d", id)
		return
	}
	return
}
