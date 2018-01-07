/*
Populates the database with a few posts so we have something to look at
*/

package main

import (
	"fmt"
	"github.com/dannyroes/golangkwblog/blog"
	"github.com/dannyroes/golangkwblog/database"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	//Clear the database
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	database.DB.C("posts").RemoveAll(bson.M{})

	//Create three posts and insert them
	post := blog.Post{}
	post.Author = "Danny"
	post.Title = "Hello World!"
	post.Body = "This is my first post on my new blog!"
	post.PostedTime = time.Now().Add(-5 * time.Hour)
	err = database.InsertPost(post)

	if err != nil {
		panic(err)
	}

	post.Author = "Danny"
	post.Title = "Is this working?"
	post.Body = "I can't seem to view my posts, something must be broken."
	post.PostedTime = time.Now().Add(-4 * time.Hour)
	err = database.InsertPost(post)

	if err != nil {
		panic(err)
	}

	post.Author = "Danny"
	post.Title = "Still coding"
	post.Body = "Still trying to fix my blog, reading lots of docs at <a href='https://golang.org/doc'>golang.org</a>"
	post.PostedTime = time.Now().Add(-5 * time.Minute)
	err = database.InsertPost(post)

	if err != nil {
		panic(err)
	}

	fmt.Println("Created test blog posts!")
}
