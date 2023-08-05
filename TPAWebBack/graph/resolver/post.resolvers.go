package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yahkerobertkertasnya/TPAWebBack/graph/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, newPost model.NewPost) (*model.Post, error) {
	var user *model.User
	userID := ctx.Value("UserID").(string)

	if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	boolVar := false
	post := &model.Post{
		ID:           uuid.NewString(),
		UserID:       userID,
		User:         user,
		Content:      newPost.Content,
		Privacy:      newPost.Privacy,
		LikeCount:    0,
		CommentCount: 0,
		ShareCount:   0,
		Files:        newPost.Files,
		Liked:        &boolVar,
		CreatedAt:    time.Now(),
	}

	if err := r.DB.Save(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, newComment model.NewComment) (*model.Comment, error) {
	var user *model.User
	userID := ctx.Value("UserID").(string)

	if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	boolVar := false
	comment := &model.Comment{
		ID:              uuid.NewString(),
		UserID:          userID,
		User:            user,
		Content:         newComment.Content,
		Liked:           &boolVar,
		LikeCount:       0,
		ReplyCount:      0,
		ParentPostID:    newComment.ParentPost,
		ParentCommentID: newComment.ParentComment,
		CreatedAt:       time.Now(),
	}

	if err := r.DB.Save(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// LikePost is the resolver for the likePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	var postLike *model.PostLike
	userID := ctx.Value("UserID").(string)

	if err := r.DB.First(&postLike, "post_id = ? AND user_id = ?", postID, userID).Error; err != nil {
		postLike = &model.PostLike{
			PostID: postID,
			UserID: userID,
		}
		if err := r.DB.Save(&postLike).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.DB.Delete(&postLike).Error; err != nil {
			return nil, err
		}
	}

	return postLike, nil
}

// Likecomment is the resolver for the likecomment field.
func (r *mutationResolver) Likecomment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	var commentLike *model.CommentLike
	userID := ctx.Value("UserID").(string)

	if err := r.DB.First(&commentLike, "post_id = ? AND user_id = ?", commentID, userID).Error; err != nil {
		commentLike = &model.CommentLike{
			CommentID: commentID,
			UserID:    userID,
		}
		if err := r.DB.Save(&commentLike).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.DB.Delete(&commentLike).Error; err != nil {
			return nil, err
		}
	}

	return commentLike, nil
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: GetPost - getPost"))
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, pagination model.Pagination) ([]*model.Post, error) {
	var posts []*model.Post

	if err := r.DB.
		Order("created_at desc").
		Preload("User").
		Preload("Likes").
		Preload("Comments").
		Offset(pagination.Start).
		Limit(pagination.Limit).
		Find(&posts).Error; err != nil {
		return nil, err
	}

	userID := ctx.Value("UserID").(string)

	for _, post := range posts {
		temp := false

		post.CommentCount = int(r.DB.Model(post).Association("Comments").Count())
		post.LikeCount = int(r.DB.Model(post).Association("Likes").Count())

		if err := r.DB.First(&model.PostLike{}, "post_id = ? AND user_id = ?", post.ID, userID).Error; err == nil {
			temp = true
		}

		//fmt.Println(temp)
		post.Liked = &temp
	}

	return posts, nil
}

// GetCommentPost is the resolver for the getCommentPost field.
func (r *queryResolver) GetCommentPost(ctx context.Context, postID string) ([]*model.Comment, error) {
	var comments []*model.Comment

	if err := r.DB.
		Preload("User").
		Preload("Likes").
		Preload("Comments").
		Preload("Comments.User").
		Find(&comments, "parent_post_id = ?", postID).Error; err != nil {
		return nil, err
	}

	userID := ctx.Value("UserID").(string)

	for _, comment := range comments {
		temp := false

		comment.LikeCount = int(r.DB.Model(comment).Association("Likes").Count())

		for _, replies := range comment.Comments {
			temp := false

			replies.LikeCount = int(r.DB.Model(replies).Association("Likes").Count())

			if err := r.DB.First(&model.CommentLike{}, "comment_id = ? AND user_id = ?", replies.ID, userID).Error; err == nil {
				temp = true
			}

			replies.Liked = &temp
		}

		if err := r.DB.First(&model.CommentLike{}, "comment_id = ? AND user_id = ?", comment.ID, userID).Error; err == nil {
			temp = true
		}

		comment.Liked = &temp
	}

	return comments, nil
}
