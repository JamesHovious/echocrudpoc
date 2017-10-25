# EchoCRUDPOC 

is a *proof of concept* web application that uses an organized group of structs instead of a database.

## Goal
* Allow for CRUD like behavior/queries
* Allow for relational data to be stored/queried similar to a relational database
* Allow the data to persist & restore should the web application crash or close

Below is a set of data that one might track in a web application. 

Notice that the `Schema` struct is just a slice of pointers to the data that we want to manage. 
Each struct (`User`, `Project`) is analogous to a database table. 
Each struct field is analogous to a table column. 
Each slice of data that is part of `Schema` would resemble a row of data in a database. 
Notice that `User` contains a field  that is a slice of pointers to `Project`. 
This let's us define relatioships similar to a relational database. 
If you want to update `Project` you create a new struct a change its pointer. `Users`'s reference will automatically update. 

If you close the application the whole data structure is encoded as a [gob](https://blog.golang.org/gobs-of-data) and saved to disk. The benefit of this is that when encoding occurs all pointers are flattened, so the underlying data is saved. When the structure is loaded up again as the applicaiton restarts the pointer relationships are restored.

```
type Schema struct {
	User    []*User
	Project []*Project
}

type User struct {
	Email    string
	Password string
	Username string
	Projects []*Project // Just specify a slice a pointers
                      // to define a relationship with another struct/"table"
}

type Project struct {
	ProjectCode string
	ProjectName string
}
```


Here are some `curl` commands to test. 
Run the applications and test the following endpoints
```
// CREATE 
curl -X POST \
-H 'Content-Type: application/json' \
-d '{"Email":"bill@aol.com","Username":"h0b0","Password":"pass"}' \
localhost:1323/users

// READ
curl \
  -X GET \
  http://localhost:1323/users/h0b0

// UPDATE
curl -X PUT \
-H 'Content-Type: application/json' \
-d '{"Email": "ballin@netscap.net","Username":"h0b0","Password":"NEWSECUREPASSWORD"}' \
localhost:1323/users/h0b0

// DELETE
curl -X DELETE \
-H 'Content-Type: application/json' \
-d '{"Username": "h0b0"}' \
localhost:1323/users/h0b0

// Dump the whole data structure to view its contents
curl \
  -X GET \
  http://localhost:1323/database
```
