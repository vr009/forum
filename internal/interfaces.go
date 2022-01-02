package internal

import (
	"forum/internal/models"
	"github.com/jackc/pgtype"
)

type Repo interface {
	CreateForum(forum models.Forum) (models.Forum, *models.InternalError)
	GetForum(forum models.Forum) (models.Forum, *models.InternalError)
	CreateThread(forum models.Forum, thread models.Thread) (models.Thread, *models.InternalError)
	GetUsers(forum models.Forum, limit int32, sinceUser models.User, desc bool) (models.Users, *models.InternalError)
	GetThreads(forum models.Forum, limit int32, since pgtype.Timestamptz, desc bool) (models.Threads, *models.InternalError)
	GetThreadInfo(thread models.Thread, related interface{}) (models.Thread, models.Error)
	UpdateMessage(post models.Post, update models.PostUpdate) (models.Post, models.Error)
	DropAllData() models.Error
	GetStatus() (models.Status, models.Error)
	CreatePosts(thread models.Thread, posts models.Posts) (models.Posts, models.Error)
	GetThreadInfoBySlug(thread models.Thread) (models.Thread, models.Error)
	UpdateThread(thread models.Thread, update models.ThreadUpdate) (models.Thread, models.Error)
	GetPosts(forum models.Forum, limit int32, sinceUser models.User, desc bool) (models.Posts, models.Error)
	VoteForThread(thread models.Thread, vote models.Vote) (models.Thread, models.Error)
	CreateUser(user models.User) (models.User, models.Error)
	GetUser(user models.User) (models.User, models.Error)
	UpdateUser(user models.User) (models.User, models.Error)
}

type Usecase interface {
	CreateForum(forum models.Forum) (models.Forum, *models.InternalError)
	GetForum(forum models.Forum) (models.Forum, *models.InternalError)
	CreateThread(forum models.Forum, thread models.Thread) (models.Thread, *models.InternalError)
	GetUsers(forum models.Forum, limit int32, sinceUser models.User, desc bool) (models.Users, *models.InternalError)
	GetThreads(forum models.Forum, limit int32, since pgtype.Timestamptz, desc bool) (models.Threads, *models.InternalError)
	GetThreadInfo(thread models.Thread, related interface{}) (models.Thread, models.Error)
	UpdateMessage(post models.Post, update models.PostUpdate) (models.Post, models.Error)
	DropAllInfo() models.Error
	GetStatus() (models.Status, models.Error)
	CreatePosts(thread models.Thread, posts models.Posts) (models.Posts, models.Error)
	GetThreadInfoBySlug(thread models.Thread) (models.Thread, models.Error)
	UpdateThread(thread models.Thread, update models.ThreadUpdate) (models.Thread, models.Error)
	GetPosts(forum models.Forum, limit int32, sinceUser models.User, desc bool) (models.Posts, models.Error)
	VoteForThread(thread models.Thread, vote models.Vote) (models.Thread, models.Error)
	CreateProfile(user models.User) (models.User, models.Error)
	GetProfile(user models.User) (models.User, models.Error)
	UpdateProfile(user models.User) (models.User, models.Error)
}
