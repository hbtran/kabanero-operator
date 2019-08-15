package collection

import (
	"testing"
)

func TestResolveIndex(t *testing.T) {
	index, err := ResolveIndex("https://raw.githubusercontent.com/kabanero-io/kabanero-collection/master/experimental")
	if err != nil {
		t.Fatal(err)
	}

	if index == nil {
		t.Fatal("Returned index was nil")
	}

	if index.ApiVersion != "v1" {
		t.Fatal("Expected apiVersion == v1")
	}
}

func TestResolveCollection(t *testing.T) {
	collection, err := ResolveCollection("https://raw.githubusercontent.com/kabanero-io/kabanero-collection/master/experimental/java-microprofile-0.2.1/collection.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if collection == nil {
		t.Fatal("Collection was nil")
	}

	t.Log(collection)
}