package main

import (
	"fmt"

	"k8s.io/metrics/pkg/apis/metrics"
)

func main() {
	fmt.Print(metrics.PodMetrics)
}
