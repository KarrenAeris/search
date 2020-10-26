package main

import (
	s "strings"
	// "github.com/KarrenAeris/search/pkg/search"
	"fmt"
)

var p = fmt.Println

func main() { 
	p("Contains:  ", s.Contains("test", "es"))
}
	
// 	 res := search.FindAllMatchTextInFile("HTTP", "./text1.txt")

// 	for _, r := range res {

// 		fmt.Println("---------------")
// 		fmt.Println("res.Phrase) => ", r.Phrase)
// 		fmt.Println("res.Line) => ", r.Line)
// 		fmt.Println("res.LineNum) => ", r.LineNum)
// 		fmt.Println("res.ColNum) => ", r.ColNum)
// 		fmt.Println("---------------")
// 	}

// 	fmt.Println(strings.Index("В случае, когда дела не  обстоят таким образом, можно обратиться к другим методам скрейпинга.", "другим методам скрейпинга."))

	
// 	s1:=""
// 	s2:="другим методам скрейпинга."

// 	 r1 := strings.Index(s1, s2)
// 	fmt.Println(r1) 
// 	fmt.Println(find(s2, s1)) 
// }
// func find(target string, text string)int64{
// 	if(len(text) < len(target)){
// 		return -1;
// 	}
// 	for i := len(target); i <= len(text); i ++{
// 		if(text[i - len(target):i] == target){
// 			return int64(i - len(target));
// 		}
// 	}
// 	return -1;
// }
