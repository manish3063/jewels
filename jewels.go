package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(r *gin.Engine) {

	r.GET("/jewels/:jewels/stone/:stones", route2)

}

//Dummy function
func route2(c *gin.Context) {

	//date_str, ok := c.Params.Get("from_date")

	//total_cases := getTotalCases(records, date)
	var jewels, err1 = c.Params.Get("jewels")
	var stones, err2 = c.Params.Get("stones")
	count := getStones(jewels, stones)
	if err1 == false {
		res := gin.H{
			"error": "jewels is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	if err2 == false {
		res := gin.H{
			"error": "stones is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"jewels": jewels,
		"stones": stones,
		"count":  count,
	}
	c.JSON(http.StatusOK, res)
}

func getStones(jewels, stones string) int {

	var count int
	jewels_split := strings.Split(jewels, "")
	stones_split := strings.Split(stones, "")

	for i := 0; i < len(jewels_split); i++ {

		for j := 0; j < len(stones_split); j++ {
			if jewels_split[i] == stones_split[j] {
				count = count + 1
			}
		}

	}

	return count

}
