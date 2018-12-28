package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(v interface{}) (int, error) {
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

func SprintPlain(v interface{}) string {
	d, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(d)
}
