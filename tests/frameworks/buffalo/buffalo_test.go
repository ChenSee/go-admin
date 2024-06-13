package buffalo

import (
	"net/http"
	"testing"

	"github.com/ChenSee/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestBuffalo(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
