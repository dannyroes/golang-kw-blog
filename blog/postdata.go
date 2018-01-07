/*
Database functions for Posts
Handles CRUD operations for Posts, Likes and Comments
*/
package blog

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func InsertPost(post Post) (string, error) {
	if !isDbConnected() {
		return "", errors.New("Database is not connected")
	}
	collection := DB.C("posts")

	if post.Id != "" {
		return "", errors.New("Tried to insert a post that already has an ID")
	}

	post.Id = bson.NewObjectId()

	//TODO: If the post doesn't have a created time set explicitly, set one

	return post.Id.Hex(), collection.Insert(post)
}

func GetAllPosts() ([]Post, error) {
	var posts []Post

	if !isDbConnected() {
		return posts, errors.New("Database is not connected")
	}

	err := DB.C("posts").Find(nil).Sort("-created").All(&posts)

	return posts, err
}

func GetPost(id string) (Post, error) {
	var post Post

	if !isDbConnected() {
		return post, errors.New("Database is not connected")
	}

	//TODO: Query for the specified post by id
	query := bson.M{}

	err := DB.C("posts").Find(query).One(&post)
	return post, err
}

func UpdatePost(post Post) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}
	collection := DB.C("posts")
	//TODO: Set the post's UpdatedTime and save to the database
	query := bson.M{}

	return collection.Update(query, post)
}

func DeletePost(id string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}
	collection := DB.C("posts")

	return collection.RemoveId(bson.ObjectIdHex(id))
}

func AddLike(id string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	//TODO: Make an update object to increment the likes counter
	update := bson.M{}

	err := DB.C("posts").Update(query, update)
	return err
}

func AddComment(id string, comment Comment) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}

	if comment.Id != "" {
		return errors.New("Tried to insert a comment that already has an ID")
	}

	comment.Id = bson.NewObjectId()
	comment.PostedTime = time.Now()

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	//TODO: Make an update object to push the new comment to the comments array
	update := bson.M{}

	err := DB.C("posts").Update(query, update)
	return err
}

func DeleteComment(postId, commentId string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}

	query := bson.M{"_id": bson.ObjectIdHex(postId)}
	//TODO: Make an update object to push remove the specified comment from the comments array
	update := bson.M{}

	err := DB.C("posts").Update(query, update)
	return err
}
