// Copyright The OpenTelemetry Authors
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

package googlecloudpubsublitereceiver

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/otlp"
)

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesReceiver(createTraceReceiver),
	)
}

func createTraceReceiver(_ context.Context, param component.ReceiverCreateSettings, cfg config.Receiver, nextConsumer consumer.Traces) (component.TracesReceiver, error) {
	rxrSettings, ok := cfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("configuration did not match expected format")
	}

	sr := &pubsubliteReceiver{
		tracesConsumer:    nextConsumer,
		logger:            param.Logger,
		Config:            *rxrSettings,
		tracesUnmarshaler: otlp.NewProtobufTracesUnmarshaler(),
	}

	return sr, nil
}
