// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/filter"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestPrometheus(t *testing.T) {
	CONFIG := []byte(`
metrics:
- type: counter
  name: message_foo_counter
  desc: "The total number of foo in message."
  key: foo
  labels:
    foo: bar
`)
	expected := `
<filter **>
  @type prometheus
  @id test_prometheus
  <metric>
    desc The total number of foo in message.
    key foo
    name message_foo_counter
    type counter
    <label>
      foo bar
    </label>
  </metric>
</filter>
`
	parser := &filter.PrometheusConfig{}
	yaml.Unmarshal(CONFIG, parser)
	test := render.NewOutputPluginTest(t, parser)
	test.DiffResult(expected)
}
