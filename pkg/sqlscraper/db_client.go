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
	"reflect"

	// register DB drivers
	_ "github.com/SAP/go-hdb/driver"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	_ "github.com/snowflakedb/gosnowflake"
	"go.uber.org/zap"
)

type DBClient interface {
	MetricRows(ctx context.Context) ([]MetricRow, error)
}

type DBSQLClient struct {
	DB     *sql.DB
	Logger *zap.Logger
	Sql    string
}

func NewDbClient(db *sql.DB, sql string, logger *zap.Logger) DBClient {
	return DBSQLClient{
		DB:     db,
		Sql:    sql,
		Logger: logger,
	}
}

type MetricRow map[string]string

func (cl DBSQLClient) MetricRows(ctx context.Context) ([]MetricRow, error) {
	sqlRows, err := cl.DB.QueryContext(ctx, cl.Sql)
	if err != nil {
		return nil, err
	}
	var out []MetricRow
	row := ReusableRow{
		Attrs: map[string]func() string{},
	}
	types, err := sqlRows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	for _, sqlType := range types {
		colName := sqlType.Name()
		var v interface{}
		row.Attrs[colName] = func() string {
			format := "%v"
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				// The Postgres driver returns a []uint8 (a string) for decimal and numeric types,
				// which we want to render as strings. e.g. "4.1" instead of "[52, 46, 49]".
				// Other slice types get the same treatment.
				format = "%s"
			}
			return fmt.Sprintf(format, v)
		}
		row.ScanDest = append(row.ScanDest, &v)
	}
	for sqlRows.Next() {
		err = sqlRows.Scan(row.ScanDest...)
		if err != nil {
			return nil, err
		}
		out = append(out, row.ToMetricRow())
	}
	return out, nil
}

type ReusableRow struct {
	Attrs    map[string]func() string
	ScanDest []interface{}
}

func (row ReusableRow) ToMetricRow() MetricRow {
	out := MetricRow{}
	for k, f := range row.Attrs {
		out[k] = f()
	}
	return out
}
