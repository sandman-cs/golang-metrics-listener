package main

import (
	"encoding/json"
)

func getRouteKey(payload []byte) string {

	var r metricPayload
	var route = "p.bad"

	//initialize Route
	r.Bu = "unknown"
	r.Env = "unknown"
	r.HostEnv = "unknown"
	r.App = "unknown"

	err = json.Unmarshal(payload, &r)
	if err == nil {
		if len(r.Route) > 1 {
			route = r.Route
		} else {
			route = r.Bu + "." + r.Env + "." + r.HostEnv + "." + r.App
		}
	}
	return (route)
}
