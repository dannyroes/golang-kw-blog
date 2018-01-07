/*
Define the routes we support in our web app.
Any action the user has on the blog should be defined here.
*/

package blog

import (
	"fmt"
)

//Helper function to simplfy getting a string from a map[string][]string
func getStringParam(paramName string, params map[string][]string) string {
	var param string

	raw, ok := params[paramName]
	if ok && len(raw) == 1 {
		param = raw[0]
	}

	return param
}

func addRoutes() {
	//TODO: Add more routes!
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
		//TODO: Implement likes and comments
	}

	output += PageFooter

	return WebResponse{Body: output}
}

func createRoute(params map[string][]string) WebResponse {
	output := PageHeader + GetCreateFormHtml() + PageFooter
	return WebResponse{Body: output}
}

func editRoute(params map[string][]string) WebResponse {
	//TODO: Load the specified post and display its edit form
	return WebResponse{}
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
		post.Author = ""
		//TODO: Set the post's fields
		//TODO: Insert or update based on the action param
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
	//TODO: Delete the specified post and redirect to the index
	return WebResponse{}
}

func saveCommentRoute(params map[string][]string) WebResponse {
	id := getStringParam("id", params)
	//TODO: Create a comment struct and save it

	redirectUrl := fmt.Sprintf("/viewPost?id=%s", id)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}

func deleteCommentRoute(params map[string][]string) WebResponse {
	//TODO: Delete the specified comment and redirect to the post

	return WebResponse{}
}

func likeRoute(params map[string][]string) WebResponse {
	id := getStringParam("id", params)
	AddLike(id)
	redirectUrl := fmt.Sprintf("/viewPost?id=%s", id)

	return WebResponse{RedirectUrl: redirectUrl, RedirectCode: 303}
}
