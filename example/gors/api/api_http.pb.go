// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v1.0.0
// - protoc             v4.23.3
// source: api.proto

package api

// This is a compile-time assertion to ensure that this generated file
// is compatible with the panda package it is being compiled against.
import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/errors"
)

type UserService interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

func AddUserServiceRouters[R gin.IRoutes](iRoutes R, ctr UserService) R {
	iRoutes.Handle("GET", "/v1/users", _ListUsersHandler(ctr))
	iRoutes.Handle("GET", "/v1/users/{name}", _GetUserHandler(ctr))
	iRoutes.Handle("POST", "/v1/users", _CreateUserHandler(ctr))
	iRoutes.Handle("PATCH", "/v1/users", _UpdateUserHandler(ctr))
	iRoutes.Handle("DELETE", "/v1/users/{name}", _DeleteUserHandler(ctr))
	return iRoutes
}

func _ListUsersHandler(ctr UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &ListUsersRequest{}
		if err := c.BindQuery(in); err != nil {
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				return
			}
		}

		reply, err := ctr.ListUsers(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			} else {
				c.JSON(200, coder)
			}
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
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}
		if err := c.BindUri(in); err != nil {
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				return
			}
		}

		reply, err := ctr.GetUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			} else {
				c.JSON(200, coder)
			}
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
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}
		if err := c.BindQuery(in); err != nil {
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				return
			}
		}

		reply, err := ctr.CreateUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			} else {
				c.JSON(200, coder)
			}
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
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}
		if err := c.BindQuery(in); err != nil {
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				return
			}
		}

		reply, err := ctr.UpdateUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			} else {
				c.JSON(200, coder)
			}
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
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}
		if err := c.BindUri(in); err != nil {
			c.AbortWithStatusJSON(200, errors.BindCoder)
			return
		}

		if v, ok := interface{}(in).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				return
			}
		}

		reply, err := ctr.DeleteUser(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			} else {
				c.JSON(200, coder)
			}
			return
		}

		c.JSON(200, reply)
	}
}
