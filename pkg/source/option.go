package source

import (
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
)

type Option func(*rmqSource) error

func OptionWithHost(v string) Option {
	return func(rs *rmqSource) error {
		rs.config.host = v
		return nil
	}
}

func OptionWithQueueName(v string) Option {
	return func(rs *rmqSource) error {
		rs.config.queueName = v
		return nil
	}
}

func OptionWithKey(v string) Option {
	return func(rs *rmqSource) error {
		rs.config.key = v
		return nil
	}
}

func OptionWithExchangeName(v string) Option {
	return func(rs *rmqSource) error {
		rs.config.exchangeName = v
		return nil
	}
}

func OptionWithConsumerName(v string) Option {
	return func(rs *rmqSource) error {
		rs.config.name = v
		return nil
	}
}

func OptionWithArgs(v map[string]interface{}) Option {
	return func(rs *rmqSource) error {
		rs.config.args = v
		return nil
	}
}

func OptionWithXExpires(n int) Option {
	return func(rs *rmqSource) error {
		rs.config.xExpires = n
		return nil
	}
}

func OptionFromViper(v *viper.Viper) Option {
	return func(rs *rmqSource) error {
		var ops []Option

		if val := v.GetString(constants.ConsumerHost); val != "" {
			ops = append(ops, OptionWithHost(val))
		}
		if val := v.GetString(constants.ConsumerQueueName); val != "" {
			ops = append(ops, OptionWithQueueName(val))
		}
		if val := v.GetString(constants.ConsumerKey); val != "" {
			ops = append(ops, OptionWithKey(val))
		}
		if val := v.GetString(constants.ConsumerExchangeName); val != "" {
			ops = append(ops, OptionWithExchangeName(val))
		}
		if val := v.GetString(constants.ConsumerName); val != "" {
			ops = append(ops, OptionWithConsumerName(val))
		}

		for _, fn := range ops {
			err := fn(rs)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
