package internal

import (
	"encoding/json"
	"errors"
	"os"
	"io"
	"passgen/internal/model"
)

func SaveCredential(cred model.Credential, path string) error {
	var creds []model.Credential
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&creds); err != nil && !errors.Is(err, io.EOF) {
			return err
		}
	}
	creds = append(creds, cred)
	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}