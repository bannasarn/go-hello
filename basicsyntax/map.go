package basicsyntax

func Map() {
	countries := map[string]string{
		"CN": "China",
		"EN": "England",
	}
	countries["TH"] = "Thailand"
	countries["US"] = "United State"

	// println(countries["TH"])

	// delete(countries, "CN")

	// Check value in map
	country, ok := countries["JP"]
	if !ok {
		println("No value")
	}
	println(country)
}
