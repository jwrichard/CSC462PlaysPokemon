package main

import "os"
import "fmt"
import "mapreduce"
import "container/list"
import "strings"
import "strconv"
import "unicode"

func Map(value string) *list.List {

	fs := strings.FieldsFunc(value, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	result := new(list.List)
	counts := createMovesMap()

  // Count all our moves.
	for _, wd := range fs {
    if val, ok := counts[wd]; ok {
      counts[wd] = val + 1
    }
	}

  //
  //  Push back all key value pairs for each movement.
  //  "up" : 1,
  //  "down": 0...
  //
	for k, v := range counts {
		result.PushBack(mapreduce.KeyValue{Key: k, Value: strconv.Itoa(v)})
	}
	return result
}

//
//  Create a map of all valid moves for pokemon.
//
func createMovesMap() map[string]int {
    movesMap := make(map[string]int)

    // Movement
    movesMap["up"] = 0
    movesMap["down"] = 0
    movesMap["left"] = 0
    movesMap["right"] = 0

    // Actions
    movesMap["start"] = 0
    movesMap["select"] = 0
    movesMap["a"] = 0
    movesMap["b"] = 0

    return movesMap
}

func Reduce(key string, values *list.List) string {

	count := 0
	for e := values.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(string)
		if ok {
			segment, err := strconv.Atoi(value)
			if err == nil {
				count += segment
			}
		}
	}
	return strconv.Itoa(count)

}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master x.txt sequential)
// 2) Master (e.g., go run wc.go master x.txt localhost:7777)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	fmt.Println("Welcome to my MapReduce!")

	if len(os.Args) != 4 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	} else if os.Args[1] == "master" {
		if os.Args[3] == "sequential" {
			mapreduce.RunSingle(5, 3, os.Args[2], Map, Reduce)
		} else {
			mr := mapreduce.MakeMapReduce(5, 3, os.Args[2], os.Args[3])
			// Wait until MR is done
			<-mr.DoneChannel
		}
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], Map, Reduce, 100)
	}
}
