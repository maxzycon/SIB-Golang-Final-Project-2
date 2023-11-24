package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/httputil"
)

func (c *GlobalController) handlerCreatePhotos(f *fiber.Ctx) (err error) {
	payload := dto.PayloadPhotos{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Photos")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreatePhotos(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create Photos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdatePhotos(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Photos")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadPhotos{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Photos")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdatePhotosById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Photos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerGetPhotosById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Photos get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetPhotosById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Photos get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllPhotos(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllPhotos(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller Photos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeletePhotos(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Photos delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeletePhotosById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Photos delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Your photo has been successfully deleted")
}
