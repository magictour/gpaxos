package common

import "github.com/syndtr/goleveldb/leveldb/errors"

const (
	CRC32_SKIP = 8
)

// error code
var (
	ErrKeyNotFound       = errors.New("key not found")
	ErrGetFail           = errors.New("get fail")
	ErrInvalidGroupIndex = errors.New("invalid group index")
	ErrInvalidInstanceId = errors.New("invalid instanceid")
	ErrDbNotInit         = errors.New("db not init yet")
)
