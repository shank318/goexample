package pingdom

import (
	"examplego/common"
	"examplego/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type URL = models.URL
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Url string `json:"url" binding:"required"`
		CrawlTimeout int `json:"crawl_timeout" binding:"required"`
		Frequency int `json:"frequency" binding:"required"`
		FailureThreshold int `json:"failure_threshold" binding:"required"`
	}
	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}
	url := URL{Url: requestBody.Url, CrawlTimeout: requestBody.CrawlTimeout, Frequency:requestBody.Frequency, Status:"active", FailureThreshold:requestBody.FailureThreshold}
	db.NewRecord(url)
	db.Create(&url)
	c.JSON(200, url.Serialize())
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var url URL
	if err := db.Where("id = ?", id).First(&url).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, url.Serialize())
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var urls []URL

	if err := db.Find(&urls).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		length := len(urls)
		serialized := make([]JSON, length, length)
		for i := 0; i < length; i++ {
			serialized[i] = urls[i].Serialize()
		}

		c.JSON(200, serialized)
	}
}


func updateStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var url URL
	id := c.Params.ByName("id")
	status := c.Params.ByName("status")

	if err := db.Where("id = ?", id).First(&url).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	if status == "activate" || status == "deactivate"{
		url.Status=status
	}
	c.BindJSON(&url)
	db.Save(&url)
	c.JSON(200, url.Serialize())

}


