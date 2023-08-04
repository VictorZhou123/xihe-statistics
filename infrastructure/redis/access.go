package redis

import (
	"context"
	"encoding/json"
	"errors"
	"project/xihe-statistics/infrastructure/repositories"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	ipPrefix = "ip_"
)

func NewAccessMapper(expire time.Duration) repositories.AccessMapper {
	return access{cli: NewDBRedis(int(expire))}
}

type access struct {
	cli dbRedis
}

func (impl access) Upsert(
	do repositories.AccessInputDO,
) (err error) {
	key := ipPrefix + do.IP

	f := func(ctx context.Context) error {
		return impl.cli.AtomicOpt(ctx, key, do, impl.upsert)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (impl access) upsert(
	ctx context.Context,
	val interface{},
) error {
	do, ok := val.(repositories.AccessInputDO)
	if !ok {
		logrus.Debugf("upsert convert val error")

		return errors.New("cannot convert val")
	}

	key := ipPrefix + do.IP

	// get value
	res, err := impl.cli.Get(ctx, key).Result()
	if err != nil {
		// create
		data, err2 := json.Marshal(map[string]int64{
			do.URL: 1,
		})
		if err2 != nil {
			logrus.Debugf("marshal map error: %s", err2)
			return err2
		}

		c := impl.cli.Create(ctx, key, data)

		return c.Err()
	}

	// update
	var vo map[string]int64
	if err := json.Unmarshal([]byte(res), &vo); err != nil {
		logrus.Debugf("upsert convert res error")

		return errors.New("cannot convert res")
	}

	if count, ok := vo[do.URL]; ok {
		vo[do.URL] = count + 1
	} else {
		vo[do.URL] = 1
	}

	data2, err3 := json.Marshal(vo)
	if err3 != nil {
		logrus.Debugf("upsert convert vo error")

		return errors.New("cannot convert vo")
	}
	cmd2 := impl.cli.Set(ctx, key, string(data2))
	if cmd2.Err() != nil {
		logrus.Debugf("upsert redis set error: %s", cmd2.Err().Error())

		return cmd2.Err()
	}

	return nil
}

func (impl access) Get(
	do repositories.AccessInputDO,
) (count int64, err error) {
	key := ipPrefix + do.IP

	var res string
	f := func(ctx context.Context) (err2 error) {
		res, err2 = impl.cli.Get(ctx, key).Result()
		if err2 != nil {
			return repositories.NewErrorNoData(err2)
		}

		return
	}

	if err = withContext(f); err != nil {
		if !repositories.IsErrorNoData(err) {
			return
		} else {
			err = nil
		}
	}

	vo := map[string]int64{}
	if res != "" {
		if err = json.Unmarshal([]byte(res), &vo); err != nil {
			logrus.Debugf("access get unmarshal error: %s", err.Error())

			return
		}
	}

	count = vo[do.URL]

	return
}
