package setup

func bail(err error) {
	if err != nil {
		panic(err)
	}
}
