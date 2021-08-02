package debug

import "encoding/json"

func Debug(value interface{}) {
	bytes, err := json.Marshal(value)
	if err != nil {
		println("Failed to debug, value was null or could not be serialized.")
		return
	}
	println(string(bytes))
}
