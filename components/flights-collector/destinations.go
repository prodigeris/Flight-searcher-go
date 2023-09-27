package main

type Flight struct {
	FromAirport string
	ToAirport   string
}

func GetVilniusToCodes() []string {
	return []string{
		"WAW",
		"HEL",
		"RIX",
		"ARN",
		"CPH",
		"IST",
		"LTN",
		"FRA",
		"OSL",
		"BER",
		"AMS",
		"BCN",
		"EIN",
		"FCO",
		"STN",
		"VIE",
		"LCY",
		"NCE",
		"TLL",
		"TLV",
		"AGP",
		"BGY",
		"BRU",
		"BVA",
		"CDG",
		"MUC",
		"TRF",
		"ZRH",
		"MXP",
		"SPU",
		"CFU",
		"DTM",
		"DUB",
		"HHN",
		"KEF",
		"KUT",
		"HAM",
		"TRN",
		"HER",
		"PMI",
		"VAR",
		"ATH",
		"BLL",
		"BRE",
		"LCA",
		"MLA",
		"NUE",
		"DBV",
		"TSF",
		"GNB",
		"DXB",
		"TFS",
	}
}

func GetKaunasToCodes() []string {
	return []string{
		"LTN",
		"CPH",
		"STN",
		"ARN",
		"BTS",
		"DUB",
		"SVG",
		"BRI",
		"WMI",
		"AAL",
		"AGP",
		"CRL",
		"GOT",
		"LPL",
		"MAD",
		"CGN",
		"AES",
		"ALC",
		"BGO",
		"BOJ",
		"BRS",
		"EDI",
		"NAP",
		"PFO",
		"RMI",
		"SNN",
		"PMI",
		"RHO",
	}
}

func GetAllDeparturesAndArrivals() ([]Flight, []Flight) {
	kaunasDestinations := GetKaunasToCodes()
	vilniusDestinations := GetVilniusToCodes()
	var departures []Flight
	var arrivals []Flight

	for _, kaunasDestination := range kaunasDestinations {
		departures = append(departures, Flight{FromAirport: "KUN", ToAirport: kaunasDestination})
	}
	for _, vilniusDestination := range vilniusDestinations {
		departures = append(departures, Flight{FromAirport: "VNO", ToAirport: vilniusDestination})
	}

	for _, kaunasDestination := range kaunasDestinations {
		arrivals = append(arrivals, Flight{FromAirport: kaunasDestination, ToAirport: "KUN"})
	}
	for _, vilniusDestination := range vilniusDestinations {
		arrivals = append(arrivals, Flight{FromAirport: vilniusDestination, ToAirport: "VNO"})
	}

	return departures, arrivals
}
