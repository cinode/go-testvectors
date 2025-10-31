/*
Copyright © 2025 Bartłomiej Święcki (byo)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testvectors

import (
	"embed"
	"encoding/json"
	"io/fs"
	"strings"

	"github.com/cinode/go-common/cutl"
)

//go:embed data/*
var testVectorsData embed.FS

var TestCases = func() []*TestCase {
	var testCases []*TestCase

	err := fs.WalkDir(testVectorsData, "data", func(path string, d fs.DirEntry, err error) error {
		cutl.PanicIfError(err)

		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		data, err := fs.ReadFile(testVectorsData, path)
		cutl.PanicIfError(err)

		testCase := TestCase{}

		err = json.Unmarshal(data, &testCase)
		cutl.PanicIfError(err)

		testCase.Details = strings.Join(testCase.DetailsLines, "\n")

		testCases = append(testCases, &testCase)

		return nil
	})
	cutl.PanicIfError(err)

	return testCases
}()

var TestCasesByName = func() map[string]*TestCase {
	testCasesByName := make(map[string]*TestCase)
	for _, tc := range TestCases {
		testCasesByName[tc.Name] = tc
	}
	return testCasesByName
}()
