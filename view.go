package main

func indexRoute(params map[string][]string) string {
	output := blog.PageHeader

	posts, err := database.GetAllPosts()
	if err != nil {
		output += fmt.Sprintf("Error: %s", err)
	} else {
		for _, post := range posts {
			output += blog.GetPostListHtml(post)
		}
	}

	output += blog.PageFooter

	return output
}

func viewRoute(params map[string][]string) string {
	output := blog.PageHeader
	id, err := getPostId(params)

	if err == nil {
		post, err := database.GetPost(id)
		if err != nil {
			output += fmt.Sprintf("Error: %s", err)
		} else {
			output += blog.GetPostHtml(post)
		}
	} else {
		output += fmt.Sprintf("Error: %s", err)
	}

	output += blog.PageFooter

	return output
}
