package gocrud_gorm_test

import (
	"github.com/kordar/gocrud"
	"testing"
)

func TestCREATE(t *testing.T) {
	formBody := gocrud.NewFormBody("gorm")
	create, err := formBody.Create(nil, "", nil)
	if err != nil {
		t.Error(err)
	}
}
