package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := uint64(10) // per second
	duration := 500 * time.Second

	userTarget := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/page/users",
	}
	productTarget := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/page/product",
	}
	aboutTarget := vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/page/about",
	}

	targeter := vegeta.NewStaticTargeter(userTarget, productTarget, aboutTarget)
	//targeter := vegeta.NewStaticTargeter(userTarget)

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
