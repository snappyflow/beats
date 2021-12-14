package elasticsearch

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/snappyflow/beats/v7/libbeat/beat"
	"github.com/snappyflow/beats/v7/libbeat/logp"
	"github.com/snappyflow/beats/v7/libbeat/outputs/codec"
)

func getEventMessage(log *logp.Logger, index string, data *beat.Event) ([]byte, error) {
	encoder, err := codec.CreateEncoder(beat.Info{}, codec.Config{})
	if err != nil {
		log.Debugf("Error initializing encoder: %v", err)
	}
	serializedEvent, err := encoder.Encode(index, data)
	if err != nil {
		if log.IsDebug() {
			log.Debugf("failed event: %v", err)
		}
		return nil, err
	}

	buf := make([]byte, len(serializedEvent))
	copy(buf, serializedEvent)

	return buf, nil
}

func getMetricName(projectName string) string {
	metricName := ""
	for _, letter := range projectName {
		if unicode.IsUpper(letter) {
			// name_list.append("_{}".format(a_letter.lower()))
			metricName += fmt.Sprintf("_%s", strings.ToLower(string(letter)))
		} else if letter == '_' {
			metricName += "__"
		} else {
			metricName += string(letter)
		}
	}
	return metricName
}
