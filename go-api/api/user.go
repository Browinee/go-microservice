package api

import (
	"context"
	"fmt"
	"go-api/proto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		// NOTE: transform grpc code to http status code
		if e, ok :=	status.FromError(err); ok {
			 switch e.Code() {
			 	case codes.NotFound:
					c.JSON(http.StatusNotFound, gin.H{
						"msg": e.Message(),
					})
				case codes.Internal:
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "Internal error",
					})
				case codes.InvalidArgument:
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": "Wrong arguments",
					})
				case codes.Unavailable:
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "Unavailable",
					})
				default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg":"Other error"+ e.Message(),
				})
			 }
		}
	}
}
func GetUserList(ctx *gin.Context) {
	ip := "localhost"
	port := 50051
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] fail", "msg", err.Error())
	}
	userSrcClient := proto.NewUserClient(userConn)
	rsp, err := userSrcClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn: 0,
		PSize:0,
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] fail to get user list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {

		user := response.UserReponse{
			Id: value.id,
			Birthday: time.Time(time.Unix(int64(value.Birthday), 0)),
			Nickname: value.Nickname,
			gender: value.Gender,
			mobile: value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}