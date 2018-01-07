package blog

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Id          bson.ObjectId `bson:"_id"`
	Author      string        `bson:"author"`
	Title       string        `bson:"title"`
	Body        string        `bson:"body"`
	PostedTime  time.Time     `bson:"created"`
	UpdatedTime time.Time     `bson:"updated,omitempty"`
	Comments    []Comment     `bson:"comments,omitempty"`
	Likes       int           `bson:"likes,omitempty"`
}

type Comment struct {
	Id         bson.ObjectId `bson:"_id"`
	Author     string        `bson:"author"`
	Text       string        `bson:"comment"`
	PostedTime time.Time     `bson:"created"`
}

func (post Post) GetListHtml() string {
	html := `<div class="postList"><a href="/viewPost?id=` + post.Id.Hex() + `">` + post.Title + `</a> <span class="byline">posted by ` +
		post.Author + ", " + post.PostedTime.Format("Mon Jan 2, 2006 15:04") + "</span></div>\n"
	return html
}

func (post Post) GetHtml() string {
	html := `<div class="post"><h2>` + post.Title + "</h2>" +
		`<span class="byline">Posted by ` + post.Author + " " +
		post.PostedTime.Format("Mon Jan 2, 2006 15:04") + ` (<a href="/edit?id=` + post.Id.Hex() + `">edit</a>)</span><div class="postBody">` +
		post.Body + "</div></div>"

	return html
}

func (post Post) GetLikesHtml() string {
	html := `<div class="likes">` + fmt.Sprintf("%d", post.Likes) + ` likes <a href="/like?id=` + post.Id.Hex() + `">+</a></div>`

	return html
}

func (post Post) GetCommentsHtml() string {
	html := `<div class="comments"><h4>Comments</h3>`
	if len(post.Comments) == 0 {
		html += "No comments yet!"
	}
	for _, comment := range post.Comments {
		html += `<div class="comment"><span class="commenter">` + comment.Author + `</span>
		 <span class="byline">` + comment.PostedTime.Format("Mon Jan 2, 2006 15:04") +
			` (<a href="/deleteComment?comment_id=` + comment.Id.Hex() + `&post_id=` + post.Id.Hex() + `">delete</a>)</span>
		<br>` +
			comment.Text + "</div>"
	}
	html += `</div>
		<h5>Add a comment</h4>
		<form method="POST" action="/addComment">
		<input type="hidden" name="post_id" value="` + post.Id.Hex() + `">
		Name: <input type="text" name="author"><br>
		Comment:<br>
		<textarea name="comment" rows=3 cols=40></textarea><br><br>
		<input type="submit" value="Add Comment">
		</form>`

	return html
}
