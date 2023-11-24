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

func (c *GlobalController) handlerCreateSocialMedia(f *fiber.Ctx) (err error) {
	payload := dto.PayloadSocialMedia{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Error("err parse body create SocialMedia")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreateSocialMedia(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create SocialMedia :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdateSocialMedia(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update SocialMedia")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateSocialMedia{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update SocialMedia")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateSocialMediaById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update SocialMedia :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerGetAllSocialMedia(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllSocialMedia(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller SocialMedia :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeleteSocialMedia(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params SocialMedia delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeleteSocialMediaById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller SocialMedia delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Your social media has been successfully deleted")
}
