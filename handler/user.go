package handler

import (
	"cloud-front-test/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// Signup godoc
// @Summary      Signup
// @Description  Register a a user
// @Accept       json
// @Produce      json
// @Param   	 user  body     model.UserSignup     yes  "user signup"
// @Success      200  {object}  model.User
// @Router       /signup [post]
func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind

	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	fmt.Println(u)
	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	affected, err := h.DB.Insert(u)
	fmt.Println("Created user with id", u.ID)
	//user.ID // returns inserted data's primary key
	fmt.Println(affected, err)

	return c.JSON(http.StatusCreated, u)
}

// Login godoc
// @Summary      Login
// @Description  Login a user
// @Accept       json
// @Produce      json
// @Param   	 user  body     model.UserSignup     yes  "user signup"
// @Success      200  {object}  model.User
// @Router       /login [post]
func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	if has, err := h.DB.Get(u); err != nil || !has {
		fmt.Println(err)
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	} else {
		fmt.Println(has)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

// BookmarksAdd godoc
// @Summary      Add Bookmark
// @Description  Add a bookmar to a user
// @Accept       json
// @Produce      json
// @Param   	 user  body     model.addBookmark     yes  "user signup"
// @Success      200  {object}  model.Bookmark
// @Router       /user/bookmarks [post]
func (h *Handler) BookmarksAdd(c echo.Context) (err error) {
	bookmark := new(model.Bookmark)
	bookmark.UserID = userIDFromToken(c)
	if err = c.Bind(bookmark); err != nil {
		return
	}

	if affected, err := h.DB.Insert(bookmark); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(affected)
	}
	fmt.Println("Created site with id", bookmark.ID)

	return c.JSON(http.StatusOK, bookmark)

}

// BookmarksDelete godoc
// @Summary      Delete Bookmark
// @Description  Delete a bookmark from a user
// @Produce      json
// @Param        id   path      int  true  "Bookmark ID"
// @Success      200
// @Router       /user/bookmarks/{id} [delete]
func (h *Handler) BookmarksDelete(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	bookmarkID, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {

	}
	affected, err := h.DB.Delete(&model.Bookmark{UserID: userID, ID: uint(bookmarkID)})
	if affected != 1 {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id or token"}
	}
	if err != nil {
		return err
	}
	fmt.Println("Deleted bookmark with id", bookmarkID)

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "OK"})

}

// BookmarksList godoc
// @Summary      ListBookmarks
// @Description  List Bookmarks from user
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Bookmark
// @Router       /user/bookmarks [get]
func (h *Handler) BookmarksList(c echo.Context) (err error) {
	var user = model.User{ID: userIDFromToken(c)}
	if has, err := h.DB.Get(&user); err != nil && !has {
		fmt.Println("There has been an error", err, has)
	}

	//if err := h.DB.First(&user, ).Error; err != nil {
	//	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid user in token"}
	//}
	fmt.Println("User id", user.ID)
	bookmarks := make([]model.Bookmark, 0)

	if err := h.DB.Where("user_id = ?", user.ID).Find(&bookmarks); err != nil {
		fmt.Println("There has been an error", err)
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func userIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims["id"])
	return uint(claims["id"].(float64))
}
