package shared

func Must(e error) {
	if e != nil {
		panic(e)
	}
}

func Try(e error) {
	if e != nil {
		println(e.Error())
	}
}
