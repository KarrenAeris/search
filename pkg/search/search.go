package search

import (
	"context"
	"os"
	"bufio"
	"log"
	"strings"
	"sync"
)

//Result описывает один результат поиска
type Result struct {
	// фраза, которую искали
	Phrase  string
	// целиком вся строкка, где нашли фразу (без \n или \r\n в конце)
	Line    string
	// номер строки (начиная с 1)
	LineNum int64
	// номер позиции (начиная с 1) - или символа? 
	ColNum  int64
}

//All ищет все появления фразы в текстовых файлах files
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}

	//var results []Result

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, path string, i int, ch chan<- []Result) {
			defer wg.Done()

			res := FindAllMatchTextInFile(phrase, path)

			if len(res) > 0 {
				ch <- res
			}

		}(ctx, files[i], i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()

	}()

	cancel()
	return ch
}

//FindAllMatchTextInFile ...
func FindAllMatchTextInFile(phrase, path string) (res []Result) {
    file, err := os.Open(path)
    if err != nil {
		log.Println("error not opened file err => ", err)
		return res
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
	}

	for i:=0; i <= len(lines); i++ {
		
		if strings.Contains(lines[i], phrase) {

			r := Result{
				Phrase:  phrase,
				Line:    lines[i],
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(lines[i], phrase)) + 1,
			}

			res = append(res, r)
		}
	}
    return res
}

	/*data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error not opened file err => ", err)
		return res
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	for i, line := range temp {
		
		if strings.Contains(line, phrase) {

			r := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}

			res = append(res, r)
		}
	}

	return res 
} */