package handlers

import (
	"testing"

	"wichitaradar/internal/testutils"
)

func TestHandleSatellite(t *testing.T) {
	testutils.InitTemplates(t)
	testutils.TestHandler(t, HandleSatellite, "/satellite")
}