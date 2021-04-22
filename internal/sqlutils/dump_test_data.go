package sqlutils

import "io/ioutil"

// DumpTestData stores data to specified file
func DumpTestData(filename, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
