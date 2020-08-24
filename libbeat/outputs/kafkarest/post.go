// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package kafkarest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func sendToDest(url string, topic string, kafkaRecords []interface{}) {

	kafkaUrl := "http://" + url +"/topics/" + topic
	fmt.Println(kafkaUrl)

	records := make(map[string]interface{})
	records["records"] = kafkaRecords

	recordsData, err := json.Marshal(records)
    if err != nil {
        fmt.Println(err)
        return 
    }

	//fmt.Println(string(recordsData))
	//fmt.Printf("No of records to be sent %d\n", len(kafkaRecords))
	req, err := http.NewRequest("POST", kafkaUrl, bytes.NewBuffer(recordsData))
    if err != nil {
		fmt.Println(err)
        return
    }

	req.Header.Set("Content-Type", "application/vnd.kafka.json.v2+json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiYWRtaW4vYWRtaW4iLCJpc3MiOiJsb2dhcmNoaXZhbCJ9.Aqhl-amaKaKDoXDc0-8TN4hhI7FFkLa76GwDMBTmR8s")

    client := &http.Client{Timeout: 30 * time.Second}

    res, err := client.Do(req)
    if err != nil {
    	fmt.Println(err)
		fmt.Println(kafkaUrl)
        return
    }
    defer res.Body.Close()
    //fmt.Println(res.StatusCode)
    if res.StatusCode == 200 {
    	//fmt.Printf("Successfully sent records to Kafka\n")
    } else {
		fmt.Println(kafkaUrl)
		fmt.Println(string(recordsData))
    	fmt.Println("Failed to send Kafka records", res.Status)
    }
}
