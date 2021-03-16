package main

import (
	"encoding/json"
	"fmt"

	"k8s.io/metrics/pkg/apis/metrics"
)

func main() {
	fmt.Println(json.Marshal(metrics.PodMetricsList))
}
