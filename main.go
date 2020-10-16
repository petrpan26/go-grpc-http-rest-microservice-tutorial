package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"log"
)

func main() {
	client, err := as.NewClient("172.17.0.2", 3000)
	if err != nil {
		log.Fatal("ok0" + err.Error())
	}
	key, err := as.NewKey("test", "set",
		"key value goes here and can be any supported primitive")
	if err != nil {
		log.Fatal("ok1" + err.Error())
	}

	bin1 := as.NewBin("bin1", "value1")
	bin2 := as.NewBin("bin2", "value2")

	// Write a record
	err = client.PutBins(nil, key, bin1, bin2)
	if err != nil {
		log.Fatal("ok2" + err.Error())
	}

	// Read a record
	record, err := client.Get(nil, key)
	if err != nil {
		log.Fatal("ok3" + err.Error())
	}

	log.Print(record.Bins)

	client.Close()
}
