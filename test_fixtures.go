package main

import (
	"encoding/json"
	"log"
	"os"
)

type knownProduct struct {
	ProductID string `json:"productId"`
}

// A list of the product IDs our customers are known to use based on beacon data
// collected within the past month across all customers.
//
// See query_for_text_fixtures.sql for query ran.
// Query ran Dec 1st 2021.
var knownProductIDs []string

func init() {
	data, err := os.ReadFile("known_product_ids_top_16000.json")
	if err != nil {
		log.Fatal(err)
	}
	if len(data) == 0 {
		log.Fatal("Empty data from disk for known product IDs fixtures.")
	}

	var knownProducts []knownProduct
	if err := json.Unmarshal(data, &knownProducts); err != nil {
		log.Fatal(err)
	}
	if len(knownProducts) == 0 {
		log.Fatal("Empty list of known products after JSON unmarshal for known product IDs fixtures.")
	}

	knownProductIDStrings := []string{}
	for _, knownProduct := range knownProducts {
		knownProductIDStrings = append(knownProductIDStrings, knownProduct.ProductID)
	}

	knownProductIDs = knownProductIDStrings
}
