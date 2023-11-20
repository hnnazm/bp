package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Counter struct {
	ID              uuid.UUID `json:"id"`
	AdjustmentValue int       `json:"adjustmentValue"`
	CurrentValue    int       `json:"currentValue"`
	CreatedAt       time.Time `json:"createdAt"`
}

type UpdateCounterRequest struct {
	Value int `json:"value"`
}

type UpdateCounterResponse struct {
	Data  *Counter `json:"data"`
	Error string   `json:"error,omitempty"`
}

func (app *Application) GetCounter(ctx context.Context) (Counter, error) {
	var counter Counter

	query := `
  SELECT id, adjustment_value, current_value, created_at
  FROM counter_view
  ORDER BY id DESC
  LIMIT 1;
  `

	err := app.DB.QueryRow(query).Scan(&counter.ID, &counter.AdjustmentValue, &counter.CurrentValue, &counter.CreatedAt)

	if err != nil {
		log.Fatal(err)
		return Counter{}, err
	}

	return counter, nil
}

func (app *Application) UpdateCounterHandler(c echo.Context) error {
	input := new(UpdateCounterRequest)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, UpdateCounterResponse{
			Data:  nil,
			Error: err.Error(),
		})
	}

	query := `
	 INSERT INTO counter (adjustment_value)
	 VALUES (%d);
  `

	_, err := app.DB.Exec(fmt.Sprintf(query, input.Value))

	if err != nil {
		return c.JSON(http.StatusBadRequest, UpdateCounterResponse{
			Data:  nil,
			Error: err.Error(),
		})
	}

	counter, err := app.GetCounter(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusBadRequest, UpdateCounterResponse{
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, UpdateCounterResponse{
		Data: &counter,
	})
}
