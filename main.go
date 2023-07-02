package main

import (
	"fmt"
	jsonparser "jsonparser/jlib"
)

func main (){
	m := jsonparser.MapJson(`
	{
		"dependencies": {
		"ejs": "^3.1.6",
		"express": "^4.17.3",
		"mongoose": "^6.2.10",
		"morgan": "^1.10.0",
		"nodemon": "^2.0.15"
		}
	}
	`) // returns Map of type map[string]any from json source file

	// m.body accesses the contents of the json file
	// jsonparser.Access changes type any to map[string]any to allow you to look through the map and see values
	// jsonparser.Access takes a map[string]any for the tree and a string for the key 
	body := jsonparser.Access(m, "body")

	fmt.Println(jsonparser.Access(body, "d"))
}