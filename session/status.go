package session

type Status struct {
	code   string
	reason string
	desc   string
	error  error
}

func Ok() Status {
	return Status{
		code: "OK",
	}
}

func Error(e error) Status {
	return Status{
		code:  "ERROR",
		error: e,
	}
}

func Code(code string, reason string, desc string) Status {
	return Status{
		code, reason, desc, nil,
	}
}
