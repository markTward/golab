package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func main() {
	db, err := bolt.Open("db/db1.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ub := []byte("user")

	// iterate over entire DB (manual ForEach)
	fmt.Println("Cursor as ForEach over DB")
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ub)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s / value: %s\n", k, v)
		}

		return nil

	})

	// choose a range of IDs
	fmt.Println("Cursor as Range")
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(ub)
		c := b.Cursor()

		min := []byte(itob(8))
		max := []byte(itob(10))

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("[]byte key: %v / value: %v\n", k, v)
			fmt.Printf("key: %s / value: %s\n", k, v)
		}

		return nil

	})

}
