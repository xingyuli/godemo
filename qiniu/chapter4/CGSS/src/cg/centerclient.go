package cg

import (
	"encoding/json"
	"errors"

	"ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return wrapErr(err, resp)
}

func (client *CenterClient) RemovePlayer(name string) error {
	resp, err := client.Call("removeplayer", name)
	if err == nil && resp.Code == "200" {
		return nil
	}
	return wrapErr(err, resp)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, err := client.Call("listplayer", params)
	if err == nil && resp.Code == "200" {
		err = json.Unmarshal([]byte(resp.Body), &ps)
	} else {
		err = wrapErr(err, resp)
	}
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message}

	b, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	resp, err := client.Call("broadcast", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return wrapErr(err, resp)
}

func wrapErr(err error, resp *ipc.Response) error {
	if err != nil {
		return err
	}
	return errors.New(resp.Code)
}
