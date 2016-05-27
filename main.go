package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
)


func main() {

	addr := "http://tmi.twitch.tv/group/user/officialfriberg/chatters"
	resp, err := http.Get(addr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    response := parseObjectFromJSON(bytes)
	chatters := getObj("chatters", response)
	viewers := getArr("staff", chatters)
	for i, v := range viewers {
		fmt.Println(i, v)
	}
}

// Go types == json types
// map[string]interface{} == object { }
// []interface{} == array []
// float64 == number
// string == string
// bool == bool
// nil == null
func getObj(keyThatWereSearchingFor string, obj map[string]interface{}) map[string]interface{} {

	for key, uncastedvalue := range obj {
		switch value := uncastedvalue.(type) {
		case map[string]interface{}:
			if key == keyThatWereSearchingFor {
				return value
			}
		}
	}
	return nil
}

func getArr(keyThatWereSearchingFor string, obj map[string]interface{}) []interface{} {

	for key, uncastedvalue := range obj {
		switch value := uncastedvalue.(type) {
		case []interface{}:
			if key == keyThatWereSearchingFor {
				return value
			}
		}
	}
	return nil
}

func parseObjectFromJSON([]byte bytes) map[string]interface{} {
	var data interface{}
    err = json.Unmarshal(bytes, &data)
    if err != nil {
		fmt.Printf("Error parsing JSON to an object: ", err)
	}

	return data.(map[string]interface{})
}


//
/*

"a"
"aa"
"aaa"
"aaaa"
"aaaaa"
"ab"
"aba"
"abaa"
"abaaa"
*/

/**/

func guessIt(){
	chars := []string{"a", "b", "c", "d", "e"}
	secretpassword := "edcba"
	str_curr := 1
	str := ""
	for i := 0; i < 5; i++ {
		fmt.Println("i", i)

		for j := 0; j < 5; j++ {

			fmt.Println("j", j)

			if str_curr == len(str) {

				if str == secretpassword {
					fmt.Printf("Busted!!!")
				} else {
					fmt.Println(str)
					str = ""
					str_curr++
					fmt.Println("curr", str_curr)
					if str_curr == 6 {
						str_curr = 1
					}
					

				}
			} else {
				str += chars[i]
				fmt.Println(str)
			}


		}
		
	}
}