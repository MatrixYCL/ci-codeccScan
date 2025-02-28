// (c) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rules

import (
	"go/ast"

	gas "github.com/GoASTScanner/gas/core"
)

type UsesWeakCryptography struct {
	gas.MetaData
	blacklist map[string][]string
}

func (r *UsesWeakCryptography) Match(n ast.Node, c *gas.Context) (*gas.Issue, error) {

	for pkg, funcs := range r.blacklist {
		if _, matched := gas.MatchCallByPackage(n, c, pkg, funcs...); matched {
			return gas.NewIssue(c, n, r.What, r.Severity, r.Confidence), nil
		}
	}
	return nil, nil
}

// Uses des.* md5.* or rc4.*
func NewUsesWeakCryptography(conf map[string]interface{}) (gas.Rule, []ast.Node) {
	calls := make(map[string][]string)
	calls["crypto/des"] = []string{"NewCipher", "NewTripleDESCipher"}
	calls["crypto/md5"] = []string{"New", "Sum"}
	calls["crypto/rc4"] = []string{"NewCipher"}
	rule := &UsesWeakCryptography{
		blacklist: calls,
		MetaData: gas.MetaData{
			Severity:   gas.Medium,
			Confidence: gas.High,
			What:       "gas/crypto->使用了弱的加密方法",
			//What:     "Use of weak cryptographic primitive",
		},
	}
	return rule, []ast.Node{(*ast.CallExpr)(nil)}
}
