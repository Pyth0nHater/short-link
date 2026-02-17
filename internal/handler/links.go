package handler

import (
	"fmt"
	"net/http"

	"github.com/Pyth0nHater/link-shorter/internal/models"
	"github.com/Pyth0nHater/link-shorter/internal/service"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LinkHandler struct {
	linkService *service.LinkService
}

func NewLinkHandler(linkService *service.LinkService) *LinkHandler {
	return &LinkHandler{linkService}
}

func (h *LinkHandler) CreateLink(gc *gin.Context) {
	var req models.MainLink
	
	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validation.ValidateStruct(&req,
        validation.Field(&req.MainLink, 
            validation.Required,     
            validation.Length(5, 2048), 
            is.URL,                    
        ),
    )

    if err != nil {
        gc.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
        return
    }

	link, err := h.linkService.CreateLink(gc.Request.Context(), &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusCreated, link)
}


func (h *LinkHandler) GetLink(gc *gin.Context) {
	var req models.ShortLink

	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := models.GetLinkFilter{
		MainLink: nil,
		ShortLink: &req.ShortLink,
	}

	link, err := h.linkService.GetLink(gc.Request.Context(), res)
	fmt.Print(link.MainLink)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	gc.JSON(http.StatusCreated, link.MainLink)
}
