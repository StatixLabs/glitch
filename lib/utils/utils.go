package utils

import "strings"

// ForceKeysToLowercase convert the keys in a map to lowercase
func ForceKeysToLowercase(currentMap map[string]interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	for k, v := range currentMap {
		newMap[strings.ToLower(k)] = v
	}
	return newMap
}
