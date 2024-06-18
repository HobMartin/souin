package context

import (
	"testing"

	"github.com/HobMartin/souin/configurationtypes"
	"github.com/HobMartin/souin/plugins/souin/configuration"
	"go.uber.org/zap"
)

func Test_ModeContext_SetupContext(t *testing.T) {
	dc := configurationtypes.DefaultCache{}
	c := configuration.Configuration{
		DefaultCache: &dc,
	}
	c.SetLogger(zap.NewNop())
	ctx := ModeContext{}

	ctx.SetupContext(&c)
	if ctx.Bypass_request || ctx.Bypass_response || !ctx.Strict {
		t.Error("The context must be strict and must not bypass either response or request.")
	}

	c.DefaultCache.Mode = "bypass"
	ctx = ModeContext{}
	ctx.SetupContext(&c)
	if !ctx.Bypass_request || !ctx.Bypass_response || ctx.Strict {
		t.Error("The context must bypass either response and request.")
	}

	c.DefaultCache.Mode = "bypass_request"
	ctx = ModeContext{}
	ctx.SetupContext(&c)
	if !ctx.Bypass_request || ctx.Bypass_response || ctx.Strict {
		t.Error("The context must bypass request only.")
	}

	c.DefaultCache.Mode = "bypass_response"
	ctx = ModeContext{}
	ctx.SetupContext(&c)
	if ctx.Bypass_request || !ctx.Bypass_response || ctx.Strict {
		t.Error("The context must bypass response only.")
	}

	c.DefaultCache.Mode = "strict"
	ctx = ModeContext{}
	ctx.SetupContext(&c)
	if ctx.Bypass_request || ctx.Bypass_response || !ctx.Strict {
		t.Error("The context must be strict.")
	}

	c.DefaultCache.Mode = "default_value"
	ctx = ModeContext{}
	ctx.SetupContext(&c)
	if ctx.Bypass_request || ctx.Bypass_response || !ctx.Strict {
		t.Error("The context must be strict.")
	}
}
