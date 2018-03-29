package mimedb

import (
	"testing"
)

func TestIsCompressibleContentType(t *testing.T) {
	tests := map[string]bool{
		"application/json":       true,
		"application/javascript": true,
		"text/css":               true,
		"text/html":              true,

		"application/zip": false,
		"image/png":       false,
		"image/jpeg":      false,
	}

	for input, expected := range tests {
		output := IsCompressibleContentType(input)
		if output != expected {
			t.Errorf("IsCompressibleContentType('%s') => %v but expected %v", input, output, expected)
		}
	}
}
