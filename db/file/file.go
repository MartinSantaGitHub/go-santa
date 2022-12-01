package file

import (
	"errors"
	"log"
	"os"
	"strings"
)

type DbFile struct {
	FilePath string
}

// region "Connection"

/* Connect creates the file */
func (db *DbFile) Connect() error {
	var err error
	var file *os.File

	if !db.IsConnection() {
		file, err = os.Create(db.FilePath)
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	if err != nil {
		return errors.New("an error has occurred trying to create the file")
	}

	return nil
}

/* IsConnection verifies that the file exist */
func (db *DbFile) IsConnection() bool {
	_, err := os.Stat(db.FilePath)

	if os.IsNotExist(err) {
		log.Println("file does not exist: creating new file")

		return false
	}

	if err != nil {
		log.Println(err.Error())

		return false
	}

	return true
}

// endregion

// region "Greetings"

/* SaveName saves a name to the file */
func (db *DbFile) SaveName(name string) (bool, error) {
	isDbEntry, err := db.isEntry(name)

	if err != nil {
		return false, err
	}

	if !isDbEntry {
		file, err := os.OpenFile(db.FilePath, os.O_WRONLY|os.O_APPEND, 0644)

		if err != nil {
			return false, err
		}

		_, err = file.WriteString(name + "\n")

		return false, err
	}

	return true, nil
}

/* GetNames gets the names saved in the file */
func (db *DbFile) GetNames() ([]string, error) {
	var names []string

	s, err := db.getFileContent()

	if err != nil {
		return names, err
	}

	s = strings.TrimSpace(s)
	names = strings.Split(s, "\n")

	if len(names) == 1 && names[0] == "" {
		names = nil
	}

	return names, nil
}

// endregion

// region "Helpers"

func (db *DbFile) getFileContent() (string, error) {
	b, err := os.ReadFile(db.FilePath)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (db *DbFile) isEntry(value string) (bool, error) {
	s, err := db.getFileContent()

	if err != nil {
		return false, err
	}

	return strings.Contains(s, value), nil
}

// endregion
