package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"

	gqlmodels "github.com/dehwyy/mugen/apps/stream_broadcaster/internal/gql/models"
)

func (r *queryResolver) GetStream(ctx context.Context, req gqlmodels.StreamRequest) (*gqlmodels.StreamResponse, error) {
	r.Log.Info().Msgf("GetStream: ID = %s, Frame = %d", req.ID, req.Frame)
	return &gqlmodels.StreamResponse{
		Data: []byte("Hello World!"),
		ID:   req.ID,
	}, nil
}
