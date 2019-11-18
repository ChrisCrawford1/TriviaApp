package helpers

import "bytes"

// BuildQuery will return a well formed url with
// the appropriate query paramters
func BuildQuery(baseURL string, params map[string]string) string {
	var count int
	buf := bytes.Buffer{}
	buf.WriteString(baseURL)

	for i, line := range params {
		count++
		if count == len(params) {
			buf.WriteString("&" + i + "=" + line)
			break
		}
		buf.WriteString(i + "=" + line)
	}

	return buf.String()
}
