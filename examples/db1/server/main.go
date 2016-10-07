/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "github.com/markTward/grpc-demo/examples/db1/grpc/db"
)

const (
	port = ":50052"
)

// server is used to implement db service.
type server struct {
	mu sync.Mutex
	db map[string]string
}

// create and seed database
func newDBServer() *server {
	s := new(server)
	s.db = make(map[string]string)
	s.db["hello"] = "world"
	return s
}

// attempt key/value lookup into db
func (s *server) Read(ctx context.Context, in *pb.RecordKey) (*pb.RecordValue, error) {
	s.mu.Lock()
	v := s.db[in.Key]
	s.mu.Unlock()
	return &pb.RecordValue{Value: v}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// declare new grpc server using db service
	s := grpc.NewServer()
	pb.RegisterRecordReaderServer(s, newDBServer())

	// debug output for service
	fmt.Println(s.GetServiceInfo())

	// start server
	s.Serve(lis)
}
