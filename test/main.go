package main

import (
	"fmt"

	JsonParser "github.com/ZombiiTheCoder/JsonParser/jlib"
)

func main (){
	m := JsonParser.MapJson(`
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
	// JsonParser.Access changes type any to map[string]any to allow you to look through the map and see values
	// JsonParser.Access takes a map[string]any for the tree and a string for the key 
	body := JsonParser.Access(m, "body")

	fmt.Println(JsonParser.Access(body, "d"))
}