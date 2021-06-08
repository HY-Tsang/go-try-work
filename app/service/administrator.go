package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"try-work/app/dao"
	"try-work/app/model"
)

// 中间件管理服务
var Administrator = administratorService{}

type administratorService struct{}

// 用户注册
func (s *administratorService) SignUp(r *model.AdministratorServiceSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.RealName == "" {
		r.RealName = r.Username
	}
	// 账号唯一性数据检查
	if !s.CheckUsername(r.Username) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Username))
	}
	// 昵称唯一性数据检查
	if !s.CheckRealName(r.RealName) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.RealName))
	}
	r.Status = 1
	r.CreatedAt = time.Now().Unix()
	r.UpdatedAt = time.Now().Unix()

	if _, err := dao.Administrator.Save(r); err != nil {
		return err
	}
	return nil
}

//
//// 判断用户是否已经登录
//func (s *userService) IsSignedIn(ctx context.Context) bool {
//	if v := Context.Get(ctx); v != nil && v.User != nil {
//		return true
//	}
//	return false
//}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *administratorService) SignIn(ctx context.Context, username, password string) error {
	admin, err := dao.Administrator.FindOne("username=? and password=?", username, password)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("账号或密码错误")
	}
	var adminModel *model.Administrator
	admin.Struct(&adminModel)
	if err := Session.SetUser(ctx, adminModel); err != nil {
		return err
	}
	Context.SetAdministrator(ctx, &model.ContextAdmin{
		Id:       adminModel.Id,
		Username: adminModel.Username,
		RealName: adminModel.RealName,
	})
	return nil
}

//// 用户注销
//func (s *userService) SignOut(ctx context.Context) error {
//	return Session.RemoveUser(ctx)
//}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *administratorService) CheckUsername(passport string) bool {
	if i, err := dao.Administrator.FindCount("username", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *administratorService) CheckRealName(nickname string) bool {
	if i, err := dao.Administrator.FindCount("real_name", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

//// 获得用户信息详情
//func (s *userService) GetProfile(ctx context.Context) *model.User {
//	return Session.GetUser(ctx)
//}
