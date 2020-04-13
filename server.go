package main

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Users struct {
	Users []User `json:"users"`
}

var (
	usertable = "users"
	seq       = 1
	conn, err = dbr.Open("postgresql", "root:@tcp(localhost:5432)/test", nil)
	sess      = conn.NewSession(nil)
)

func main() {
	// Opening db
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/employees", showAllEmployees)
	//e.GET("/employee/:id", showEmployee)
	//e.PUT("/employee/:id", updateEmployees)
	e.POST("/employee", newEmployees)
	//e.DELETE("/employee/:id", deleteEmployee)

	e.Logger.Fatal(e.Start(":9090"))
}

func showAllEmployees(c echo.Context) error {
	//session begin
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var users []User
	// db.Find(&users)
	// return c.JSON(http.StatusOK, users)
}

func newEmployees(c echo.Context) error {
	//session begin
	// if err != nil {
	// 	panic(err.Error())
	// }

	// user := new(User)
	// if err := c.Bind(user); err != nil {
	// 	return err
	// }

	// db.Create(&user)
	// return c.String(http.StatusOK, "OK")
}
