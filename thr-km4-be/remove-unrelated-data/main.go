package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	dataMap1 := map[string]any{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	key := "address"
	newMap1 := removeUnrelated(dataMap1, key)
	fmt.Printf("%+v\n", newMap1)

	dataMap2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	key2 := "jump"
	newMap2 := removeUnrelated(dataMap2, key2)
	fmt.Printf("%+v\n", newMap2)

	dataMap3 := map[string]interface{}{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "berenang",
		"empat": "yon",
		"lima":  "go",
	}
	key3 := "tiga"
	newMap3 := removeUnrelated(dataMap3, key3)
	fmt.Printf("%+v\n", newMap3)

}
