package session

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"
)

type Session struct {
	addr   string
	reader *bufio.Reader
	writer *bufio.Writer
	cmd    chan []string
	err    chan error
}

func NewSession(addr string, r *bufio.Reader, w *bufio.Writer) Session {
	return Session{
		addr:   addr,
		reader: r,
		writer: w,
		cmd:    make(chan []string),
		err:    make(chan error),
	}
}

func (s Session) Talk() Status {
	s.Ready()
	go s.readCmd()

	for {
		select {
		case <-time.After(5 * time.Second): //TODO: configurable timeout
			return Code("CLOSED", "SERVER_TIMEOUT", "Session closed after timeout")
		case err := <-s.err:
			return Error(err)
		case cmd := <-s.cmd:
			status := s.processCmd(cmd)
			if status.code == "ERROR" ||
				status.code == "CLOSED" {
				return status
			}
		}
	}
}

func (s Session) processCmd(cmd []string) Status {
	var cmdError error

	log.Printf("Command '%v' received from %s", cmd[0], s.addr)

	switch {
	case cmd[0] == "CLOSE":
		return Code("CLOSED", "CLIENT", "Session closed by client")
	case cmd[0] == "PING":
		cmdError = s.Pong()
	case cmd[0] == "GET":
		cmdError = s.Get(cmd[1:])
	case cmd[0] == "PUT":
		cmdError = s.Put(cmd[1:])
	case cmd[0] == "DELETE":
		cmdError = s.Delete(cmd[1:])
	default:
		cmdError = s.Unknown(strings.Join(cmd, " "))
	}
	if cmdError != nil {
		return Error(cmdError)
	}
	// TODO try to notify client about the error
	return Ok()
}

func (s Session) readCmd() {
	for {
		resp, err := read(s.reader)
		if err != nil {
			s.err <- err
			return
		}

		cmd := strings.Split(resp, " ")
		for i, s := range cmd {
			cmd[i] = strings.Trim(s, " ")
			cmd[i] = strings.Trim(s, "\x00") // Remove null bytes
		}

		cmd[0] = strings.ToUpper(cmd[0])

		s.cmd <- cmd
	}
}

func write(msg string, w *bufio.Writer) error {
	w.WriteString(fmt.Sprintf("%s\n", msg))
	return w.Flush()
}

func read(r *bufio.Reader) (string, error) {
	var buffer []byte
	buffer = make([]byte, 10)

	for {
		line, isPrefix, err := r.ReadLine()

		if err != nil {
			return "", err
		}

		buffer = append(buffer, line...)

		if !isPrefix {
			return string(buffer[:]), nil
		}
	}
}
