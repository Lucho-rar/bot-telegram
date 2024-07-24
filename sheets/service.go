package sheets

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewService(apiKey string) (*sheets.Service, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return srv, nil
}
