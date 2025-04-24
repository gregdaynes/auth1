package main

import (
	"context"
)

// GreetingOutput represents the greeting operation response.
type HomeOutput struct {
	Body []byte
}

type HomeInput struct {
}

func (app *application) home(ctx context.Context, input *HomeInput) (*HomeOutput, error) {
	app.logger.Info("home")
	resp := &HomeOutput{}
	resp.Body = []byte("Hello world!")

	return resp, nil
}
