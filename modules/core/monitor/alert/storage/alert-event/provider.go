// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alert_event

import (
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/kafka"
	"github.com/erda-project/erda-infra/providers/mysql"

	"github.com/erda-project/erda/modules/core/monitor/alert/alert-apis/db"
)

type define struct{}

type config struct {
	Input kafka.ConsumerConfig `file:"input"`
}

func (d *define) Services() []string {
	return []string{"alert-event-storage"}
}

func (d *define) Dependencies() []string {
	return []string{"kafka", "mysql", "kafka.topic.initializer"}
}

func (d *define) Summary() string {
	return "alert event storage"
}

func (d *define) Description() string {
	return d.Summary()
}

func (d *define) Config() interface{} {
	return &config{}
}

func (d *define) Creator() servicehub.Creator {
	return func() servicehub.Provider {
		return &provider{}
	}
}

type provider struct {
	C     *config
	L     logs.Logger
	kafka kafka.Interface

	alertEventDB *db.AlertEventDB
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.alertEventDB = &db.AlertEventDB{DB: ctx.Service("mysql").(mysql.Interface).DB()}
	p.kafka = ctx.Service("kafka").(kafka.Interface)
	return nil
}

func (p *provider) Start() error {
	err := p.kafka.NewConsumer(&p.C.Input, p.invoke)
	return err
}

func (p *provider) Close() error {
	logrus.Debug("not support close kafka consumer")
	return nil
}

func init() {
	servicehub.RegisterProvider("alert-event-storage", &define{})
}
