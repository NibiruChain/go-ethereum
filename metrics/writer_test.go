package metrics

import (
	"sort"
	"testing"
)

func TestMetricsSorting(t *testing.T) {
	var namedMetrics = []namedMetric{
		{name: "zzz"},
		{name: "bbb"},
		{name: "fff"},
		{name: "ggg"},
	}

	sort.Slice(namedMetrics, func(i, j int) bool {
		return namedMetrics[i].cmp(namedMetrics[j]) < 0
	})
	for i, name := range []string{"bbb", "fff", "ggg", "zzz"} {
		if namedMetrics[i].name != name {
			t.Fail()
		}
	}
}
