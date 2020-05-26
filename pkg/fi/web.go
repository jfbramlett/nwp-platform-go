package fi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jfbramlett/go-aop/pkg/rest"

	"github.com/jfbramlett/nwp-platform-go/pkg/eelmodel"

	"github.com/gorilla/mux"
	"github.com/jfbramlett/go-aop/pkg/aop"
	"github.com/jfbramlett/go-aop/pkg/web"
)

type Resource struct {
	Name string
	URL  string
}

type WebResponse struct {
	Msg       string
	Error     error
	Type      string
	Resources []Resource
}

type Runner interface {
	Run()
}

func NewFIWebRunner(webRoot string) Runner {
	return &webRunner{webRoot: webRoot}
}

type webRunner struct {
	webRoot string
}

func (wr *webRunner) Run() {
	// set up HTTP sequence
	router := mux.NewRouter()
	router.HandleFunc("/accountlist", wr.accountListHandler)

	loggingMiddleware := &web.LoggingMiddleware{}
	spanMiddleware := &web.SpanMiddleware{}
	traceMiddleware := &web.TracingMiddleware{}
	router.Use(traceMiddleware.Middleware)
	router.Use(loggingMiddleware.Middleware)
	router.Use(spanMiddleware.Middleware)

	rest.InitRestClient()
	rest.AddRequestProxy(rest.NewTraceRequestProxy())
	rest.AddRequestProxy(rest.NewLoggingRequestProxy())

	fmt.Println("Service listening on port 8090")
	fmt.Println(http.ListenAndServe(":8090", router))
}

func (wr *webRunner) accountListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := aop.Before(r.Context())
	defer func() { aop.After(ctx, err) }()

	if r.Method == http.MethodGet {
		resp := eelmodel.NewAccountListResponse([]*eelmodel.AccountResponse{eelmodel.NewAccountResponse("10",
			"123", "CHECKING", 55.7, "CHECK")})

		writeResponse(resp, http.StatusOK, w)

		return
	}

	writeResponse(map[string]string{"error": "GET is only supported method"}, http.StatusBadRequest, w)
}

func writeResponse(data interface{}, status int, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "    ")
	if err := enc.Encode(&data); err != nil {
		return
	}
}
