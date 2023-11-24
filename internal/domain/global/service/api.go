package service

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/authutil"
)

type GlobalService interface {
	// ----- Photos
	GetAllPhotos(ctx context.Context, user *authutil.UserClaims) (resp []*dto.PhotosRow, err error)
	GetPhotosById(ctx context.Context, id int) (resp *dto.PhotosRow, err error)
	CreatePhotos(ctx context.Context, payload *dto.PayloadPhotos, user *authutil.UserClaims) (resp *dto.PhotoCreteResponse, err error)
	UpdatePhotosById(ctx context.Context, id int, payload *dto.PayloadPhotos) (resp *dto.PhotoUpdateResponse, err error)
	DeletePhotosById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Comments
	GetAllComment(ctx context.Context, user *authutil.UserClaims) (resp []*dto.CommentRow, err error)
	CreateComment(ctx context.Context, payload *dto.PayloadComment, user *authutil.UserClaims) (resp *dto.CommentCreteResponse, err error)
	UpdateCommentById(ctx context.Context, id int, payload *dto.PayloadUpdateComment) (resp *dto.CommentUpdateResponse, err error)
	DeleteCommentById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Social media
	GetAllSocialMedia(ctx context.Context, user *authutil.UserClaims) (resp *dto.SocialMediaWrappper, err error)
	CreateSocialMedia(ctx context.Context, payload *dto.PayloadSocialMedia, user *authutil.UserClaims) (resp *dto.SocialMediaCreteResponse, err error)
	UpdateSocialMediaById(ctx context.Context, id int, payload *dto.PayloadUpdateSocialMedia) (resp *dto.SocialMediaUpdateResponse, err error)
	DeleteSocialMediaById(ctx context.Context, id int) (resp *int64, err error)
}
