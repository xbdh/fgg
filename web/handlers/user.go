package handlers

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	"time"
	"web/global"
	"web/middlewares"

	"web/models"

	//friendshipV1 "web/api/friendship/v1"
	//newsfeedV1 "web/api/newsfeed/v1"
	//tweetV1 "web/api/tweet/v1"
	accountV1 "web/api/user/v1"

)

func PassWordLogin(c *gin.Context) {
	passwordFrom := models.PassWordLoginForm{}
	if err := c.ShouldBindJSON(&passwordFrom); err != nil {
		log.Println("解析登录数据出错")
		return
	}

	resp, err := global.AccountCli.GetUserByName(context.Background(),&accountV1.NameRequest{
		Name: passwordFrom.Name,
	})

	if  err != nil {
		log.Println("没用此用户")
		c.JSON(400,gin.H{
			"message":"user not found",
		})
		return
	}

	if resp.Userinfo.Password!=passwordFrom.Password {
		log.Println("password wrong")
		c.JSON(400,gin.H{
			"message":"password wrong",
		})
		return

	}

	j := middlewares.NewJWT()

	claims := models.CustomClaims{
		UserId: resp.Userinfo.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 24*60*60*30,
				Issuer:    "rain",
				NotBefore: time.Now().Unix(),
			},
		}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "生产token失败",})
			return
	}

	c.JSON(http.StatusOK, gin.H{
			"id":        resp.Userinfo.Id,
			"token":     token,
			"ExpiresAt": (time.Now().Unix() + 24*60*60*30) * 1000,
		})
}



func Register(c *gin.Context) {

	registerForm := models.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		log.Println("解析表单错误")
		return
	}



	reply, err :=global.AccountCli.CreateUser(context.Background(),&accountV1.CreateRequest{
		Name:    registerForm.Name,
		Password: registerForm.Password,
		Email:    registerForm.Email,
	})

	if err != nil {
		log.Println("[Register] 查询 【新建用户失败】失败: %s", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"id":         reply.Userinfo.Id,
		"user":reply,
	})
}





// 将grpc的错误转换成http状态码
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

