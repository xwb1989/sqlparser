/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sqlparser

// Additional tests to address the GitHub issues for this fork.

import (
	"testing"
)

func TestParsing(t *testing.T) {
	tests := []struct {
		github int
		sql    string
	}{
		//{9, "select 1 as 测试 from dual", nil}, // Broken due to ReadByte()
		{12, "SELECT * FROM AccessToken LIMIT 10 OFFSET 13"},
		{14, "SELECT DATE_SUB(NOW(), INTERVAL 1 MONTH)"},
		{15, "select STRAIGHT_JOIN t1.* FROM t1 INNER JOIN  t2 ON t1.CommonID = t2.CommonID WHERE t1.FilterID = 1"},
		{16, "SELECT a FROM t WHERE FUNC(a) = 1"}, // Doesn't seem broken, need better example
	}

	for _, test := range tests {
		_, err := Parse(test.sql)
		if err != nil {
			t.Errorf("https://github.com/xwb1989/sqlparser/issues/%d:\nParse(%q) err = %s, want nil", test.github, test.sql, err)
		}
	}
}
