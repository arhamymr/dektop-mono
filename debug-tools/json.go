package debugs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func PrintPrettyJSON(str string) {
	res, err := PrettyString(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res)
}
