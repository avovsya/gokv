package session

import (
	"errors"
	"fmt"
	"github.com/avovsya/gokv/store"
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
	var err error
	if len(args) != 1 {
		return errors.New("GET command has exactly one arg: GET <key>")
	}

	if value, err := store.Get(args[0]); err == nil {
		err = write(fmt.Sprintf("OK %s", value), s.writer)
		if err != nil {
			return err
		}
	}

	return err
}

func (s Session) Put(args []string) error {
	var err error
	if len(args) != 2 {
		return errors.New("PUT command has exactly two args: PUT <key> <value>")
	}

	if err := store.Put(args[0], args[1]); err == nil {
		err = write("OK", s.writer)
		if err != nil {
			return err
		}
	}

	return err
}

func (s Session) Delete(args []string) error {
	var err error
	if len(args) != 1 {
		return errors.New("DELETE command has exactly one arg: DELETE <key>")
	}

	if err := store.Delete(args[0]); err == nil {
		err = write("OK", s.writer)
		if err != nil {
			return err
		}
	}

	return err
}

func (s Session) Unknown(resp string) error {
	err := write(fmt.Sprintf("UNKNOWN_CMD %s", resp), s.writer)
	if err != nil {
		return err
	}
	return nil
}
