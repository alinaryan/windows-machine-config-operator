package prometheus

import(
	v1 "k8s.io/api/core/v1"
)

// 1. given a node, find its ip
// 2. create an endpoint object
// 3. service in main.go