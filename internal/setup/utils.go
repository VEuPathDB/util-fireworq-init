package setup

func bail(err error) {
	if err != nil {
		panic(err)
	}
}

func prefixUrl(url string) string {
	if url[0:4] != "http" {
		return "http://" + url
	}

	return url
}