package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {

	ch := All(context.Background(), "Aang", []string{"../../text1.txt"})  

	s, ok := <-ch

	if !ok {
		t.Errorf(" function All error => %v", ok)
	}

	log.Println("=======>>>>>", s)

}

func TestAny_user(t *testing.T) {

	ch := Any(context.Background(), "Aang", []string{"../../text1.txt"})  

	s, ok := <-ch

	if !ok {
		t.Errorf(" function All error => %v", ok)
	}

	log.Println("---------------")
	log.Println("res.Phrase) => ", s.Phrase)
	log.Println("res.Line) => ", s.Line)
	log.Println("res.LineNum) => ", s.LineNum)
	log.Println("res.ColNum) => ", s.ColNum)
	log.Println("---------------")


}


