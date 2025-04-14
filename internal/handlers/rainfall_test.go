package handlers

import (
	"testing"

	"wichitaradar/internal/testutils"
)

func TestHandleRainfall(t *testing.T) {
	testutils.InitTemplates(t)
	testutils.TestHandler(t, HandleRainfall, "/rainfall")
}