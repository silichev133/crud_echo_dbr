package main

import (
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Users struct {
	Users []User `json:"users"`
}

var (
	usertable = "users"
	seq       = 1
	conn, err = dbr.Open("postgres", "postgres:@tcp(localhost:5432)/test", nil)
	sess      = conn.NewSession(nil)
)

func main() {
	// Opening db
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", showAllEmployees)
	//e.GET("/user/:id", showEmployee)
	//e.PUT("/user/:id", updateEmployees)
	e.POST("/user", newEmployees)
	//e.DELETE("/user/:id", deleteEmployee)

	e.Logger.Fatal(e.Start(":9090"))
}

func showAllEmployees(c echo.Context) error {
	var users Users
	sess.Select("*").From(usertable).Load(&users)
	return c.JSON(http.StatusOK, users)
}

// func newEmployees(c echo.Context) error {
// 	user := new(User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}
// 	sess.InsertInto(usertable).Columns("id", "name").Values(user.Id, user.Name).Exec()
// 	return c.JSON(http.StatusCreated, user)
// }

func newEmployees(c echo.Context) error {
	user := &User{
		Id: seq,
	}
	if err := c.Bind(user); err != nil {
		return err
	}
	//Users[u.Id] = u
	sess.InsertInto(usertable).Columns("name").Values(user.Name).Exec()
	seq++
	return c.JSON(http.StatusCreated, user)
}
