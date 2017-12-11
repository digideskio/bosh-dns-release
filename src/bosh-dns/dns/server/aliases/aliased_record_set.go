package aliases

import "bosh-dns/dns/server/records"

//go:generate counterfeiter . RecordSet

type RecordSet interface {
	Resolve(string) ([]string, error)
	ResolveFullRecord(string) ([]records.Record, error)
	Domains() []string
	Subscribe() <-chan bool
}

type AliasedRecordSet struct {
	recordSet RecordSet
	config    Config
}

func NewAliasedRecordSet(recordSet RecordSet, config Config) *AliasedRecordSet {
	return &AliasedRecordSet{
		recordSet: recordSet,
		config:    config,
	}
}

func (a *AliasedRecordSet) Resolve(domain string) ([]string, error) {
	records, err := a.ResolveFullRecord(domain)
	if err != nil {
		return []string{}, err
	}
	var s []string
	for _, rec := range records {
		s = append(s, rec.IP)
	}
	return s, nil
}

func (a *AliasedRecordSet) ResolveFullRecord(domain string) ([]records.Record, error) {
	resolutions := a.config.Resolutions(domain)
	if len(resolutions) > 0 {
		var err error
		recs := []records.Record{}

		for _, resolution := range resolutions {
			var rec []records.Record
			rec, err = a.recordSet.ResolveFullRecord(resolution)
			recs = append(recs, rec...)
		}

		if len(recs) == 0 && err != nil {
			return nil, err
		}
		return recs, nil
	}

	return a.recordSet.ResolveFullRecord(domain)
}

func (a *AliasedRecordSet) Subscribe() <-chan bool {
	return a.recordSet.Subscribe()
}

func (a *AliasedRecordSet) Domains() []string {
	return append(a.recordSet.Domains(), a.config.AliasHosts()...)
}
