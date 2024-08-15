package shared

import (
	"github.com/gofrs/uuid"
	"github.com/jaevor/go-nanoid"
	"github.com/oklog/ulid/v2"
)

func GenerateUlid() string {
	return ulid.Make().String()
}

func GenerateUuid() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}

func GenerateNanoId() string {
	newNanoId, err := nanoid.CustomASCII("0123456789", 15)

	if err != nil {
		panic(err)
	}
	return newNanoId()
}
