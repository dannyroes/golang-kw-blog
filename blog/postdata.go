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

	//If the post doesn't have a created time set explicitly, set one
	if post.PostedTime.IsZero() {
		post.PostedTime = time.Now()
	}

	return post.Id.Hex(), collection.Insert(post)
}

func UpdatePost(post Post) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}
	post.UpdatedTime = time.Now()

	collection := DB.C("posts")
	query := bson.M{"_id": post.Id}

	return collection.Update(query, post)
}

func DeletePost(id string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}
	collection := DB.C("posts")

	return collection.RemoveId(bson.ObjectIdHex(id))
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

	query := bson.M{"_id": bson.ObjectIdHex(id)}

	err := DB.C("posts").Find(query).One(&post)
	return post, err
}

func AddLike(id string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}

	query := bson.M{"_id": bson.ObjectIdHex(id)}
	update := bson.M{"$inc": bson.M{"likes": 1}}

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
	update := bson.M{"$push": bson.M{"comments": comment}}

	err := DB.C("posts").Update(query, update)
	return err
}

func DeleteComment(postId, commentId string) error {
	if !isDbConnected() {
		return errors.New("Database is not connected")
	}

	query := bson.M{"_id": bson.ObjectIdHex(postId)}
	update := bson.M{"$pull": bson.M{"comments": bson.M{"_id": bson.ObjectIdHex(commentId)}}}

	err := DB.C("posts").Update(query, update)
	return err
}
