package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type User struct {
	ID   int
	Name string
}

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
	err = db.Update(func(tx *bolt.Tx) error {

		fmt.Println("Update Transation Info")
		fmt.Println(tx.DB(), "database")
		fmt.Println(tx.ID(), "ID")
		fmt.Println(tx.Size(), "Size")
		fmt.Println(tx.Stats(), "Stats")
		fmt.Println(tx.Writable(), "/", tx.WriteFlag, "Writable / Writeflag")
		fmt.Printf("Tx(%T): %#v\n", tx, tx)

		b, err := tx.CreateBucketIfNotExists(ub)
		if err != nil {
			return err
		}
		fmt.Println("Update Bucket Meta INfo")
		fmt.Println(b.FillPercent, "fill percent")
		fmt.Println(b.Writable(), "Writeable")
		fmt.Println(b.Stats(), "Stats")
		fmt.Printf("Tx(): %#v\n", b.Tx())
		fmt.Println("bucket seq", b.Sequence())

		// assigned ID and create User
		id, _ := b.NextSequence()
		u := User{ID: int(id), Name: "mtw"}
		fmt.Println("User Init", u)

		buf, err := json.Marshal(u)
		log.Println("json as bytes", buf)

		fmt.Println("ID as byte key", itob(u.ID))
		err = b.Put(itob(u.ID), buf)
		return nil
	})

	// access bucket and READ all pairs
	// err = db.View(func(tx *bolt.Tx) error {
	// 	fmt.Println("View Transation Info")
	// 	fmt.Println(tx.DB(), "database")
	// 	fmt.Println(tx.ID(), "ID")
	// 	fmt.Println(tx.Size(), "Size")
	// 	fmt.Println(tx.Stats(), "Stats")
	// 	fmt.Println(tx.Writable(), "/", tx.WriteFlag, "Writable / Writeflag")
	//
	// 	bucket := tx.Bucket(ub)
	// 	if bucket == nil {
	// 		return fmt.Errorf("bucket not found: %v", string(ub))
	// 	}
	//
	// 	log.Println("attempt ForEach over entire DB")
	// 	err = bucket.ForEach(func(k, v []byte) error {
	// 		fmt.Printf("A %s is %s.\n", k, v)
	// 		return nil
	// 	})
	// 	return nil
	//
	// })

}
