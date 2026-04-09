package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPercentiles(t *testing.T) {
	t.Run("empty values", func(t *testing.T) {
		_, err := Percentiles(PercentilesArgs{Values: []float64{}})
		require.Error(t, err)
	})

	t.Run("single value", func(t *testing.T) {
		result, err := Percentiles(PercentilesArgs{Values: []float64{42}})
		require.NoError(t, err)
		require.Equal(t, 1, result.Count)
		require.Equal(t, 42.0, result.Min)
		require.Equal(t, 42.0, result.P50)
		require.Equal(t, 42.0, result.P95)
		require.Equal(t, 42.0, result.P99)
		require.Equal(t, 42.0, result.Max)
		require.Equal(t, 42.0, result.Avg)
	})
}

func Test_percentile(t *testing.T) {
	t.Run("single value", func(t *testing.T) {
		values := []float64{42}
		require.Equal(t, 42.0, percentile(values, 50))
	})

	t.Run("multiple values", func(t *testing.T) {
		values := []float64{1, 2, 3, 4, 5}
		require.Equal(t, 3.0, percentile(values, 50))
	})

	t.Run("even number of values", func(t *testing.T) {
		values := []float64{1, 2, 3, 4}
		require.Equal(t, 2.5, percentile(values, 50))
	})
}
