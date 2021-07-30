package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alexsosic/fizz-buzz-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func updateStats(path string) {
	var stat models.Stats
	if err := models.DB.Where("request = ?", path).First(&stat).Error; err != nil {
		stat = models.Stats{Request: path, Hits: 1}
		models.DB.Create(&stat)
	} else {
		err := models.DB.Transaction(func(tx *gorm.DB) error {
			// Get if exist
			if err := tx.Where("request = ?", path).First(&stat).Error; err != nil {
				return err
			}
			// Increment Counter
			if err := tx.Model(&stat).Update("hits", stat.Hits+1).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}

// FizzBuzz GET /api/:int1/:int2/:limit/:str1/:str2
// FizzBuzz endpoint
func FizzBuzz(c *gin.Context) {
	var result string
	int1, err := strconv.Atoi(c.Param("int1"))
	if err != nil {
		log.Fatal(err)
	}
	int2, err := strconv.Atoi(c.Param("int2"))
	if err != nil {
		log.Fatal(err)
	}
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		log.Fatal(err)
	}
	str1 := c.Param("str1")
	str2 := c.Param("str2")

	for i := 1; i <= limit; i++ {

		if i%int1 == 0 && i%int2 == 0 {
			result = result + str1 + str2
		} else if i%3 == 0 {
			result = result + str1
		} else if i%5 == 0 {
			result = result + str2
		} else {
			result = result + strconv.Itoa(i)
		}

		if i < limit {
			result = result + ","
		}
	}
	updateStats(c.Request.URL.Path)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetStats GET /
// GetStats find all stats
func GetStats(c *gin.Context) {
	var stats []models.Stats
	models.DB.Find(&stats)

	c.JSON(http.StatusOK, gin.H{"data": stats})
}
