package evaluations

import (
	"github.com/hashicorp/nomad/api"
	"github.com/jippi/hashi-ui/backend/nomad/helper"
	"github.com/jippi/hashi-ui/backend/structs"
)

const (
	fetchedList = "NOMAD_FETCHED_EVALS"
	UnwatchList = "NOMAD_UNWATCH_EVALS"
	WatchList   = "NOMAD_WATCH_EVALS"
)

type list struct {
	action structs.Action
	client *api.Client
	query  *api.QueryOptions
}

func NewList(action structs.Action, client *api.Client, query *api.QueryOptions) *list {
	return &list{
		action: action,
		client: client,
		query:  query,
	}
}

func (w *list) Do() (*structs.Action, error) {
	evaluations, meta, err := w.client.Evaluations().List(w.query)
	if err != nil {
		return nil, err
	}

	if !helper.QueryChanged(w.query, meta) {
		return nil, nil
	}

	return &structs.Action{
		Type:    fetchedList,
		Payload: evaluations,
		Index:   meta.LastIndex,
	}, nil
}

func (w *list) Key() string {
	return "/evaluations/list"
}

func (w *list) IsMutable() bool {
	return false
}
