package util

import (
	"testing"
)

func TestGetFormattingPatternsByLocal(t *testing.T) {
	resp := GetFormattingPatternsByLocal("en")

	if resp.Response.Code != 200{
		t.Fatal("GetFormattingPatternsByLocal failed!!!")
	}

	t.Log(resp.Data.Messages)
}