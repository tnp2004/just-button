package main

import (
	"clientserver/services"
	"fmt"
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RequestData struct {
	N int64 `json:"n"`
}

type ResponseResult struct {
	N int64 `json:"n"`
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calcClient := services.NewCalcClient(cc)
	calcService := services.NewCalcService(calcClient)

	app := fiber.New()
	app.Use(cors.New())
	app.Patch("/update", func(c *fiber.Ctx) error {

		var data RequestData
		err := c.BodyParser(&data)
		if err != nil {
			return c.SendString("Something went wrong!")
		}

		var result int
		if data.N >= 0 {
			err = calcService.Plus(uint64(data.N), &result)
		} else if data.N < 0 {
			err = calcService.Minus(int64(math.Abs(float64(data.N))), &result)
		}

		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(ResponseResult{
			N: int64(result),
		})
	})

	app.Listen(":3000")

}
