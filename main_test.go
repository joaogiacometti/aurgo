package main

import (
	"testing"

	"github.com/joaogiacometti/aurgo/helpers"
)

func TestRun_Help(t *testing.T) {
	config := helpers.FlagOptions{ShowHelp: true}
	err := run(config, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRun_ConflictTags(t *testing.T) {
	cases := []struct {
		name   string
		config helpers.FlagOptions
	}{
		{
			name:   "Install and Remove",
			config: helpers.FlagOptions{Install: true, Remove: true},
		},
		{
			name:   "Install and Update",
			config: helpers.FlagOptions{Install: true, Update: true},
		},
		{
			name:   "Remove and Update",
			config: helpers.FlagOptions{Remove: true, Update: true},
		},
		{
			name:   "Install and Remove and Update",
			config: helpers.FlagOptions{Install: true, Remove: true, Update: true},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := run(tc.config, nil)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	}
}

func TestRun_UpdateWithPackage(t *testing.T) {
	config := helpers.FlagOptions{Update: true}
	err := run(config, []string{"somepackage"})
	if err == nil {
		t.Errorf("Expected error for -U with package argument, got nil")
	}
}

func TestRun_NoPackageName(t *testing.T) {
	cases := []struct {
		name   string
		config helpers.FlagOptions
	}{
		{
			name:   "Search without package",
			config: helpers.FlagOptions{Search: true},
		},
		{
			name:   "Install without package",
			config: helpers.FlagOptions{Install: true},
		},
		{
			name:   "Remove without package",
			config: helpers.FlagOptions{Remove: true},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := run(tc.config, nil)
			if err == nil {
				t.Errorf("Expected error for %s with no package name, got nil", tc.name)
			}
		})
	}
}

func TestRun_NoArgs(t *testing.T) {
	config := helpers.FlagOptions{}
	err := run(config, nil)
	if err == nil {
		t.Errorf("Expected error for no arguments, got nil")
	}
}
