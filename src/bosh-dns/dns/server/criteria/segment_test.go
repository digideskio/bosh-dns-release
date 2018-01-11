package criteria_test

import (
	"bosh-dns/dns/server/criteria"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewSegment", func() {
	It("creates segment from long-form fqdn", func() {
		s, err := criteria.NewSegment("q-s0.ig.net.depl.bosh", []string{"bosh"})
		Expect(err).NotTo(HaveOccurred())
		Expect(s).To(Equal(criteria.Segment{
			Query:      "q-s0",
			Group:      "",
			Instance:   "ig",
			Network:    "net",
			Deployment: "depl",
			Domain:     "bosh",
		}))
	})

	It("creates segment from short-form fqdn", func() {
		s, err := criteria.NewSegment("q-short.q-group.bosh", []string{"bosh"})
		Expect(err).NotTo(HaveOccurred())
		Expect(s).To(Equal(criteria.Segment{
			Query:      "q-short",
			Group:      "q-group",
			Instance:   "",
			Network:    "",
			Deployment: "",
			Domain:     "bosh",
		}))
	})

	It("errors when the fqdn cannot be split into a query and group segment", func() {
		_, err := criteria.NewSegment("garbage", []string{"bosh"})
		Expect(err).To(MatchError("domain is malformed"))
	})

	It("errors when tld is not in the list of domains", func() {
		_, err := criteria.NewSegment("garbage.fire.potato.bosh", []string{"bosh"})
		Expect(err).To(MatchError("bad group segment query nad 3 values"))
	})

	// TODO: q-s0. and q-s0.bosh should error
})
