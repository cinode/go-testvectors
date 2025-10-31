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
	"testing"

	"github.com/cinode/go-common/picotestify/require"
)

func TestEmbedded(t *testing.T) {
	for _, tc := range TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Log(tc.Details)
		})
	}

	require.NotEmpty(t, TestCases)
	require.Len(t, TestCasesByName, len(TestCases))

	for name, tc := range TestCasesByName {
		require.Equal(t, tc.Name, name)
	}
}
