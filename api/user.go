package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mbShopApi/global/response"
	"mbShopApi/proto"
	"net/http"
)

var (
	conn       *grpc.ClientConn
	userClient proto.UserClient
	ip         = "127.0.0.1"
	port       = 50051
)

func init() {
	var err error
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())

	if err != nil {
		zap.S().Panic("err:" + err.Error())
	}

	userClient = proto.NewUserClient(conn)
}

func GetUserList(c *gin.Context) {

	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
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
			Birthday: v.Birthday,
			Gender:   v.Gender,
			Role:     v.Role,
		}
		userList.Data = append(userList.Data, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})

}
