package main

import (
	"context"
	"fmt"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

func (app *application) greeting(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
	app.logger.Info("greeting", "name", input.Name)

	rows, err := app.db.QueryContext(ctx, "SELECT id FROM sample_table WHERE name = $1", input.Name)
	if err != nil {
		app.logger.Error("something bad happened")
	}

	var id int

	for rows.Next() {
		app.logger.Info("found", "name", input.Name)

		err := rows.Scan(&id)
		if err != nil {
			app.logger.Error("something bad happened", "err", err)
		}

		app.logger.Info("found id", "id", id)
	}

	if id == 0 {
		app.logger.Info("inserting", "name", input.Name)

		newEntries, err := app.db.QueryContext(ctx, "INSERT INTO sample_table (name) VALUES ($1) RETURNING id", input.Name)
		if err != nil {
			app.logger.Error("xxx something bad happened", "err", err)
		}

		for newEntries.Next() {
			err := newEntries.Scan(&id)

			if err != nil {
				app.logger.Error("something bad happened")
			}

			app.logger.Info("inserted id", "id", id)
		}
	}

	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s! (%d)", input.Name, id)
	return resp, nil
}
