package digest

import (
	"github.com/dwarvesf/fortress-discord/pkg/adapter"
	"github.com/dwarvesf/fortress-discord/pkg/logger"
	"github.com/dwarvesf/fortress-discord/pkg/model"
)

type Digest struct {
	adapter adapter.IAdapter
	l       logger.Logger
}

func New(adapter adapter.IAdapter, l logger.Logger) DigestServicer {
	return &Digest{
		adapter: adapter,
		l:       l,
	}
}

func (e *Digest) GetInteralUpdates() ([]*model.Digest, error) {
	// get response from fortress
	adapterDigest, err := e.adapter.Fortress().GetInternalDigest()
	if err != nil {
		e.l.Error(err, "can't get open digest from fortress")
		return nil, err
	}

	// normalized into in-app model
	digest := adapterDigest.Data

	return digest, nil
}

func (e *Digest) GetExternalUpdates() ([]*model.Digest, error) {
	// get response from fortress
	adapterDigest, err := e.adapter.Fortress().GetExternalDigest()
	if err != nil {
		e.l.Error(err, "can't get open digest from fortress")
		return nil, err
	}

	// normalized into in-app model
	digest := adapterDigest.Data

	return digest, nil
}
