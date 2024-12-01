package handlers

import (
	"context"

	"github.com/platinumscatter/simple_api/internal/userService"
	"github.com/platinumscatter/simple_api/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		id := int(usr.ID)
		user := users.User{
			Id:        &id,
			Email:     &usr.Email,
			CreatedAt: &usr.CreatedAt,
			UpdatedAt: &usr.UpdatedAt,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}

	createdUser, err := u.Service.CreateUser(userToCreate, userToCreate)
	if err != nil {
		return nil, err
	}
	id := int(createdUser.ID)
	response := users.PostUsers201JSONResponse{
		Id:        &id,
		Email:     &createdUser.Email,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	}
	return response, nil
}

func (u *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userToUpdate := userService.User{
		Email:    string(*request.Body.Email),
		Password: *request.Body.Password,
	}

	userId := uint(request.Id)
	updatedUser, err := u.Service.UpdateUserByID(userId, userToUpdate)
	if err != nil {
		return nil, err
	}
	id := int(updatedUser.ID)
	response := users.PatchUsersId200JSONResponse{
		Id:        &id,
		Email:     &updatedUser.Email,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	}
	return response, nil
}

func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userId := uint(request.Id)
	err := u.Service.DeleteUserByID(userId)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
