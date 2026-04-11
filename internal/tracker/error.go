package tracker

import "errors"

var ErrNotFound = errors.New("item not found")
var ErrHasAlreadyExist = errors.New("item has already exist")
