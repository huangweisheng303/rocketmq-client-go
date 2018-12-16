/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/core"
	"time"
)

func main() {
	cfg := &rocketmq.ProducerConfig{}
	cfg.GroupID = "testGroup"
	cfg.NameServer = "47.101.55.250:9876"
	producer, err := rocketmq.NewProducer(cfg)
	if err != nil {
		fmt.Println("create Producer failed, error:", err)
		return
	}

	producer.Start()
	defer producer.Shutdown()

	fmt.Printf("Producer: %s started... \n", producer)
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Hello RocketMQ-%d", i)
		result := producer.SendMessageSync(&rocketmq.Message{Topic: "wwf1", Body: msg})
		fmt.Println(fmt.Sprintf("send message: %s result: %s", msg, result))
	}
	time.Sleep(10 * time.Second)
}
