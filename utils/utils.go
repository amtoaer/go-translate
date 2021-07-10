package utils

func CheckEmpty(fileToken, defaultToken string) string {
	if fileToken == "" {
		return defaultToken
	}
	return fileToken
}
