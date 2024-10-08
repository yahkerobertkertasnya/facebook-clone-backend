package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/graph/model"
)

// CreateTextStory is the resolver for the createTextStory field.
func (r *mutationResolver) CreateTextStory(ctx context.Context, input model.NewTextStory) (*model.Story, error) {
	userID := ctx.Value("UserID").(string)
	return r.StoryService.CreateTextStory(userID, input)
}

// CreateImageStory is the resolver for the createImageStory field.
func (r *mutationResolver) CreateImageStory(ctx context.Context, input model.NewImageStory) (*model.Story, error) {
	userID := ctx.Value("UserID").(string)
	return r.StoryService.CreateImageStory(userID, input)
}

// GetStories is the resolver for the getStories field.
func (r *queryResolver) GetStories(ctx context.Context, username string) ([]*model.Story, error) {
	return r.StoryService.GetStories(username)
}

// GetUserWithStories is the resolver for the getUserWithStories field.
func (r *queryResolver) GetUserWithStories(ctx context.Context) ([]*model.User, error) {
	userID := ctx.Value("UserID").(string)
	return r.StoryService.GetUserWithStories(userID)
}

// User is the resolver for the user field.
func (r *storyResolver) User(ctx context.Context, obj *model.Story) (*model.User, error) {
	return r.StoryService.User(obj)
}

// Story returns graph.StoryResolver implementation.
func (r *Resolver) Story() graph.StoryResolver { return &storyResolver{r} }

type storyResolver struct{ *Resolver }
