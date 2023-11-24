package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/service"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/middleware"
)

const (
	PhotosBasePath   = "photos"
	PhotosBaseIdPath = "photos/:id"

	CommentsBasePath   = "comments"
	CommentsBaseIdPath = "comments/:id"
)

type GlobalControllerParams struct {
	V1            fiber.Router
	Conf          *config.Config
	GlobalService service.GlobalService
	Middleware    middleware.GlobalMiddleware
}
type GlobalController struct {
	v1            fiber.Router
	conf          *config.Config
	globalService service.GlobalService
	middleware    middleware.GlobalMiddleware
}

func New(params *GlobalControllerParams) *GlobalController {
	return &GlobalController{
		v1:            params.V1,
		conf:          params.Conf,
		globalService: params.GlobalService,
		middleware:    params.Middleware,
	}
}

func (pc *GlobalController) Init() {
	// ---- Photos API
	pc.v1.Get(PhotosBasePath, pc.middleware.Protected([]uint{}), pc.handlerGetAllPhotos)
	pc.v1.Get(PhotosBaseIdPath, pc.middleware.Protected([]uint{}), pc.handlerGetPhotosById)
	pc.v1.Post(PhotosBasePath, pc.middleware.Protected([]uint{}), pc.handlerCreatePhotos)
	pc.v1.Put(PhotosBaseIdPath, pc.middleware.Protected([]uint{}), pc.handlerUpdatePhotos)
	pc.v1.Delete(PhotosBaseIdPath, pc.middleware.Protected([]uint{}), pc.handlerDeletePhotos)

	// ---- Comments API
	pc.v1.Get(CommentsBasePath, pc.middleware.Protected([]uint{}), pc.handlerGetAllComment)
	pc.v1.Post(CommentsBasePath, pc.middleware.Protected([]uint{}), pc.handlerCreateComment)
	pc.v1.Put(CommentsBaseIdPath, pc.middleware.Protected([]uint{}), pc.handlerUpdateComment)
	pc.v1.Delete(CommentsBaseIdPath, pc.middleware.Protected([]uint{}), pc.handlerDeleteComment)
}
