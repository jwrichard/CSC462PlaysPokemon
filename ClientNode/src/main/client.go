package main

import (
	"mapreduce"
	"container/list"
	"strings"
	"strconv"
	"unicode"
)

const (
	MASTER =	"54.213.250.201:1337"
	SELF =		"54.201.255.6:1337"
	PORT =		":1337"
)

//
//	The mapper
//
func Map(value string) *list.List {

	fs := strings.FieldsFunc(value, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	result := new(list.List)
	counts := createMovesMap()

  // Count all our moves.
	for _, line := range fs {
		word := strings.ToLower(line)
    if val, ok := counts[word]; ok {
      counts[word] = val + 1
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

//
//	The reducer
//
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

func main() {
	mapreduce.RunWorker(MASTER, SELF, PORT, Map, Reduce, 100)
}
