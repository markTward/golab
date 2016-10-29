package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {

	b := []byte{31, 139, 8, 0, 0, 9, 110, 136, 2, 255, 84, 144, 65, 79, 194, 64, 16, 133, 45, 96, 67, 135, 16, 116, 98, 12, 225, 164, 245, 66, 98, 220, 4, 61, 24, 175, 13, 63, 128, 84, 253, 1, 219, 238, 4, 9, 109, 186, 238, 46, 40, 255, 222, 217, 45, 104, 189, 205, 188, 253, 222, 203, 219, 129, 161, 42, 132, 54, 141, 107, 176, 167, 138, 244, 22, 70, 57, 73, 149, 211, 231, 142, 172, 67, 132, 193, 150, 14, 118, 26, 221, 244, 231, 73, 30, 230, 244, 14, 146, 22, 209, 213, 1, 175, 33, 222, 203, 138, 217, 35, 114, 220, 210, 103, 24, 191, 107, 75, 198, 157, 146, 46, 160, 207, 110, 166, 34, 166, 252, 136, 87, 112, 30, 224, 105, 47, 104, 237, 194, 233, 163, 147, 209, 231, 255, 66, 81, 7, 122, 36, 72, 150, 217, 43, 153, 253, 166, 36, 156, 195, 192, 247, 193, 137, 224, 159, 116, 202, 207, 198, 127, 2, 71, 165, 103, 40, 32, 110, 179, 241, 210, 63, 253, 43, 56, 155, 116, 165, 192, 103, 47, 112, 95, 54, 181, 88, 111, 220, 199, 174, 16, 181, 52, 219, 183, 47, 105, 148, 88, 27, 93, 62, 40, 170, 27, 65, 223, 178, 214, 21, 89, 246, 46, 178, 225, 50, 91, 172, 252, 37, 87, 81, 17, 135, 147, 62, 253, 4, 0, 0, 255, 255, 237, 111, 58, 134, 94, 1, 0, 0}
	fmt.Println("Bytes", b)

	fd, _ := decodeFileDesc(b)
	fmt.Println("Decoded:", fd)

	fmt.Println("Service:", fd.GetService())

	for _, mt := range fd.GetMessageType() {
		// fmt.Printf("message type: %T\t%v\n", mt, mt)
		fmt.Printf("Name:\t%v\n", *mt.Name)
		fmt.Println("Fields:")
		for _, f := range mt.Field {
			fmt.Printf("\t%v\n", f)
		}
	}

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
