package writer

import (
	"github.com/mymhimself/simple-csv-reader/internal/models/writer"
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
)

type InitOption func(*iWriter) error

func InitModel(v writer.IWriter) InitOption {
	return func(s *iWriter) error {
		s.model = v
		return nil
	}
}

func InitConfigsFromViper(v *viper.Viper) InitOption {
	return func(s *iWriter) error {
		s.config.rmq.queueName = v.GetString(constants.ConsumerQueueName)
		s.config.rmq.host = v.GetString(constants.ConsumerHost)
		s.config.rmq.exchangeName = v.GetString(constants.ConsumerExchangeName)
		s.config.rmq.key = v.GetString(constants.ConsumerKey)
		s.config.rmq.consumerName = v.GetString(constants.ConsumerName)
		s.config.numberOfConsumingThread = v.GetInt(constants.ConsumerThreads)
		return nil
	}
}
