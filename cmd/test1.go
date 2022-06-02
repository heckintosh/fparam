package main

import (
	"fmt"
	
	"math"

	// "github.com/davecgh/go-spew/spew"
	// "encoding/json"
	// "strings"
	"reflect"
)


func main() {
	DictionarySlice := make([]map[string]string,0)
	input := map[string]string{
		"TEST1": "value1",
		"TEST2": "value2",
		"TEST3": "value3",
		"CITE":"HANOI",
		"CSOC":"value5",
	}
	var inp_key_slice []string
	var inp_val_slice []string
	var n = 3

	for k, v := range input {
		inp_key_slice = append(inp_key_slice, v)
		inp_val_slice = append(inp_val_slice, k)
	}

	// fmt.Println(listed)
	k := len(inp_key_slice) / n
	m := len(inp_key_slice) % n
	
	for i := 0; i < n; i++ {
		index_start := i*k + int(math.Min(float64(i), float64(m)))
		index_end := (i+1)*k + int(math.Min(float64(i+1), float64(m)))
		
		tmp :=make(map[string]string)
		
		for j := index_start; j < index_end; j++ {
			tmp[inp_key_slice[j]] = inp_val_slice[j]
		}
		// spew.Dump(tmp)
		fmt.Println(tmp)
		// result = append(result,array_dict)
		
		DictionarySlice = append(DictionarySlice,tmp)
	}
	// spew.Dump(DictionarySlice)
	fmt.Println(DictionarySlice)
	fmt.Println(reflect.TypeOf(DictionarySlice))
}



type Dict struct{
	key string 
	value string
}
type SliceDict struct{
	dicts [] Dict
}