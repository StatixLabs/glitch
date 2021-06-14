package file

import "io/ioutil"

// ReadContentsFromFile returns contents of file at filepath
func ReadContentsFromFile(file_path string) (string, error) {
	// read contents of file
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		return "", err
	}
	// return contents of file
	return string(content), nil
}

// WriteContentsToFile writes contents of data to a filepath
func WriteContentsToFile(file_path string, data []byte) error {
	err := ioutil.WriteFile(file_path, data, 0644)
	return err
}
