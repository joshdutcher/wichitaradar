package handlers

import (
	"testing"

	"wichitaradar/internal/testutils"
)

func TestHandleHome(t *testing.T) {
	testutils.InitTemplates(t)
	testutils.TestHandler(t, HandleHome, "/")
}