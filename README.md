# Basic Usage:

* Make sure MongoDB is installed, configured, and running. https://docs.mongodb.com/manual/installation/#mongodb-community-edition
* Populate the database by running `go run blog-setup/main.go`. Modify MongoURI in blog/database.go if you aren't connecting to localhost or need to add credentials
* Run the webserver with `go run main.go` or install a binary with `go install github.com/dannyroes/golangkwblog` and run `golangkwblog`
* Visit http://localhost:8080/ in a browser

# Ideas for further development
* Generate the page html using templates (https://golang.org/pkg/html/template/)
* Implement basic error pages for 404, 500 errors
* Change timestamp displays to be relative to current time
* Add the ability to attach tags to a post
* Add support to change an author/commenter's username across any posts/comments they've made
