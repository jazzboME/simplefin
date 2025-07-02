package simplefin

import (
	"encoding/json"
	"fmt"
	"os"
)

func SfinFile(fileName string) (Accounts, error) {
	var accts Accounts

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return Accounts{}, fmt.Errorf("failed to open file: %w", err)
	}
	
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&accts); err != nil {
		return Accounts{}, fmt.Errorf("failed to parse file: %w", err)
	}

	return accts, nil
}
