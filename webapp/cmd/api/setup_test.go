package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var (
	app          application
	expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2NjQ1MTM2NDQsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.VjODp43PPUlT-nVDSrPNTaXhkhlNuyz9KhJaHEukFMs"
)

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "very secret"

	os.Exit(m.Run())
}
