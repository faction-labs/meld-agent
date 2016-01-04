package main

import (
	"net/rpc"
)

func getClient(addr string) (*rpc.Client, error) {
	return rpc.DialHTTP("tcp", addr)
}
