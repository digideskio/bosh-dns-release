package handlers

import (
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/miekg/dns"
)

type ArpaHandler struct {
	logger             logger.Logger
	logTag             string
	recursionAvailable bool
}

func NewArpaHandler(logger logger.Logger, recursionAvailable bool) ArpaHandler {
	return ArpaHandler{
		logger:             logger,
		logTag:             "ArpaHandler",
		recursionAvailable: recursionAvailable,
	}
}

func (a ArpaHandler) ServeDNS(resp dns.ResponseWriter, req *dns.Msg) {
	m := &dns.Msg{}

	m.Authoritative = true
	m.RecursionAvailable = a.recursionAvailable

	a.logger.Info(a.logTag, "received a request with %d questions", len(req.Question))

	if len(req.Question) == 0 {
		m.SetRcode(req, dns.RcodeSuccess)
	} else {
		m.SetRcode(req, dns.RcodeServerFailure)
	}

	if err := resp.WriteMsg(m); err != nil {
		a.logger.Error(a.logTag, err.Error())
	}
}
