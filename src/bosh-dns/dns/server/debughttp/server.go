package debughttp

import (
	"fmt"
	"net/http"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Server interface {
	Serve()
}

type recordSet interface {
	Resolve(string) ([]string, error)
	ResolveFullRecord(string) ([]Record, error)
	Domains() []string
	Subscribe() <-chan bool
}

type concreteServer struct {
	logger    boshlog.Logger
	port      int
	recordset recordSet
}

const logTag = "debugHTTP"

func NewServer(logger boshlog.Logger, port int, recordset recordSet) Server {
	return &concreteServer{
		logger:    logger,
		port:      port,
		recordset: recordset,
	}
}

func (c *concreteServer) instancesHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.FormValue("domain")
	if len(domain) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := c.recordset.ResolveFullRecord(domain)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (c *concreteServer) Serve() {
	http.HandleFunc("/instances", c.instancesHandler)

	server := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", c.port),
	}
	server.SetKeepAlivesEnabled(false)

	serveErr := server.ListenAndServe()
	c.logger.Error(logTag, "server ending with %s", serveErr)
}
