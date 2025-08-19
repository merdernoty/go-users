package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/domain"
)

type AnimeController struct {
	animeService domain.AnimeService
}

func NewAnimeController(s domain.AnimeService) *AnimeController {
	return &AnimeController{animeService: s}
}

func (c *AnimeController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/anime/search", c.search)
	r.GET("/anime/:id", c.getById)
	r.GET("/anime", c.list)
}

func (c *AnimeController) search(ctx *gin.Context) {
	q := ctx.Query("q")
	res, err := c.animeService.Search(ctx, q)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *AnimeController) getById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	res, err := c.animeService.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *AnimeController) list(ctx *gin.Context) {
	res, err := c.animeService.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
