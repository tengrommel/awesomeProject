package configSaver

import (
	"encoding/json"
	"log"
	"os"
)

/*
Sadly, this refactor suffers from similar issues as the previous one,
not least of which is that it has the potential to be a considerable amount of changes.
 */

type Config struct {
	ID int64
	Name string
}

type ConfigSaver struct {
	FileWriter func(filename string, data []byte, perm os.FileMode) error
}

func (c ConfigSaver) Save(filename string, cfg *Config) error {
	// convert to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	// save file
	err = c.FileWriter(filename, data, 0666)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
		return err
	}
	return nil
}
