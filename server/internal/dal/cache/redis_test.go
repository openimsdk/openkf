// Copyright © 2023 OpenIM open source community. All rights reserved.
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

package cache

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestSub(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "openkf",
		DB:       0,
	})

	ctx := context.Background()
	pubsub := client.PSubscribe(ctx, "__keyevent@0__:expired")

	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			fmt.Println("Events:", msg.Payload)
		}
	}()

	err := client.Set(ctx, "mykey", "myvalue", 5*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	err = pubsub.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}
}
