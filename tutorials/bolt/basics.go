package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	var wbuck = []byte("world")

	db, err := bolt.Open("db/db1.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	log.Println("db", db)

	key := []byte("hello")
	value := []byte("hello world world")

	// create bucket and PUT k/v pair
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(wbuck)
		if err != nil {
			return err
		}
		log.Println("bucket ==>", bucket)

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte("dog"), []byte("fun")); err != nil {
			return err
		}
		if err := bucket.Put([]byte("cat"), []byte("lame")); err != nil {
			return err
		}
		if err := bucket.Put([]byte("liger"), []byte("awesome")); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	// access bucket and READ k/v pair
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(wbuck)
		if bucket == nil {
			return fmt.Errorf("bucket not found: %v", string(wbuck))
		}

		val := bucket.Get(key)

		if val == nil {
			fmt.Printf("nil value for key: %v", key)
		} else {
			fmt.Printf("found key / value pair: %v / %v\n", string(key), string(val))
		}

		log.Println("attempt ForEach over entire DB")
		err = bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("A %s is %s.\n", k, v)
			return nil
		})
		return nil

	})

	if err != nil {
		log.Fatalln(err)
	}

	// PUT / GET struct using JSON
	type Post struct {
		Created time.Time
		Title   string
		Content string
	}

	post := &Post{time.Now(), "First Post", "can it really work?"}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(wbuck)
		if err != nil {
			return err
		}

		jpost, err := json.Marshal(post)

		err = bucket.Put([]byte("jpost"), jpost)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(wbuck)
		if bucket == nil {
			return fmt.Errorf("bucket not found: %v", string(wbuck))
		}

		val := bucket.Get([]byte("jpost"))
		log.Println("get jpost ==>", val)

		var dat map[string]interface{}
		json.Unmarshal(val, &dat)
		log.Println("unmarshalled jpost into dat interface{}", dat)

		if val == nil {
			fmt.Printf("nil value for key: %v", key)
		} else {
			fmt.Printf("found key / value pair: %v / %v\n", string(key), string(val))
		}
		return nil

	})

	if err != nil {
		log.Fatalln(err)
	}

	// // DELETE k/v pair
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket(wbuck)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Println("bucket ==>", bucket)
	//
	// 	err = bucket.Delete(key)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Println("attempt DELETE", err)
	//
	// 	val := bucket.Get(key)
	//
	// 	if val == nil {
	// 		fmt.Printf("nil value for key after DELETE: %v", key)
	// 	} else {
	// 		fmt.Printf("found after DELETE key / value pair: %v / %v\n", string(key), string(val))
	// 	}
	//
	// 	return nil
	// })
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// // DELETE bucket
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket(wbuck)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Println("bucket ==>", bucket)
	//
	// 	err = bucket.DeleteBucket([]byte("123"))
	// 	if err != nil {
	// 		// log.Println("attempt DELETE bucket", err)
	// 		return err
	// 	}
	//
	// 	val := bucket.Get(key)
	//
	// 	if val == nil {
	// 		fmt.Printf("nil value for key after DELETE: %v", key)
	// 	} else {
	// 		fmt.Printf("found after DELETE key / value pair: %v / %v\n", string(key), string(val))
	// 	}
	//
	// 	return nil
	// })
	//
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
