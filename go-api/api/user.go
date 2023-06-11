package api

import (
	"context"
	"fmt"
	"go-api/forms"
	"go-api/global"
	"go-api/global/response"
	"go-api/pkg/jwt"
	"go-api/proto"
	"go-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
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
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserServiceInfo.Host, global.ServerConfig.UserServiceInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] fail", "msg", err.Error())
	}
	userSrcClient := proto.NewUserClient(userConn)
	offset, limit := getPageInfo(ctx)
	rsp, err := userSrcClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn: uint32(offset),
		PSize:uint32(limit),
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] fail to get user list")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {

		user := response.UserReponse{
			Id: value.Id,
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			// Birthda y: time.Time(time.Unix(int64(value.Birthday), 0)),
			Nickname: value.Nickname,
			Gender: value.Gender,
			Mobile: value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}


func PassWordLogin(ctx *gin.Context) {
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
    zap.S().Errorf("valid %v", errs.Translate(global.Trans))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.RemoveTopStruct(errs.Translate(global.Trans)),
		})
		return
	}


	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"captcha":"captcha error",
		})
		return
	}


	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserServiceInfo.Host, global.ServerConfig.UserServiceInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[PasswordLogin] fail", "msg", err.Error())
	}
	userSrcClient := proto.NewUserClient(userConn)

	rsp, err := userSrcClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})

	if err != nil {
		zap.S().Errorw("[PassWordLogin] user not found")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

  if passRsp,pasErr := userSrcClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:passwordLoginForm.Password,
		EncryptedPassword: rsp.Password,
	}); pasErr != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"msg": "登錄失敗",
		})
	} else {
		 if passRsp.Success {
				accessToken, err := jwt.GenToken(rsp.Id, rsp.Nickname, rsp.Role)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"msg":"fail to create token"})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"id": rsp.Id,
					"token": accessToken,
				})
				return
		 } else {
			 ctx.JSON(http.StatusBadRequest,gin.H{
				 "msg":"Fail to login.",
			 })
			 return
		 }
	}

}



func Register(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBindJSON(&registerForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
    zap.S().Errorf("valid %v", errs.Translate(global.Trans))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": utils.RemoveTopStruct(errs.Translate(global.Trans)),
		})
		return
	}
zap.S().Info("before get validation code----")
	// validate code
	value, err := global.RedisClient.Get(registerForm.Mobile).Result()
	if err == redis.Nil {
		zap.S().Error("mobile key not existed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":"Wrong mobile",
		})
		return
	}

	if value != registerForm.Code {
		zap.S().Error("wrong code")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":"Wrong code",
		})
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserServiceInfo.Host, global.ServerConfig.UserServiceInfo.Port), grpc.WithInsecure())

	userSrcClient := proto.NewUserClient(userConn)
	user, userErr := userSrcClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Nickname: registerForm.Mobile,
		Password: registerForm.Password,
		Mobile: registerForm.Mobile,
	})
	zap.S().Infof("user created %v", user)
	if userErr != nil {
		zap.S().Errorf("[Register] fail creat user:  %s", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return

}