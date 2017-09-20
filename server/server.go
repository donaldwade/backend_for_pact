package main

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type MyResponse struct {
  OK bool `json:"ok"`
}

func main() {
  e := echo.New()
  e.Use(middleware.CORS())
  endPoint := MyResponse{
    OK: true,
  }
  jsonString, _ := json.Marshal(endPoint)
  fmt.Printf("%+v", endPoint)
  fmt.Println(string(jsonString))
  e.GET("/go/test", func(c echo.Context) error {
    return c.String(http.StatusOK, string(jsonString))
  })
  e.Logger.Fatal(e.Start(":1323"))
}
