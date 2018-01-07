/*
Constants and functions for generating HTML responses
*/

package blog

const (
	PageHeader = `
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<title>Golang KW Blog</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.3/css/bootstrap.min.css">
		<style type="text/css">
			div.footer {
				margin-top: 2em;
				padding: 1em;
				background-color: #CCCCCC;
			}
			span.byline {
				font-size: 0.8em;
				font-style: italic;
			}
			div.postBody {
				margin-top: 1em;
			}
			div.comments {
				margin-top: 1em;
				margin-bottom: 1em;
			}
			span.commenter {
				font-weight: bold;
			}
			div.comment {
				margin-bottom: 0.5em;
			}
		</style>
	</head>
	<body>
	<div class="container">
		<div class="content">
			<h1><a href="/">Golang KW Blog</a></h1>
`
	PageFooter = `
		</div>
		<div class="footer">
			<a href="/create">New Post</a> | 
			<a href="https://www.meetup.com/Golang-KW/" target="_blank">Golang KW</a> | 
			<a href="https://golang.org" target="_blank">Golang</a> | 
			<a href="https://dejero.com" target="_blank">Dejero</a>
		</div>
	</div>
`
)

func GetCreateFormHtml() string {
	html := `<div class="create"><h2>Create Post</h2>` +
		`<form method="POST" action="/save">
		<input type="hidden" name="action" value="create">
		Title: <input type="text" name="title"><br>
		Author: <input type="text" name="author"><br>
		Text:<br>
		<textarea name="body" rows=5 cols=60></textarea><br><br>
		<input type="submit" value="Post">
		</form>`
	return html
}

func GetEditFormHtml(post Post) string {
	html := `<div class="create"><h2>Edit Post</h2>` +
		`<form method="POST" action="/save">
		<input type="hidden" name="action" value="update">
		<input type="hidden" name="id" value="` + post.Id.Hex() + `">
		Title: <input type="text" name="title" value="` + post.Title + `"><br>
		Author: <input type="text" name="author" value="` + post.Author + `"><br>
		Text:<br>
		<textarea name="body" rows=5 cols=60>` + post.Body + `</textarea><br><br>
		<input type="submit" value="Update">
		</form>
		<form method="POST" action="/delete">
		<input type="hidden" name="id" value="` + post.Id.Hex() + `">
		<input type="submit" value="Delete">`
	return html
}
