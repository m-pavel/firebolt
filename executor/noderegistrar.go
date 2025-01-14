package executor

import (
	"reflect"

	"github.com/m-pavel/firebolt"
	"github.com/m-pavel/firebolt/node/elasticsearch"

	"github.com/m-pavel/firebolt/node"
	"github.com/m-pavel/firebolt/node/kafkaconsumer"
	"github.com/m-pavel/firebolt/node/kafkaproducer"
	"github.com/m-pavel/firebolt/node/syslogparser"
)

// RegisterBuiltinSourceTypes initializes the node registry with all built-in source types
func RegisterBuiltinSourceTypes() {
	node.GetRegistry().RegisterSourceType("kafkaconsumer", func() node.Source {
		return &kafkaconsumer.KafkaConsumer{}
	}, reflect.TypeOf(([]byte)(nil)))
}

// RegisterBuiltinNodeTypes initializes the node registry with all built-in node types
func RegisterBuiltinNodeTypes() {
	node.GetRegistry().RegisterNodeType("syslogparser", func() node.Node {
		return &syslogparser.SyslogParser{}
	}, reflect.TypeOf(([]byte)(nil)), reflect.TypeOf(""))

	node.GetRegistry().RegisterNodeType("kafkaproducer", func() node.Node {
		return &kafkaproducer.KafkaProducer{}
	}, reflect.TypeOf((*firebolt.ProduceRequest)(nil)).Elem(), nil)

	node.GetRegistry().RegisterNodeType("errorkafkaproducer", func() node.Node {
		return &kafkaproducer.ErrorProducer{}
	}, reflect.TypeOf(&firebolt.EventError{}), nil)

	node.GetRegistry().RegisterNodeType("elasticsearch", func() node.Node {
		return &elasticsearch.Elasticsearch{}
	}, reflect.TypeOf(elasticsearch.IndexRequest{}), reflect.TypeOf(elasticsearch.IndexRequest{}))
}
