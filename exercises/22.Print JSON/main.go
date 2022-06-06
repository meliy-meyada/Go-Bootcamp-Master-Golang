package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	j := `{
						"Items": [{
							"Item": {
								"name": "Teddy Bear"
							}
						}]
					}`
	var obj map[string]interface{}
	json.Unmarshal([]byte(j), &obj) 

	fmt.Println(string(j))
}