package io_test

import (
	"bytes"
	"salmon/src/io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadFrom(t *testing.T) {
	input := "hello world"

	output := io.ReadFrom(bytes.NewReader([]byte(input)))

	if diff := cmp.Diff(output, input); diff != "" {
		t.Fatalf(diff)
	}
}
