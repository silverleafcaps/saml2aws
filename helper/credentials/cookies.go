package credentials

func LookupCookies(cookiesJson *string, url string) error {
	_, secret, err := CurrentHelper.Get(url + "/cookie")
	if err != nil {
		return err
	}

	*cookiesJson = secret
	return nil
}

func SaveCookies(cookiesJson string, url string) error {
	creds := &Credentials{
		ServerURL: url + "/cookie",
		Secret:    cookiesJson,
	}

	return CurrentHelper.Add(creds)
}