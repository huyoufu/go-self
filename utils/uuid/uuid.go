package uuid

import _uuid "github.com/satori/go.uuid"

func Uuid() string {
	return _uuid.Must(_uuid.NewV4(), nil).String()
}
