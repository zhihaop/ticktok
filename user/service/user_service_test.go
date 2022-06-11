package service

import (
	"github.com/go-playground/assert/v2"
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/entity/mocks"
	"testing"
)

var (
	repository entity.UserRepository
	service    entity.UserService
)

func init() {
	repository = mocks.NewMockUserRepository()
	service = NewUserService(repository, nil)
}

func TestUserServiceImpl_Register(t *testing.T) {
	token, err := service.Register("testRegister", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.FindByUsername("testRegister")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, user.ID, token.ID)
	assert.Equal(t, user.Name, "testRegister")
	assert.Equal(t, user.Token, token.Token)
	t.Logf("register success, id: %d, token: %s", token.ID, token.Token)
}

func TestUserServiceImpl_Login(t *testing.T) {
	tokenRegister, err := service.Register("testLogin", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	tokenLogin, err := service.Login("testLogin", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tokenRegister, tokenLogin)
	t.Logf("login success, id: %d, token: %s", tokenLogin.ID, tokenLogin.Token)
}

func TestUserServiceImpl_GetUserID(t *testing.T) {
	token, err := service.Register("testUserID", "passwd")
	if err != nil {
		t.Fatal(err)
	}

	id, err := service.GetUserID(token.Token)
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
