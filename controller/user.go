package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/war3dgjddhh/douyin-dm/common"
	"github.com/war3dgjddhh/douyin-dm/repository"
	"golang.org/x/crypto/bcrypt"
)

var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := repository.GetUserByUsername(username)
	if user.Id != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 500, StatusMsg: "System Error"},
		})
		return
	}

	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := repository.User{
		Username: username,
		Password: string(hashedPassword),
		Name:     username,
		Avatar:   "/images/default_avatar.png",
	}
	repository.CreateUser(&newUser)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   newUser.Id,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := repository.GetUserByUsername(username)

	logrus.Infof("user is %v", user)

	// test user is exist
	if user.Id != 0 {

		// 判断密码是否正确
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{Response: Response{StatusCode: 422, StatusMsg: "User password incorrent"}})
			return
		}

		token, err := common.ReleaseToken(user)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{Response: Response{StatusCode: 500, StatusMsg: "System Error"}})
		}

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	user = Convert(user.(repository.User))
	if user.(User).Id != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user.(User),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func Convert(su repository.User) User {
	user := User{
		Id:            int64(su.Id),
		Name:          su.Name,
		FollowerCount: int64(su.FollowerCount),
		FollowCount:   int64(su.FollowCount),
	}
	return user

}
