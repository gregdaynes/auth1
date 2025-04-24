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
	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	fmt.Println(app.config)
	return resp, nil
}
