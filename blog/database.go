/*
Generic functions for database access
*/

package blog

import (
	"gopkg.in/mgo.v2"
)

const (
	MongoURI = "localhost"
	DBName   = "golang-kw-blog"
)

var DB *mgo.Database
var Session *mgo.Session

//Connect to MongoDB using the hardcoded credentials
func ConnectDB() error {
	var err error
	Session, err = mgo.Dial(MongoURI)
	if err != nil {
		return err
	}

	DB = Session.DB(DBName)
	return nil
}

//Confirm we're currently connected to the database
func isDbConnected() bool {
	return Session != nil && DB != nil
}
