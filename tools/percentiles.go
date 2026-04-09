// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package tools

import (
	"math"
	"sort"

	"github.com/pkg/errors"
)

// PercentilesArgs defines the input for the Percentiles function.
type PercentilesArgs struct {
	Values []float64 `json:"values"`
}

// PercentilesResult defines the output of the Percentiles function.
type PercentilesResult struct {
	Count int     `json:"count"`
	Min   float64 `json:"min"`
	P50   float64 `json:"p50"`
	P95   float64 `json:"p95"`
	P99   float64 `json:"p99"`
	Max   float64 `json:"max"`
	Avg   float64 `json:"avg"`
}

// Percentiles calculates the count, min, p50, p95, p99,
// max, and average of the given values.
func Percentiles(args PercentilesArgs) (PercentilesResult, error) {
	if len(args.Values) == 0 {
		return PercentilesResult{}, errors.New("values must not be empty")
	}

	values := make([]float64, len(args.Values))
	copy(values, args.Values)
	sort.Float64s(values)

	var sum float64
	for _, v := range values {
		sum += v
	}

	return PercentilesResult{
		Count: len(values),
		Min:   values[0],
		P50:   percentile(values, 50),
		P95:   percentile(values, 95),
		P99:   percentile(values, 99),
		Max:   values[len(values)-1],
		Avg:   sum / float64(len(values)),
	}, nil
}

// percentile calculates the p-th percentile of a sorted slice of float64 values.
func percentile(sorted []float64, p float64) float64 {
	if len(sorted) == 1 {
		return sorted[0]
	}

	pos := (p / 100.0) * float64(len(sorted)-1)
	lower := int(math.Floor(pos))
	upper := int(math.Ceil(pos))

	if lower == upper {
		return sorted[lower]
	}

	weight := pos - float64(lower)
	return sorted[lower] + weight*(sorted[upper]-sorted[lower])
}
