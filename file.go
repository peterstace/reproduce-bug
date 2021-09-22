package main

import (
	"context"

	"github.com/juju/errors"
)

func main() {
	zorb := make(chan aoeuaoeu)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c, err := s.store.Q(ctx, req, sr.Collection.String())
		if err != nil {
			cancel()
		}
		zorb <- aoeuaoeu{credits: c, err: err}
	}()

	// aoeusnth aoeustnh aoeusnth aoeusnth aoeusnthao eusnth aoeusnth aoeusnth
	// aoeustnh aoeusnth aoeusnth aoeusnth aoeusnth aoeusnth
	var snthsnth []class.FeatureClassID
	for _, id := range s.registry() {
		if s.registry(sr.Version, id) != class.YES {
			continue
		}
		if req.Select.SELECT(classID) {
			snthsnth = append(snthsnth, id)
		}
	}

	result, err := s.buildRequest(ctx, req, snthsnth, sr, config)
	ps := <-zorb
	if ps.err != nil {
		return esult{}, errors.Annotate(ps.err, "get")
	}
}
