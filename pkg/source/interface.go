package source

type Source interface {
	ForeachN(fn func([]byte) error, n int) error
}

//go:generate mockgen -destination=../../test/mocks/Source.go -package=mocks -source=interface.go

//go:generate mockgen -destination=../../test/mocks/IRabbitConsumer.go -package mocks github.com/saage-tech/commons/rabbitmq IRabbitConsumer
