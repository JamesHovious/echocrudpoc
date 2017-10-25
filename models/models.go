package models

// The "schema" for the database is simply a struct that
// contains a slice of pointers to all of your structs.
// In this fashion your structs act as database tables.
// The struct feilds act as database columns
// And every item within the slice of pointser for the Schema
// struct is a row of data.
type Schema struct {
	User []*User
}

type User struct {
	Email    string
	Password string
	Username string
}

/*
If you wanted to add "foreign key" type relationships you could simply that like so:

type Schema struct {
	User    []*User
	Project []*Project // Remember to update your schema
}

type User struct {
	Email    string
	Password string
	Username string
	Projects []*Project // Just specify a slice a pointers
}


type Project struct {
	ProjectCode string
	ProjectName string
}

*/
