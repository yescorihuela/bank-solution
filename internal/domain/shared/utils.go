package shared

import (
	"github.com/jaevor/go-nanoid"
	"github.com/oklog/ulid/v2"
)

func NewUlid() string {
	return ulid.Make().String()
}

func NewNanoId() string {
	newNanoId, err := nanoid.CustomASCII("0123456789", 12)
	if err != nil {
		panic(err)
	}
	return newNanoId()
}
