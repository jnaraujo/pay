package errs

import "errors"

var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrUserNotFound = errors.New("user not found")
var ErrGameNotFound = errors.New("game not found")
var ErrInvalidTransfer = errors.New("invalid transfer")
