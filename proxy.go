package transports

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
)

type Proxy struct {
	Port      int
	Transport interface{}
}

func (proxy *Proxy) Listen() {
	fmt.Println("Proxy listening on", proxy.Port)

	err := errors.New("no transport specified")

	if proxy.Transport == nil {
		panic(err)
	}

	// transport := proxy.Transport.(FacebookTransport)
	transport := proxy.Transport.(WhatsappTransport)
	transport.Prepare()

	http.HandleFunc("/", transport.Handler)
	http.HandleFunc( "/stats", stats_api.Handler )

	http.ListenAndServe(":8080", nil)

	return
}

func MarshalRequest(request *http.Request) []byte {
	r := Request{
		Method:  request.Method,
		URL:     request.URL.String(),
		Proto:   request.Proto,
		Headers: request.Header,
	}
	output, _ := json.Marshal(r)
	return output
}
