package core

var (
	inputBuilders  = make(map[string]InputBuilder)
	filterBuilders = make(map[string]FilterBuilder)
	outputBuilders = make(map[string]OutputBuilder)
	codecBuilders  = make(map[string]CodecBuilder)
)

type InputBuilder interface {
	NewConf() InputConf
	Build(conf InputConf) Input
}
type FilterBuilder interface {
	NewConf() FilterConf
	Build(conf FilterConf) Filter
}
type OutputBuilder interface {
	NewConf() OutputConf
	Build(conf OutputConf) Output
}
type CodecBuilder interface {
	NewConf() CodecConf
	Build(conf CodecConf) Codec
}

func RegInput(name string, builder InputBuilder) {
	inputBuilders[name] = builder
}

func RegFilter(name string, builder FilterBuilder) {
	filterBuilders[name] = builder
}

func RegOutput(name string, builder OutputBuilder) {
	outputBuilders[name] = builder
}

func RegCodec(name string, builder CodecBuilder) {
	codecBuilders[name] = builder
}
