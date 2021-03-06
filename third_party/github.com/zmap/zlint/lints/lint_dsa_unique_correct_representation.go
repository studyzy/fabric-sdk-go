package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"crypto/dsa"
	"math/big"

	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zcrypto/x509"
	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zlint/util"
)

type dsaUniqueCorrectRepresentation struct{}

func (l *dsaUniqueCorrectRepresentation) Initialize() error {
	return nil
}

func (l *dsaUniqueCorrectRepresentation) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaUniqueCorrectRepresentation) Execute(c *x509.Certificate) *LintResult {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return &LintResult{Status: NA}
	}
	// Verify that 2 ≤ y ≤ p-2.
	two := big.NewInt(2)
	pMinusTwo := big.NewInt(0)
	pMinusTwo.Sub(dsaKey.P, two)
	if two.Cmp(dsaKey.Y) > 0 || dsaKey.Y.Cmp(pMinusTwo) > 0 {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_unique_correct_representation",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Citation:      "BRs: 6.1.6",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &dsaUniqueCorrectRepresentation{},
	})
}
