package main

import (
	"context"
	"runtime"
)

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.Di

}

func main() {

}
