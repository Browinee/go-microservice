package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"grpc/global"
	"grpc/model"
	"grpc/proto"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)



type UserServer struct {
	proto.UnimplementedUserServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    if page <= 0 {
      page = 1
    }

    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}
func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoResp := proto.UserInfoResponse{
		Id: int32(user.ID),
		Password: user.Password,
		Nickname: user.NickName,
		Gender: user.Gender,
		Role: int32(user.Role),

	}
	if user.Birthday != nil {
		// if there is attribute having default value in grpc message
		// we can't set nil
		userInfoResp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoResp
}


func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo)(*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil,result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *UserServer) 	GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error){
  var user model.User
	result := global.DB.Where((&model.User{Mobile:req.Mobile})).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User not exist.")
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}


func (s *UserServer) 	GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error){
  var user model.User
	result := global.DB.First(&user, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User not exist.")
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error){
	// check req
	// check if user existed
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Error(codes.AlreadyExists, "User existed")
	}
	// create user object
	user.Mobile = req.Mobile
	user.NickName = req.Nickname

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, &encodedPwd)
	user.Password = newPassword
	// save to db
	result = global.DB.Create((&user))
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	fmt.Print("---------1")
	userInfoRsp := ModelToResponse(user)
fmt.Print("---------3")
	return &userInfoRsp, nil

}
func (s *UserServer) 	CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckPasswordResponse, error){
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckPasswordResponse{
		Success: check,
	}, nil

}
func (s *UserServer)	UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not existed")
	}
	birthday := time.Unix(int64(req.Birthday), 0)
	user.NickName = req.Nickname
	user.Birthday = &birthday
	user.Gender = req.Gender
	result = global.DB.Save(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
 return &empty.Empty{}, nil
}