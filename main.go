package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GeolocationHTTP", geolocationHTTP)
}

// geolocationHTTP is an HTTP Cloud Function with a request parameter.
func geolocationHTTP(w http.ResponseWriter, r *http.Request) {
	var dout struct {
		Country     string `json:"Country,omitempty"`
		Region      string `json:"Region,omitempty"`
		City        string `json:"City,omitempty"`
		CityLatLong string `json:"CityLatLong,omitempty"`
		UserIP      string `json:"User-IP,omitempty"`
	}
	var din struct {
		XAppengineCountry     []string `json:"X-Appengine-Country,omitempty"`
		XAppengineRegion      []string `json:"X-Appengine-Region,omitempty"`
		XAppengineCity        []string `json:"X-Appengine-City,omitempty"`
		XAppengineCityLatLong []string `json:"X-Appengine-CityLatLong,omitempty"`
		XAppengineUserIP      []string `json:"X-Appengine-User-IP,omitempty"`
	}
	header, err := json.Marshal(r.Header)
	if err != nil {
		return
	}
	if err := json.Unmarshal([]byte(header), &din); err != nil {
		return
	}
	if len(din.XAppengineCountry) > 0 {
		dout.Country = din.XAppengineCountry[0]
	}
	if len(din.XAppengineRegion) > 0 {
		dout.Region = din.XAppengineRegion[0]
	}
	if len(din.XAppengineCity) > 0 {
		dout.City = din.XAppengineCity[0]
	}
	if len(din.XAppengineCityLatLong) > 0 {
		dout.CityLatLong = din.XAppengineCityLatLong[0]
	}
	if len(din.XAppengineUserIP) > 0 {
		dout.UserIP = din.XAppengineUserIP[0]
	}
	data, err := json.Marshal(dout)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(data))
}
