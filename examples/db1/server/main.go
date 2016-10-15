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
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
	return s
}

// attempt lookup into db over range of keys / values
func (s *server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadReply, error) {
	// get exclusive lock on server, deferring close
	s.mu.Lock()
	defer s.mu.Unlock()

	// consturct values array for multiple keys
	var values []string
	for _, key := range in.Keys {
		values = append(values, s.db[key])
	}

	log.Printf("Read: Keys: %v\t Values: %v\n", in.Keys, values)
	desc, _ := in.Descriptor()
	log.Printf("Descriptor: %v\n", desc)

	// lookup key in db
	return &pb.ReadReply{Values: values}, nil
}

func (s *server) ServiceInfo(ctx context.Context, in *pb.ServiceInfoRequest) (*pb.ServiceInfoReply, error) {
	si, _ := in.Descriptor()
	fd, _ := decodeFileDesc(si)

	log.Println("Decoded:", fd)

	gs := fd.GetService()
	log.Println("Service:", gs)

	var md []string
	for _, mt := range fd.GetMessageType() {
		md = append(md, *mt.Name)
	}
	log.Println("metadata:", md)

	// for _, mt := range fd.GetMessageType() {
	// 	// fmt.Printf("message type: %T\t%v\n", mt, mt)
	// 	fmt.Printf("Name:\t%v\n", *mt.Name)
	// 	fmt.Println("Fields:")
	// 	for _, f := range mt.Field {
	// 		fmt.Printf("\t%v\n", f)
	// 	}
	// }

	return &pb.ServiceInfoReply{Values: md}, nil
}

// decompress does gzip decompression.
func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v\n", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v\n", err)
	}

	return out, nil
}

func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(enc)
	fmt.Println("Decompressed:", raw)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress enc: %v", err)
	}

	fd := new(dpb.FileDescriptorProto)

	if err := proto.Unmarshal(raw, fd); err != nil {
		return nil, fmt.Errorf("bad descriptor: %v", err)
	}
	return fd, nil
}

// attempt key/value insert/update into db
func (s *server) Upsert(ctx context.Context, in *pb.UpsertRequest) (*pb.UpsertReply, error) {
	// get exclusive lock on server, deferring close
	s.mu.Lock()
	defer s.mu.Unlock()

	// assign value to key in db
	s.db[in.Key] = in.Value
	log.Printf("Upsert: Key: %v\t Value: %v\n", in.Key, in.Value)
	desc, _ := in.Descriptor()
	log.Printf("Descriptor: %v\n", desc)

	return &pb.UpsertReply{Value: s.db[in.Key]}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// declare new grpc server using db service
	s := grpc.NewServer()
	pb.RegisterDBServiceServer(s, newDBServer())

	// // debug output for service
	// si := s.GetServiceInfo()

	// start server
	s.Serve(lis)
}
