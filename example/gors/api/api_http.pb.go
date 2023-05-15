// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.9
// source: api.proto

package api

// This is a compile-time assertion to ensure that this generated file
// is compatible with the panda package it is being compiled against.
import (
	"context"
	"github.com/gin-gonic/gin"
	httpserver "github.com/go-leo/leo/runner/net/http/server"
	"github.com/go-leo/errors"
)

type UserService interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

func UserServiceAPIRouters(ctr UserService) []httpserver.Router {
	return []httpserver.Router{
		{
			HTTPMethod:   "GET",
			Path:         "/v1/users",
			HandlerFuncs: []gin.HandlerFunc{_ListUsersHandler(ctr)},
		},
		{
			HTTPMethod:   "GET",
			Path:         "/v1/users/{name}",
			HandlerFuncs: []gin.HandlerFunc{_GetUserHandler(ctr)},
		},
		{
			HTTPMethod:   "POST",
			Path:         "/v1/users",
			HandlerFuncs: []gin.HandlerFunc{_CreateUserHandler(ctr)},
		},
		{
			HTTPMethod:   "PATCH",
			Path:         "/v1/users",
			HandlerFuncs: []gin.HandlerFunc{_UpdateUserHandler(ctr)},
		},
		{
			HTTPMethod:   "DELETE",
			Path:         "/v1/users/{name}",
			HandlerFuncs: []gin.HandlerFunc{_DeleteUserHandler(ctr)},
		},
	}
}

func _ListUsersHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &ListUsersRequest{}
		if err := c.BindQuery(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.JSON(200, errors.ValidationCoder)
				c.Abort()
				return
			}
		}

		reply, err := ctr.ListUsers(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			return
		}

		c.JSON(200, reply)
	}
}

func _GetUserHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &GetUserRequest{}
		if err := c.BindQuery(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}
		if err := c.BindUri(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.JSON(200, errors.ValidationCoder)
				c.Abort()
				return
			}
		}

		reply, err := ctr.GetUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			return
		}

		c.JSON(200, reply)
	}
}

func _CreateUserHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &CreateUserRequest{}
		if err := c.Bind(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}
		if err := c.BindQuery(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.JSON(200, errors.ValidationCoder)
				c.Abort()
				return
			}
		}

		reply, err := ctr.CreateUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			return
		}

		c.JSON(200, reply)
	}
}

func _UpdateUserHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &UpdateUserRequest{}
		if err := c.Bind(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}
		if err := c.BindQuery(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.JSON(200, errors.ValidationCoder)
				c.Abort()
				return
			}
		}

		reply, err := ctr.UpdateUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			return
		}

		c.JSON(200, reply)
	}
}

func _DeleteUserHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &DeleteUserRequest{}
		if err := c.BindQuery(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}
		if err := c.BindUri(in); err != nil {
			c.JSON(200, errors.BindCoder)
			c.Abort()
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.JSON(200, errors.ValidationCoder)
				c.Abort()
				return
			}
		}

		reply, err := ctr.DeleteUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			return
		}

		c.JSON(200, reply)
	}
}
