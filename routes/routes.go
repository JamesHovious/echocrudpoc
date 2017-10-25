package routes

import (
	"net/http"

	"github.com/JamesHovious/echocrudpoc/database"
	"github.com/JamesHovious/echocrudpoc/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

// CreateUser creates a 'record' for a new user.
// Sample input: {"Email":"bill@aol.com","Username":"h0b0","Password":"pass"}
func CreateUser(c echo.Context) error {
	// Parse the incoming JSON
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusCreated, err)
	}
	// Append the parsed json string to our []struct
	database.GobDB.User = append(database.GobDB.User, u)

	return c.String(http.StatusOK, "{\"message\":\"success\"}")
}

// GetUser retreives a record for a user
// Sample input: ?user=h0b0
func GetUser(c echo.Context) error {
	var u models.User
	u.Username = c.Param("username")
	ret := database.QueryUser(u.Username, database.GobDB)
	return c.String(http.StatusOK, ret)
}

// UpdateUser updates the record of a user
// Sample input: {"Email": "ballin@netscap.net","Username":"h0b0","Password":"NEWSECUREPASSWORD"}
func UpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusCreated, err)
	}
	userName := c.Param("username")
	if u.Username != userName {
		return c.String(http.StatusOK, "{\"message\":\"wrong username\"}")
	}
	ret := database.UpdateUser(*u, database.GobDB)
	return c.String(http.StatusOK, ret)
}

// DeleteUser deletes a record for a user.
// Sample input: {"Username":"h0b0"}
func DeleteUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusCreated, err)
	}
	userName := c.Param("username")
	if u.Username != userName {
		return c.String(http.StatusOK, "{\"message\":\"wrong username\"}")
	}
	ret := database.DeleteUser(*u, database.GobDB)
	return c.String(http.StatusOK, ret)
}

// Return the contents of the database. For debugging and POC purposes.
func ShowDatabase(c echo.Context) error {
	return c.JSON(http.StatusOK, spew.Sdump(database.GobDB))
}
