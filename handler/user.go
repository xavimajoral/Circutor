package handler

import (
	"fmt"
	"frontend-test-api/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Signup godoc
// @Summary      Signup
// @Description  Signs up a user
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.User
// @Router       /signup [get]
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
// @Param   	user  body     model.UserSignup     yes  "user signup"
// @Success      200  {object}  model.User
// @Router       /login [post]
func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	if has, err := h.DB.Get(u); err != nil {
		fmt.Println(err)
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	} else {
		fmt.Println(has)
	}

	//if err := h.DB.Where("email = ? AND password = ?", u.Email, u.Password).First(&u).Error; err != nil {
	//	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	//}

	//-----
	// JWT
	//-----

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

func (h *Handler) SitesAdd(c echo.Context) (err error) {
	site := new(model.Site)
	site.UserID = userIDFromToken(c)
	if err = c.Bind(site); err != nil {
		return
	}

	if affected, err := h.DB.Insert(site); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(affected)
	}
	fmt.Println("Created site with id", site.ID)
	fmt.Println(err)

	return c.JSON(http.StatusOK, site)

}

func (h *Handler) SitesList(c echo.Context) (err error) {
	var user = model.User{ID: userIDFromToken(c)}
	if has, err := h.DB.Get(&user); err != nil && !has {
		fmt.Println("There has been an error", err, has)
	}

	//if err := h.DB.First(&user, ).Error; err != nil {
	//	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid user in token"}
	//}
	fmt.Println("User id", user.ID)
	var sites []model.Site

	if err := h.DB.Where("user_id = ?", user.ID).Find(&sites); err != nil {
		fmt.Println("There has been an error", err)
	}
	fmt.Println(sites)

	return c.JSON(http.StatusOK, sites)
}

func userIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims["id"])
	return uint(claims["id"].(float64))
}
