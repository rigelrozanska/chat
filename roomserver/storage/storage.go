// Copyright 2020 The Matrix.org Foundation C.I.C.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !wasm
// +build !wasm

package storage

import (
	"fmt"

	"freemasonry.cc/chat/internal/caching"
	"freemasonry.cc/chat/roomserver/storage/postgres"
	"freemasonry.cc/chat/roomserver/storage/sqlite3"
	"freemasonry.cc/chat/setup/base"
	"freemasonry.cc/chat/setup/config"
)

// Open opens a database connection.
func Open(base *base.BaseDendrite, dbProperties *config.DatabaseOptions, cache caching.RoomServerCaches) (Database, error) {
	switch {
	case dbProperties.ConnectionString.IsSQLite():
		return sqlite3.Open(base, dbProperties, cache)
	case dbProperties.ConnectionString.IsPostgres():
		return postgres.Open(base, dbProperties, cache)
	default:
		return nil, fmt.Errorf("unexpected database type")
	}
}
