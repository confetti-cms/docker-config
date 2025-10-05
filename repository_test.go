package dockerconfig

import (
	"testing"

	"github.com/matryer/is"
)

func TestRepository_no_request_does_not_match_any_granted(t *testing.T) {
	is := is.New(t)

	// Given
	requested := []Requested{}

	// When
	result := GetGranted(requested)

	// Then
	is.Equal(len(result), 0)
}
