package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mbShopApi/global"
	"mbShopApi/global/response"
	"mbShopApi/global/utils/timeutil"
	"mbShopApi/proto"
	"net/http"
	"time"
)

func InitConn() proto.UserClient {
	srvInfo := global.ServerConfig.UserSrvInfo

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", srvInfo.Host, srvInfo.Port), grpc.WithInsecure())

	if err != nil {
		zap.S().Panic("err:" + err.Error())
	}

	return proto.NewUserClient(conn)
}

func GetUserList(c *gin.Context) {

	userClient := InitConn()

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
