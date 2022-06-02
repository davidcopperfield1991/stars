package server

import (
	"math/rand"
	"strconv"

	"github.com/davidcopperfield1991/stars/internal"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func Server() {
	r := gin.Default()
	r.GET("/chart/:num", chart)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":6060")
}

func chart(c *gin.Context) {
	num := c.Param("num")
	numint, _ := strconv.Atoi(num)
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by go-echarts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	bar.SetXAxis(internal.Render(numint)).
		AddSeries("Category A", internal.RenderInt(numint))

	bar.Render(c.Writer)
	c.JSON(200, gin.H{
		"message": num,
	},
	)

}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}