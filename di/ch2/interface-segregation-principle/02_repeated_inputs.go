package interface_segregation_principle

import "errors"

type Value interface {
	Value(key interface{}) interface{}
}

type Monitor interface {
	Done() <-chan struct{}
}

func EncryptV2(keyValue Value, monitor Monitor, data []byte) ([]byte, error) {
	stop := monitor.Done()
	result := make(chan []byte, 1)
	go func() {
		defer close(result)
		keyRaw := keyValue.Value("encryption-key")
		if keyRaw == nil {
			panic("encryption key not found in context")
		}
		key := keyRaw.([]byte)
		// perform encryption
		ciperText := performEncryption(key, data)
		result <- ciperText
	}()
	select {
	case ciperText := <- result:
		// happy path
		return ciperText, nil
	case <- stop:
		// cancelled
		return nil, errors.New("operation cancelled")
	}
}