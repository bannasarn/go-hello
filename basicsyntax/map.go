package basicsyntax

func Map() {
	countries := map[string]string{}
	countries["TH"] = "Thailand"
	countries["US"] = "United State"

	println(countries["TH"])

	// Check value in map
	country, ok := countries["JP"]
	if !ok {
		println("No value")
	}
	println(country)
}
