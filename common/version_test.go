package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVersion(t *testing.T) {
	v1, _ := NewVersion("0.1.0")
	v2, _ := NewVersion("0.0.1")
	require.True(t, v2.IsLessThan(v1))
}
