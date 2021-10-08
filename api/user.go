package api

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mbShopApi/forms"
	"mbShopApi/global"
	"mbShopApi/global/response"
	"mbShopApi/middlewares"
	"mbShopApi/models"
	"mbShopApi/proto"
	"mbShopApi/utils/timeutil"
	"net/http"
	"time"
)

func GetUserList(c *gin.Context) {

	rsp, err := global.UserClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    0,
		PSize: 10,
	})

	if err != nil {
		zap.S().Panic(err.Error())
		return
	}

	if rsp == nil || len(rsp.Data) == 0 {
		zap.S().Info("empty")
		return
	}

	var userList response.UserList

	userList.Total = rsp.Total

	for _, v := range rsp.Data {
		user := &response.UserInfo{
			Id:       v.Id,
			Mobile:   v.Mobile,
			NickName: v.NickName,
			Birthday: timeutil.JsonTime(time.Unix(int64(v.Birthday), 0)),
			Gender:   v.Gender,
			Role:     v.Role,
		}
		userList.Data = append(userList.Data, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})

}

func PasswordLogin(c *gin.Context) {
	var form = forms.LoginForm{}
	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//校验验证码
	if !store.Verify(form.CaptchaId, form.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	//登录逻辑
	userClient := global.UserClient
	if rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: form.Mobile}); err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "用户不存在",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "登录失败:" + err.Error(),
			})
		}
	} else {
		if passRsp, passErr := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          form.Password,
			EncryptedPassword: rsp.Password,
		}); passErr == nil && passRsp.Status {
			j := middlewares.NewJWT()
			timeStamp := time.Now().Unix()
			claims := models.CustomClaims{
				ID:          rsp.Id,
				NickName:    rsp.NickName,
				AuthorityId: rsp.Role,
				StandardClaims: jwt.StandardClaims{
					NotBefore: timeStamp,
					ExpiresAt: timeStamp * 60 * 60 * 24,
					Issuer:    "cowboy",
				},
			}
			token, err := j.CreateToken(claims)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "token生成失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"id":        rsp.Id,
				"nickName":  rsp.NickName,
				"expiresAt": timeStamp * 60 * 60 * 24 * 1000,
				"token":     token,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "登录失败",
			})
		}
	}

}
