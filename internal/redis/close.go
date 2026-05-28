package redis

import "errors"

func (r *Redis) Close() error {
	if r.client == nil {
		return errors.New("already closed")
	}

	return r.client.Close()
}
