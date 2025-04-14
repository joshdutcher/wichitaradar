package handlers

import (
	"testing"

	"wichitaradar/internal/testutils"
)

func TestHandleTemperatures(t *testing.T) {
	testutils.InitTemplates(t)
	testutils.TestHandler(t, HandleTemperatures, "/temperatures")
}