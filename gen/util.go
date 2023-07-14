package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const env = "DM_PATH"

func DMPath() (string, bool) {
	path := os.Getenv(env)
	return path, path != ""
}

func IsDMAvailable() bool {
	_, ok := DMPath()
	return ok
}

func ReadDMFile[T any](result T, path ...string) error {
	dm, ok := DMPath()
	if !ok {
		return fmt.Errorf("dm not available: the %v environment variable is not set", env)
	}

	fullPath := make([]string, 0, len(path)+1)
	fullPath = append(fullPath, dm)
	fullPath = append(fullPath, path...)

	jsonFile := filepath.Join(fullPath...)
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	return nil
}
