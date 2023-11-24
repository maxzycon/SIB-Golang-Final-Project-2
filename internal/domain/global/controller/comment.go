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

func (c *GlobalController) handlerCreateComment(f *fiber.Ctx) (err error) {
	payload := dto.PayloadComment{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Error("err parse body create Comment")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreateComment(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create Comment :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdateComment(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Comment")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateComment{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Comment")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateCommentById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Comment :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerGetAllComment(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllComment(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller Comment :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeleteComment(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Comment delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeleteCommentById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Comment delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Your comment has been successfully deleted")
}
