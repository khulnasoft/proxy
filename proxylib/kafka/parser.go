// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package kafka

import (
	"encoding/binary"

	"github.com/sirupsen/logrus"

	khulnasoft "github.com/khulnasoft/proxy/go/khulnasoft/api"
	"github.com/khulnasoft/proxy/proxylib/kafka/kafkalib"
	. "github.com/khulnasoft/proxy/proxylib/proxylib"
)

const (
	parserName = "kafka"
)

// KafkaRuleParser parses protobuf L7 rules to enforcement objects
// May panic
func KafkaRuleParser(rule *khulnasoft.PortNetworkPolicyRule) []L7NetworkPolicyRule {
	l7Rules := rule.GetKafkaRules()
	if l7Rules == nil {
		return nil
	}

	allowRules := l7Rules.GetKafkaRules()
	rules := make([]L7NetworkPolicyRule, 0, len(allowRules))
	for _, r := range allowRules {
		rules = append(rules, kafkalib.NewRule(r.ApiVersion, r.ApiKeys, r.ClientId, r.Topic))
	}
	return rules
}

type KafkaParserFactory struct{}

var kafkaParserFactory *KafkaParserFactory

func init() {
	logrus.Info("init(): Registering kafkaParserFactory")
	RegisterParserFactory(parserName, kafkaParserFactory)
	RegisterL7RuleParser(parserName, KafkaRuleParser)
}

type KafkaParser struct {
	connection *Connection
}

func (pf *KafkaParserFactory) Create(connection *Connection) interface{} {
	p := KafkaParser{connection: connection}
	return &p
}

func (p *KafkaParser) OnData(reply bool, reader *Reader) (OpType, int) {
	length := reader.Length()
	if length == 0 {
		return NOP, 0
	}

	correlationID := int32(0)
	framelength := 4          // account for the length field
	lenbuf := make([]byte, 8) // Peek the first eight bytes
	n, err := reader.PeekFull(lenbuf)
	if err == nil {
		framelength += int(binary.BigEndian.Uint32(lenbuf[:4]))
		correlationID = int32(binary.BigEndian.Uint32(lenbuf[4:]))
	} else {
		// Need more data
		return MORE, 8 - n
	}

	if reply {
		// Replies are always passed as-is. No need to parse them
		// on top of the frame length and correlation ID.
		p.connection.Log(khulnasoft.EntryType_Response,
			&khulnasoft.LogEntry_Kafka{Kafka: &khulnasoft.KafkaLogEntry{
				CorrelationId: correlationID,
			}})
		return PASS, framelength
	}

	// Ask for more if full frame has not been received yet
	if length < framelength {
		// Not enough data, ask for more and try again
		return MORE, framelength - length
	}

	req, err := kafkalib.ReadRequest(reader)
	if err != nil {
		logrus.WithError(err).Debug("Unable to parse Kafka request; closing Kafka connection")
		p.connection.Log(khulnasoft.EntryType_Denied,
			&khulnasoft.LogEntry_Kafka{Kafka: &khulnasoft.KafkaLogEntry{
				CorrelationId: correlationID,
				ErrorCode:     kafkalib.ErrInvalidMessage,
			}})
		return ERROR, int(ERROR_INVALID_FRAME_TYPE)
	}

	logEntry := &khulnasoft.LogEntry_Kafka{Kafka: &khulnasoft.KafkaLogEntry{
		CorrelationId: correlationID,
		ApiVersion:    int32(req.GetVersion()),
		ApiKey:        int32(req.GetAPIKey()),
		Topics:        req.GetTopics(),
	}}
	if p.connection.Matches(req) {
		p.connection.Log(khulnasoft.EntryType_Request, logEntry)
		return PASS, framelength
	}
	logEntry.Kafka.ErrorCode = kafkalib.ErrTopicAuthorizationFailed

	resp, err := req.CreateAuthErrorResponse()
	if err != nil {
		logrus.WithError(err).Debug("Unable to create Kafka response")
	} else {
		// inject response
		p.connection.Inject(!reply, resp.GetRaw())
	}

	p.connection.Log(khulnasoft.EntryType_Denied, logEntry)
	return DROP, framelength
}
