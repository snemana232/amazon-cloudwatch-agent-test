// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package common

import (
	"fmt"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"go.uber.org/multierr"
)

// StartLogWrite starts go routines to write logs to each of the logs that are monitored by CW Agent according to
// the config provided
func StartSendingMetrics(receivers []string, agentRunDuration time.Duration, dataRate int) error {
	//create wait group so main test thread waits for log writing to finish before stopping agent and collecting data
	var (
		err      error
		multiErr error
	)

	for _, receiver := range receivers {
		go func(receiver string, durationMinute time.Duration) {

			switch receiver {
			case "statsd":
				err = sendStatsdMetrics(dataRate, durationMinute)

			default:
			}

			multiErr = multierr.Append(multiErr, err)
		}(receiver, agentRunDuration)
	}
	return multiErr
}

func sendStatsdMetrics(dataRate int, durationMinute time.Duration) error {
	// https://github.com/cactus/go-statsd-client#example
	statsdClientConfig := &statsd.ClientConfig{
		Address:     ":8125",
		Prefix:      "statsd",
		UseBuffered: true,
		// interval to force flush buffer. full buffers will flush on their own,
		// but for data not frequently sent, a max threshold is useful
		FlushInterval: 300 * time.Millisecond,
	}
	client, err := statsd.NewClientWithConfig(statsdClientConfig)

	if err != nil {
		return err
	}

	defer client.Close()

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	endTimeout := time.After(durationMinute)

	for {
		select {
		case <-ticker.C:
			for time := 0; time < dataRate; time++ {
				go func(time int) {
					client.Inc(fmt.Sprintf("%v", time), int64(time), 1.0)
				}(time)
			}
		case <-endTimeout:
			return nil
		}
	}

}