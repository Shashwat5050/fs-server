package error

import (
	"errors"
	"fmt"
)

type ControllerAlreadyExists struct {
	Name string
}

func (e ControllerAlreadyExists) Error() string {
	return fmt.Sprintf("%s already exists", e.Name)
}

type ResourceNotFound struct {
	Resource string
}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("resource %s not found", e.Resource)
}

type BadJWT struct {
	Msg string
}

func (e BadJWT) Error() string {
	return fmt.Sprintf("JWT verification failed - %s", e.Msg)
}

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrIsClosed         = errors.New("closed")
	ErrInvalidInput     = errors.New("invalid input")
	ErrAlreadyExists    = errors.New("already exists")
	ErrAuthentication   = errors.New("authentication failed")
	ErrNotImplemented   = errors.New("not implemented")
	ErrEmptyQueue       = errors.New("queue is empty")
	ErrBackupLocked     = errors.New("backup is locked")
)
