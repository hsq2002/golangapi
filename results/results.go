package results

var resultsStore = make(map[string][]byte)

func StoreResult(streamID string, results []byte) {
	resultsStore[streamID] = result
}

func ReterieveResults(streamID string) []byte {
	return resultsStore[streamId]
}


