package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/line", func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			legendData := []string{
				"A列車：開始吧 觀光開發計畫",
				"集合啦！動物森友會",
				"Fitness Boxing 2",
				"紙片瑪利歐：摺紙國王",
				"皮克敏™３ 豪華版",
				"ZELDA無雙 災厄啟示錄"}
			xAxisData := []string{
				"http://localhost:8081/70010000030778.jpg",
				"http://localhost:8081/animal_crossing.jpg",
				"http://localhost:8081/fitness_boxing_2.jpg",
				"http://localhost:8081/paper_mario.jpg",
				"http://localhost:8081/pikumi3.jpg",
				"http://localhost:8081/zelda_musou.jpg"}
			c.JSON(200, gin.H{
				"legend_data": legendData,
				"xAxis_data":  xAxisData,
			})
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000")

}
