package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(v interface{}) string {
	return fmt.Println(Sprint(v))
}

func Sprint(v interface{}) string {
	d, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(d)
}
