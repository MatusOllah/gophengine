package config

type KeyError struct {
	Op  string
	Key string
	Err error
}

func (e *KeyError) Error() string {
	return e.Op + " " + e.Key + ": " + e.Err.Error()
}

func (e *KeyError) Unwrap() error {
	return e.Err
}
