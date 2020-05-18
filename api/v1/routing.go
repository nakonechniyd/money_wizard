package apiv1

import "github.com/gorilla/mux"

// AddRouts -
func AddRouts(router *mux.Router) {
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()

	var route *mux.Route

	// calc_years
	route = apiV1Router.Path("/calc_years/{years:[0-9]+}")
	route.Queries("tax_percent", "{tax_percent}")
	route.Queries("passive_percent", "{passive_percent}")
	route.Name("CalcYears")
	route.HandlerFunc(CalcYears)
}
