package natstesting

import (
	"github.com/formancehq/go-libs/v2/logging"
	. "github.com/formancehq/go-libs/v2/testing/utils"
	. "github.com/onsi/ginkgo/v2"
)

func WithNewNatsServer(logger logging.Logger, fn func(p *Deferred[*NatsServer])) bool {
	return Context("With new nats server", func() {
		ret := NewDeferred[*NatsServer]()
		BeforeEach(func() {
			ret.Reset()
			ret.SetValue(CreateServer(
				GinkgoT(),
				true,
				logger,
			))
		})
		fn(ret)
	})
}
