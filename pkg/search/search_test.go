package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {

	ch := All(context.Background(), "HTTP", []string{"../../text1.txt"})

	s, ok := <-ch

	if !ok {
		t.Errorf(" function All error => %v", ok)
	}

	log.Println("=======>>>>>", s)

}


