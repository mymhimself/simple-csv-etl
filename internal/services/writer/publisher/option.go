package publisher

import (
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
	"github.com/spf13/viper"
)

// ─── Publisher Option ────────────────────────────────────────────────────────

type PublisherOption func(*iPublisher) error

func PublisherOptionWithRMQHost(host string) PublisherOption {
	return func(p *iPublisher) error {
		p.config.host = host
		return nil
	}
}

func PublisherOptionWithRMQExchangeName(exchangeName string) PublisherOption {
	return func(p *iPublisher) error {
		p.config.exchange = exchangeName
		return nil
	}
}

func PublisherOptionWithRMQKey(key string) PublisherOption {
	return func(p *iPublisher) error {
		p.config.key = key
		return nil
	}
}

func PublisherOptionWithRMQExchangeType(exchangeType string) PublisherOption {
	return func(p *iPublisher) error {
		p.config.exchangeType = exchangeType
		return nil
	}
}

func PublisherOptionFromViper(v *viper.Viper) PublisherOption {
	return func(p *iPublisher) error {
		var ops []PublisherOption
		if val := v.GetString(constants.Host); val != "" {
			ops = append(ops, PublisherOptionWithRMQHost(val))
		}

		if val := v.GetString(constants.ExchangeName); val != "" {
			ops = append(ops, PublisherOptionWithRMQExchangeName(val))
		}

		if val := v.GetString(constants.Key); val != "" {
			ops = append(ops, PublisherOptionWithRMQKey(val))
		}

		if val := v.GetString(constants.ExchangeType); val != "" {
			ops = append(ops, PublisherOptionWithRMQExchangeType(val))
		}

		for _, fn := range ops {
			err := fn(p)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
