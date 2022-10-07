package learn_go_one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSayHi(t *testing.T) {
	result := SayHi("eko")
	require.Equal(t, "hi eko", result, "ada yg salah")
}
