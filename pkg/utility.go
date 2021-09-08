package pkg

import "os"

// Check wether an absolute path exists
func DoesPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { 
		return true, nil 
	}
    
	if os.IsNotExist(err) { 
		return false, nil 
	}
    
	return false, err
}

// Create a certain directory if it doesn't exist
func CheckOrCreateDir(path string) error {
	pathExisting, _ := DoesPathExist(path)
	if pathExisting {
		return nil
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Get the index of a certain item in a string slice
func IndexOf(target string, data []string) (int) {
    for k, v := range data {
        if target == v {
            return k
        }
    }
    return -1
}

// Check wether a certain item exists in a string slice
func StringSliceContains(target string, data []string) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}

	return false
}