package session

import (
	"fmt"
)

func (s Session) Ready() error {
	err := write("READY", s.writer)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Pong() error {
	err := write("PONG", s.writer)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Get(args []string) error {
	err := write("PUT", s.writer)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Put(args []string) error {
	err := write("PUT", s.writer)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Delete(args []string) error {
	err := write("DELETE", s.writer)
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Unknown(resp string) error {
	err := write(fmt.Sprintf("UNKNOWN_CMD %s", resp), s.writer)
	if err != nil {
		return err
	}
	return nil
}
