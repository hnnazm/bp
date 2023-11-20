package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Application struct {
	DB *sql.DB
}

func initDB() (*sql.DB, error) {
	var (
		DATABASE_HOST      = os.Getenv("DATABASE_HOST")
		DATABASE_USER      = os.Getenv("DATABASE_USER")
		DATABASE_PASSSWORD = os.Getenv("DATABASE_PASSWORD")
		DATABASE_NAME      = os.Getenv("DATABASE_NAME")
	)

	if DATABASE_HOST == "" && DATABASE_USER == "" && DATABASE_PASSSWORD == "" && DATABASE_NAME == "" {
		log.Fatal("DATABASE environment variable is required")
	}

	DATABASE_DSN := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		DATABASE_USER,
		DATABASE_PASSSWORD,
		DATABASE_HOST,
		DATABASE_NAME,
	)

	db, err := sql.Open("postgres", DATABASE_DSN)

	if err != nil {
		return nil, err
	}

	m, err := migrate.New(
		"file://migrations",
		DATABASE_DSN,
	)

	if err != nil {
		return nil, err
	}

	m.Up()

	return db, nil
}

func main() {
	db, err := initDB()

	if err != nil {
		log.Fatal(err)
	}

	app := Application{
		DB: db,
	}

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", app.indexHandler)
	e.POST("/console/global-settings", app.updateSiteHandler)
	e.POST("/counter", app.UpdateCounterHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func (app *Application) indexHandler(c echo.Context) error {
	var applicationName string

	err := app.DB.QueryRow("SELECT application_name FROM global_settings").Scan(&applicationName)

	if err != nil {
		log.Fatal(err)
	}

	counter, err := app.GetCounter(c.Request().Context())

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"applicationName": applicationName,
		"counter":         counter.CurrentValue,
	})
}

func (app *Application) updateSiteHandler(c echo.Context) error {
	applicationName := c.FormValue("application_name")

	query := `
  INSERT INTO global_settings (id, application_name)
  VALUES (1, %s)
  ON CONFLICT (id) DO UPDATE
  SET application_name = EXCLUDED.application_name;
  `

	_, err := app.DB.Exec(fmt.Sprintf(query, applicationName))

	if err != nil {
		log.Fatal(err)
	}

	return c.NoContent(http.StatusOK)
}
