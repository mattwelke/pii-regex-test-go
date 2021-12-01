package main

// https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// redactResultForDisplay is used to display failed tests nicer (string instead
// of bytes).
type redactResultForDisplay struct {
	numRedacted int64
	redacted    string
}

func redactResultForDisplayFromRedactResult(result redactResult) redactResultForDisplay {
	return redactResultForDisplay{
		numRedacted: result.numRedacted,
		redacted:    string(result.redacted),
	}
}
