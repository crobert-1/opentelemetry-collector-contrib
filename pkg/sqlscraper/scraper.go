// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sqlscraper"

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type Scraper struct {
	id                 config.ComponentID
	query              Query
	scrapeCfg          scraperhelper.ScraperControllerSettings
	startTime          pcommon.Timestamp
	clientProviderFunc ClientProviderFunc
	dbProviderFunc     DBProviderFunc
	logger             *zap.Logger
	client             DBClient
	db                 *sql.DB
}

var _ scraperhelper.Scraper = (*Scraper)(nil)

func (s Scraper) ID() config.ComponentID {
	return s.id
}

func (s *Scraper) Start(context.Context, component.Host) error {
	var err error
	s.db, err = s.dbProviderFunc()
	if err != nil {
		return fmt.Errorf("failed to open DB connection: %w", err)
	}
	s.client = s.clientProviderFunc(s.db, s.query.SQL, s.logger)
	s.startTime = pcommon.NewTimestampFromTime(time.Now())

	return nil
}

func (s Scraper) Scrape(ctx context.Context) (pmetric.Metrics, error) {
	out := pmetric.NewMetrics()
	rows, err := s.client.MetricRows(ctx)
	ts := pcommon.NewTimestampFromTime(time.Now())
	if err != nil {
		return out, fmt.Errorf("Scraper: %w", err)
	}
	rms := out.ResourceMetrics()
	rm := rms.AppendEmpty()
	sms := rm.ScopeMetrics()
	sm := sms.AppendEmpty()
	ms := sm.Metrics()
	var errs error
	for _, metricCfg := range s.query.Metrics {
		for i, row := range rows {
			if err = rowToMetric(row, metricCfg, ms.AppendEmpty(), s.startTime, ts, s.scrapeCfg); err != nil {
				err = fmt.Errorf("row %d: %w", i, err)
				errs = multierr.Append(errs, err)
			}
		}
	}
	if errs != nil {
		errs = fmt.Errorf("Scraper.Scrape row conversion errors: %w", errs)
	}
	return out, errs
}

func (s Scraper) Shutdown(ctx context.Context) error {
	return s.db.Close()
}
