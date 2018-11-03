package pretty

import (
	"encoding/json"
	"fmt"
)

func Sprint(v interface{}) string {
	d, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(d)
}
