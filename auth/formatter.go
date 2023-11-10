package auth

type TokenFormatter struct {
	AccessToken string `json:"access_token"`
}

func FormatToken(accessToken string) TokenFormatter {
	formatter := TokenFormatter{
		AccessToken: accessToken,
	}

	return formatter
}
