package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	min := int64(1)
	max := int64(100)
	result := RandomInt(min, max)
	require.GreaterOrEqual(t, result, min)
	require.LessOrEqual(t, result, max)
}

func TestRandomString(t *testing.T) {
	n := 10
	result := RandomString(n)
	require.Len(t, result, n)

	for _, c := range result {
		require.Contains(t, alphabet, string(c))
	}
}

func TestRandomOwner(t *testing.T) {
	result := RandomOwner()
	require.Len(t, result, 8)

	// Verify that the result only contains characters from the alphabet
	for _, c := range result {
		require.Contains(t, alphabet, string(c))
	}
}

func TestRandomAmount(t *testing.T) {
	result := RandomAmount()
	require.GreaterOrEqual(t, result, 0)
	require.LessOrEqual(t, result, 1000000)
}

func TestRandomCurrency(t *testing.T) {
	currencies := []string{"USD", "EUR", "CAD"}
	result := RandomCurrency()
	require.Contains(t, currencies, result)
}
