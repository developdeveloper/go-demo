package keyvalue

func updateDict(dict map[string]string) {
	dict["four"] = "4"
	delete(dict, "one")
}
