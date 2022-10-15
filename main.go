package main

import ( "net/http"

"github.com/gin-gonic/gin"

)

//car // { // "id":"1" // "brand":"honda" // "car_type":"" // }

type car struct { ID string json:"id" Brand string json:"brand" Type string json:"car_type" }

var cars = []car{ {ID: "1", Brand: "Honda", Type: "City"}, {ID: "2", Brand: "Toyota", Type: "Avanza"}, }

func main() { r := gin.New() r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{ "status": "ok", })

})
//GET /car - list cars
r.GET("/cars", func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cars)

})
//POST / cars - create cars
r.POST("/cars", func(ctx *gin.Context) {
	var car car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	cars = append(cars, car)
	ctx.JSON(http.StatusCreated, car)

})
//PUT /cars - update car
//delete / car - delete car
r.DELETE("/cars/:car_id", func(ctx *gin.Context) {
	id := ctx.Param("car_id")
	for i, car := range cars {
		if car.ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			break
		}
	}
	// 0 1 2 3
	// cars[:2]
	// 0 1

	ctx.Status(http.StatusNoContent)
})
r.Run()

}