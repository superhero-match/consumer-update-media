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

package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"

	"github.com/superhero-match/consumer-update-media/cmd/consumer/reader"
	"github.com/superhero-match/consumer-update-media/internal/cache"
	"github.com/superhero-match/consumer-update-media/internal/config"
	"github.com/superhero-match/consumer-update-media/internal/consumer"
	"github.com/superhero-match/consumer-update-media/internal/db"
	"github.com/superhero-match/consumer-update-media/internal/es"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetURL(
			fmt.Sprintf(
				"http://%s:%s",
				cfg.ES.Host,
				cfg.ES.Port,
			),
		),
	)
	if err != nil {
		panic(err)
	}

	e := es.New(client, cfg.ES.Index)

	dtbs, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%d)/%s",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		),
	)
	if err != nil {
		panic(err)
	}

	dbs, err := db.New(dtbs)
	if err != nil {
		panic(err)
	}

	c := consumer.New(cfg)

	rc := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s%s", cfg.Cache.Address, cfg.Cache.Port),
		Password:     cfg.Cache.Password,
		DB:           cfg.Cache.DB,
		PoolSize:     cfg.Cache.PoolSize,
		MinIdleConns: cfg.Cache.MinimumIdleConnections,
		MaxRetries:   cfg.Cache.MaximumRetries,
	})

	_, err = rc.Ping().Result()
	if err != nil {
		panic(err)
	}

	ch := cache.New(rc)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	r := reader.New(e, dbs, c, ch, logger, cfg.Cache.SuggestionKeyFormat)

	err = r.Read()
	if err != nil {
		panic(err)
	}
}
