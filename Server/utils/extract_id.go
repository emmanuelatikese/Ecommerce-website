package response

import "strings"

func ExtractPublicID(url string) string {
	segments := strings.Split(url, "/")
	if len(segments) < 2 {
		return ""
	}
	publicIDWithExt := segments[len(segments)-1]
	publicID := strings.TrimSuffix(publicIDWithExt, ".jpg")
	return publicID
}