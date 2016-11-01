package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

type User struct {
	ID   int
	Name string
}

func main() {
	db, err := bolt.Open("db/db1.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	acct11 := []byte("accountOne")
	usr11 := []byte("USERS")

	err = db.Update(func(tx *bolt.Tx) error {
		ab11, _ := tx.CreateBucketIfNotExists(acct11)
		ub11, _ := ab11.CreateBucketIfNotExists(usr11)

		fmt.Println("buckets")
		fmt.Println("accountONe", ab11)

		id, _ := ub11.NextSequence()
		u := User{ID: int(id), Name: "markTward"}
		ju, _ := json.Marshal(u)

		idstr := strconv.FormatUint(id, 10)
		ub11.Put([]byte(idstr), ju)

		return nil
	})

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(acct11)
		if bucket == nil {
			return fmt.Errorf("bucket not found: %v", string(acct11))
		}

		log.Println("attempt ForEach over account")
		err = bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("A %s is %s.\n", k, v)
			return nil
		})
		return nil

	})

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(acct11)
		users := bucket.Bucket(usr11)

		log.Println("attempt ForEach over users")
		err = users.ForEach(func(k, v []byte) error {
			fmt.Printf("A %s is %s.\n", k, v)
			return nil
		})
		return nil

	})

}
