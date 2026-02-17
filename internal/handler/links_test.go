package handler_test

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Pyth0nHater/link-shorter/internal/database"
	"github.com/Pyth0nHater/link-shorter/internal/handler"
	"github.com/Pyth0nHater/link-shorter/internal/models"
	"github.com/Pyth0nHater/link-shorter/internal/repository"
	"github.com/Pyth0nHater/link-shorter/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func getHash(link string) string {
    h := sha1.New()
    h.Write([]byte(link))

    return hex.EncodeToString(h.Sum(nil))[:10]
}

func TestHandler_CreateLink_Memory(t *testing.T){
	t.Parallel()

	mem:= database.InitMemory()
	linkRepo := repository.NewMemeoryLinkRepository(mem)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/create", linkHandler.CreateLink)

	t.Run("succses", func(t *testing.T) {
		t.Parallel()

		link:="vk.com"
		expectedHash := getHash(link)

		body, _ := json.Marshal(models.MainLink{MainLink: link})
		req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `"`+expectedHash+`"`, w.Body.String())
	})
}

func TestHandler_CreateLink_Postgres(t *testing.T){
	t.Parallel()

		err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	ctx := context.Background()	

	var linkRepo repository.LinkRepositoryInterface

	db := database.InitDB(ctx, dbUrl)
	linkRepo = repository.NewPostgresLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/create", linkHandler.CreateLink)

	t.Run("succses", func(t *testing.T) {
		t.Parallel()

		link:="vk.com"
		expectedHash := getHash(link)

		body, _ := json.Marshal(models.MainLink{MainLink: link})
		req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `"`+expectedHash+`"`, w.Body.String())
	})
}
