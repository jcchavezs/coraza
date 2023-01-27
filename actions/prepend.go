// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"github.com/corazawaf/coraza/v3/internal/corazawaf"
	"github.com/corazawaf/coraza/v3/macro"
	"github.com/corazawaf/coraza/v3/rules"
)

type prependFn struct {
	data macro.Macro
}

func (a *prependFn) Init(r rules.RuleMetadata, data string) error {
	m, err := macro.NewMacro(data)
	if err != nil {
		return err
	}
	a.data = m
	return nil
}

func (a *prependFn) Evaluate(r rules.RuleMetadata, txS rules.TransactionState) {
	tx := txS.(*corazawaf.Transaction)

	if !tx.ContentInjection() {
		tx.DebugLogger().Debug("append rejected because of ContentInjection")
		return
	}

	data := a.data.Expand(tx)

	tx.PrependInResponseBody([]byte(data))
}

func (a *prependFn) Type() rules.ActionType {
	return rules.ActionTypeNondisruptive
}

func prepend() rules.Action {
	return &prependFn{}
}

var (
	_ rules.Action      = &prependFn{}
	_ ruleActionWrapper = prepend
)
