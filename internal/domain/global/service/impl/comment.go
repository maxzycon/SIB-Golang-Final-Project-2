package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllComment(ctx context.Context, user *authutil.UserClaims) (resp []*dto.CommentRow, err error) {
	data := make([]*model.Comment, 0)

	if err = s.db.WithContext(ctx).Preload("User").Preload("Photo").Model(&model.Comment{}).Where("user_id = ?", user.ID).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.CommentRow, 0)

	for _, v := range data {
		resp = append(resp, &dto.CommentRow{
			ID:        v.ID,
			UserID:    v.UserID,
			Message:   v.Message,
			PhotoID:   v.PhotoID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: &dto.UserCommentRow{
				ID:       v.UserID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: &dto.PhotoCommentRow{
				ID:       v.PhotoID,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoURL: v.Photo.PhotoUrl,
				UserID:   v.Photo.UserID,
			},
		})
	}

	return resp, nil
}

func (s *GlobalService) CreateComment(ctx context.Context, payload *dto.PayloadComment, user *authutil.UserClaims) (resp *dto.CommentCreteResponse, err error) {
	entity := &model.Comment{
		UserID:  user.ID,
		PhotoID: payload.PhotoID,
		Message: payload.Message,
	}
	if err = s.db.Create(&entity).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err create Comment")
		return
	}

	resp = &dto.CommentCreteResponse{
		ID:        entity.ID,
		Message:   entity.Message,
		PhotoID:   entity.PhotoID,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdateCommentById(ctx context.Context, id int, payload *dto.PayloadUpdateComment) (resp *dto.CommentUpdateResponse, err error) {
	entity := &model.Comment{
		Model: gorm.Model{
			ID: uint(id),
		},
		Message: payload.Message,
	}
	if err = s.db.Updates(&entity).Error; err != nil {
		return
	}

	data := &model.Comment{}
	if err = s.db.Preload("Photo").Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.CommentUpdateResponse{
		ID:        data.ID,
		Title:     data.Photo.Title,
		Caption:   data.Photo.Caption,
		PhotoURL:  data.Photo.Caption,
		UpdatedAt: data.UpdatedAt,
	}
	return
}

func (s *GlobalService) DeleteCommentById(ctx context.Context, id int) (resp *int64, err error) {
	if err = s.db.Delete(&model.Comment{}, id).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err delete Comment %d", id)
		return
	}
	return
}
