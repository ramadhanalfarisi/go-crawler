package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-crawler/helpers"
	"github.com/ramadhanalfarisi/go-crawler/model"
)

func GetDataManga(g *gin.Context) {
	var mangas []model.Manga
	helpers.CrawlManga(&mangas)
	json, err := json.Marshal(mangas)
	if err != nil {
		fmt.Println(err)
	}
	g.JSON(200, json)
}
