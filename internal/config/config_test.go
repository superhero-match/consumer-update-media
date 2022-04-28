/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_New(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// Elasticsearch
	assert.Equal(t, "localhost", cfg.ES.Host, "The ES host should be localhost.")
	assert.Equal(t, "9200", cfg.ES.Port, "The port should be 9200.")
	assert.Equal(t, "superheromatch", cfg.ES.Cluster, "The cluster should be superheromatch.")
	assert.Equal(t, "superhero", cfg.ES.Index, "The index should be superhero.")

	// Cache
	assert.Equal(t, "localhost", cfg.Cache.Address, "The address should be localhost.")
	assert.Equal(t, ":6379", cfg.Cache.Port, "The port should be :6379.")
	assert.Equal(t, "Awesome85**", cfg.Cache.Password, "The password should be Awesome85**.")
	assert.Equal(t, 0, cfg.Cache.DB, "The db should be 0.")
	assert.Equal(t, 25, cfg.Cache.PoolSize, "The pool size should be 25.")
	assert.Equal(t, 10, cfg.Cache.MinimumIdleConnections, "The minimum idle connections should be 10.")
	assert.Equal(t, 1, cfg.Cache.MaximumRetries, "The maximum retries should be 1.")
	assert.Equal(t, "suggestion.%s", cfg.Cache.SuggestionKeyFormat, "The suggestion key format should be suggestion.%s.")

	// DB
	assert.Equal(t, "localhost", cfg.DB.Host, "The DB host should be localhost.")
	assert.Equal(t, 3306, cfg.DB.Port, "The DB port should be 3306.")
	assert.Equal(t, "dev", cfg.DB.User, "The DB user should be dev.")
	assert.Equal(t, "Awesome85**", cfg.DB.Password, "The DB password should be Awesome85**.")
	assert.Equal(t, "municipality", cfg.DB.Name, "The DB name should be municipality.")

	// Consumer
	assert.Equal(t, "localhost:9092", cfg.Consumer.Brokers[0], "The brokers should be localhost:9092.")
	assert.Equal(t, "update.municipality.profilepicture", cfg.Consumer.Topic, "The topic should be update.municipality.profilepicture.")
	assert.Equal(t, "consumer.update.media.group", cfg.Consumer.GroupID, "The group id should be consumer.update.media.group.")
}
