package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/linkeunid/api.linkeun.com/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
