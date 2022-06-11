package service_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/entity/mocks"
	"github.com/zhihaop/ticktok/user/service"
	"testing"
)

var userService entity.UserService

func init() {
	userService = service.NewUserService(
		mocks.NewMockUserRepository(),
		mocks.NewMockFollowRepository(),
	)
}

func TestUserServiceImpl_Register(t *testing.T) {
	token, err := userService.Register("testRegister", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	username, err := userService.GetUsername(token.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, username, "testRegister")
	t.Logf("register success, id: %d, token: %s", token.ID, token.Token)
}

func TestUserServiceImpl_Login(t *testing.T) {
	tokenRegister, err := userService.Register("testLogin", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	tokenLogin, err := userService.Login("testLogin", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tokenRegister, tokenLogin)
	t.Logf("login success, id: %d, token: %s", tokenLogin.ID, tokenLogin.Token)
}

func TestUserServiceImpl_GetUserID(t *testing.T) {
	token, err := userService.Register("testUserID", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	id, err := userService.GetUserID(token.Token)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, token.ID)
}

func TestUserServiceImpl_GetFollowCount(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func TestUserServiceImpl_GetFollowerCount(t *testing.T) {
	//TODO implement me
	panic("implement me")
}

func TestUserServiceImpl_IsFollow(t *testing.T) {
	//TODO implement me
	panic("implement me")
}
