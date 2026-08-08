package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tolu-c/go-shop/entity"
	"github.com/tolu-c/go-shop/service"
	"github.com/tolu-c/go-shop/validators"
)

type VideoHandlers interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type Handlers struct {
	service service.VideoService
}

var validate *validator.Validate

func New(s service.VideoService) VideoHandlers {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &Handlers{
		service: s,
	}
}

func (h *Handlers) FindAll() []entity.Video {
	return h.service.FindAll()
}

func (h *Handlers) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	h.service.Save(video)
	return nil
}

func (h *Handlers) ShowAll(ctx *gin.Context) {
	videos := h.service.FindAll()

	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
