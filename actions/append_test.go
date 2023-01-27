package actions

import (
	"testing"

	"github.com/corazawaf/coraza/v3/internal/corazawaf"
)

func TestContentInjection(t *testing.T) {
	waf := corazawaf.NewWAF()
	waf.ContentInjection = true

	tx := waf.NewTransaction()
	tx.Variables().TX().Set("a", []string{"hello"})
	tx.Variables().TX().Set("b", []string{"world"})

	t.Run("append", func(t *testing.T) {
		append := appendFn{}
		err := append.Init(nil, "a=%{tx.a}&")
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		append.Evaluate(nil, tx)

		if want, have := "a=hello&", tx.ResponseBodyAppend(); want != string(have) {
			t.Errorf("unexpected response body prepend, want: %q, have: %q", want, have)
		}
	})

	t.Run("prepend", func(t *testing.T) {
		prepend := prependFn{}
		err := prepend.Init(nil, "b=%{tx.b}&")
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		prepend.Evaluate(nil, tx)

		if want, have := "b=world&", tx.ResponseBodyPrepend(); want != string(have) {
			t.Errorf("unexpected response body prepend, want: %q, have: %q", want, have)
		}
	})
}
