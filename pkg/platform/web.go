package platform

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	eel2 "github.com/jfbramlett/nwp-platform-go/pkg/platform/eelserver"

	"github.com/gorilla/mux"
	"github.com/jfbramlett/go-aop/pkg/aop"
	"github.com/jfbramlett/go-aop/pkg/web"
	"github.com/jfbramlett/nwp-platform-go/pkg/platform/protocols/dda10"
)

var (
	homeTpl *template.Template
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

func NewWebRunner(webRoot string) Runner {
	return &webRunner{webRoot: webRoot, dda10Handler: dda10.NewDDA10Handler(eel2.NewEELHandler())}
}

type webRunner struct {
	webRoot      string
	dda10Handler *dda10.DDA10Handler
}

func (wr *webRunner) Run() {
	base := template.Must(template.ParseFiles(path.Join(wr.webRoot, "web", "index.html")))
	homeTpl = template.Must(template.Must(base.Clone()).ParseFiles(path.Join(wr.webRoot, "web", "templates", "home.html")))

	// set up HTTP sequence
	router := mux.NewRouter()
	router.HandleFunc("/", wr.indexHandler)
	router.HandleFunc("/dda10/accountlist", wr.dda10AccountListHandler)

	loggingMiddleware := &web.LoggingMiddleware{}
	spanMiddleware := &web.SpanMiddleware{}
	traceMiddleware := &web.TracingMiddleware{}
	router.Use(traceMiddleware.Middleware)
	router.Use(loggingMiddleware.Middleware)
	router.Use(spanMiddleware.Middleware)

	fmt.Println("Service listening on port 8085")
	fmt.Println(http.ListenAndServe(":8085", router))
}

func (wr *webRunner) indexHandler(w http.ResponseWriter, r *http.Request) {
	webResp := WebResponse{Msg: "Successful call", Type: "Sample",
		Resources: []Resource{{Name: "Resource1", URL: "resource1.html"}, {Name: "Resource2", URL: "resource2.html"}}}
	_ = homeTpl.Execute(w, webResp)
}

func (wr *webRunner) dda10AccountListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := aop.Before(r.Context())
	defer func() { aop.After(ctx, err) }()

	if r.Method == http.MethodGet {
		resp, err := wr.dda10Handler.GetAccountList(ctx, dda10.NewDDA10AccountListRequest("23", "1111"))
		if err == nil {
			writeResponse(resp, http.StatusOK, w)
		} else {
			writeResponse(map[string]string{"error": err.Error()}, http.StatusBadRequest, w)
		}
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
