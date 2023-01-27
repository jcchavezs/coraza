// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"github.com/corazawaf/coraza/v3/internal/corazawaf"
	"github.com/corazawaf/coraza/v3/macro"
	"github.com/corazawaf/coraza/v3/rules"
)

type appendFn struct {
	data macro.Macro
}

func (a *appendFn) Init(_ rules.RuleMetadata, data string) error {
	macro, err := macro.NewMacro(data)
	if err != nil {
		return err
	}
	a.data = macro
	return nil
}

func (a *appendFn) Evaluate(_ rules.RuleMetadata, txS rules.TransactionState) {
	tx := txS.(*corazawaf.Transaction)

	if !tx.ContentInjection() {
		tx.DebugLogger().Debug("append rejected because of ContentInjection")
		return
	}

	data := a.data.Expand(tx)
	if len(data) > 0 {
		tx.AppendInResponseBody([]byte(data))
	}
}

func (a *appendFn) Type() rules.ActionType {
	return rules.ActionTypeNondisruptive
}

func append2() rules.Action {
	return &appendFn{}
}

var (
	_ rules.Action      = &appendFn{}
	_ ruleActionWrapper = append2
)
