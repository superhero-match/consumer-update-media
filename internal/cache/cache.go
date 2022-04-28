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

package cache

import (
	"github.com/go-redis/redis"
)

// Cache interface defines cache methods.
type Cache interface {
	DeleteSuperhero(keys []string) error
}

// cache is the Redis client.
type cache struct {
	Redis redis.Cmdable
}

// New creates a client connection to Redis.
func New(rc redis.Cmdable) Cache {
	return &cache{
		Redis: rc,
	}
}
