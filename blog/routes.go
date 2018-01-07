/*
Define the routes we support in our web app.
Any action the user has on the blog should be defined here.
*/

package blog

import (
	"fmt"
)

func addRoutes() {
	CreateRoute("/viewPost", viewRoute)
	CreateRoute("/create", createRoute)
	CreateRoute("/edit", editRoute)
	CreateRoute("/save", saveRoute)
	CreateRoute("/delete", deleteRoute)
	CreateRoute("/like", likeRoute)
	CreateRoute("/addComment", saveCommentRoute)
	CreateRoute("/deleteComment", deleteCommentRoute)
	CreateRoute("/", indexRoute)
}

func indexRoute(params map[string][]string) WebResponse {
	output := PageHeader

	posts, err := GetAllPosts()
	if err != nil {
		output += fmt.Sprintf("Error: %s", err)
		output += PageFooter
		return WebResponse{Body: output}
	}

	for _, post := range posts {
		output += post.GetListHtml()
	}

	output += PageFooter

	return WebResponse{Body: output}
}

func viewRoute(params map[string][]string) WebResponse {
	output := PageHeader
	id := getStringParam("id", params)

	post, err := GetPost(id)
	if err != nil {
		output += fmt.Sprintf("Error: %s", err)
	} else {
		output += post.GetHtml()
		output += post.GetLikesHtml()
		output += post.GetCommentsHtml()
	}

	output += PageFooter

	return WebResponse{Body: output}
}

func createRoute(params map[string][]string) WebResponse {
	output := PageHeader + GetCreateFormHtml() + PageFooter
	return WebResponse{Body: output}
}

func editRoute(params map[string][]string) WebResponse {
	output := PageHeader
	id := getStringParam("id", params)

	post, err := GetPost(id)
	if err != nil {
		output += fmt.Sprintf("Error: %s", err)
	} else {
		output += GetEditFormHtml(post)
	}

	output += PageFooter
	return WebResponse{Body: output}
}

func saveRoute(params map[string][]string) WebResponse {
	var err error
	var id string
	post := Post{}

	action := getStringParam("action", params)
	if action == "update" {
		id = getStringParam("id", params)
		post, err = GetPost(id)
	}

	if err == nil {
		post.Author = getStringParam("author", params)
		post.Title = getStringParam("title", params)
		post.Body = getStringParam("body", params)

		if action == "create" {
			id, err = InsertPost(post)
		} else {
			err = UpdatePost(post)
		}
	}

	if err != nil {
		output := PageHeader
		output += fmt.Sprintf("Could not %s post: %s", action, err)
		output += PageFooter
		return WebResponse{Body: output}
	}

	redirectUrl := fmt.Sprintf("/viewPost?id=%s", id)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}

func deleteRoute(params map[string][]string) WebResponse {
	id := getStringParam("id", params)
	DeletePost(id)
	fmt.Println(id)

	return WebResponse{RedirectUrl: "/", RedirectCode: 303}
}

func saveCommentRoute(params map[string][]string) WebResponse {
	comment := Comment{}
	id := getStringParam("post_id", params)

	comment.Author = getStringParam("author", params)
	comment.Text = getStringParam("comment", params)

	AddComment(id, comment)

	redirectUrl := fmt.Sprintf("/viewPost?id=%s", id)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}

func deleteCommentRoute(params map[string][]string) WebResponse {
	postId := getStringParam("post_id", params)
	commentId := getStringParam("comment_id", params)

	DeleteComment(postId, commentId)

	redirectUrl := fmt.Sprintf("/viewPost?id=%s", postId)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}

func likeRoute(params map[string][]string) WebResponse {
	id := getStringParam("id", params)
	AddLike(id)
	redirectUrl := fmt.Sprintf("/viewPost?id=%s", id)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}

func getStringParam(paramName string, params map[string][]string) string {
	var param string

	raw, ok := params[paramName]
	if ok && len(raw) == 1 {
		param = raw[0]
	}

	return param
}
