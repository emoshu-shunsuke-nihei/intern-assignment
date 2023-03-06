package main

import (
	"intern-assignment/handlers"

	_ "intern-assignment/docs"

	"intern-assignment/database"

	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
)

// func hello(c echo.Context) error {
// 	database.Connect()
// 	sqlDB, _ := database.DB.DB()
// 	defer sqlDB.Close()
// 	err := sqlDB.Ping()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "データベース接続失敗")
// 	} else {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	}
// }

// @title        intern-assignment
// @version      1.0
// @license.name intern-assignment
// @host         localhost:3000
// @BasePath     /
func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/swagger/*", echoswagger.WrapHandler)
	e.POST("/members", handlers.CreateMember)
	e.GET("/members", handlers.GetMemberList)
	e.GET("/members/:id", handlers.GetMemberDetailInfo)
	e.PUT("/members/:id", handlers.UpdateMember)
	e.DELETE("/members/:id", handlers.DeleteMember)
	e.Logger.Fatal(e.Start(":3000"))
}
