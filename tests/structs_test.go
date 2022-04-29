package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/vranyes/goranger/policy"
)

func TestStructUnmarshal(t *testing.T) {
	var policy policy.Policy
	dat, _ := os.ReadFile("sample-policy.json")
	_ = json.Unmarshal(dat, &policy)
}
