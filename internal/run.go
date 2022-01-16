package internal

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Run() error {

	wr, err := getDownstream()
	if err != nil {
		return err
	}
	defer func() {
		closer, isCloser := wr.(io.Closer)
		if isCloser {
			_ = closer.Close()
		}
	}()

	in := os.Stdin
	err = pump(in, wr)
	if err != nil {
		return err
	}

	return nil
}

func pump(in io.Reader, wr io.Writer) error {

	buf := make([]byte, 5)
	tsWritten := false
	nlWritten := false
	for {
		n, err := in.Read(buf)
		if err == io.EOF {
			if !nlWritten {
				_, err := wr.Write([]byte("\n"))
				if err != nil {
					return err
				}
			}
			return nil
		} else if err != nil {
			return err
		}
		buf = buf[:n]

		prev := 0
		for i := 0; i < len(buf); i++ {

			if i == len(buf)-1 || buf[i] == '\n' {
				if !tsWritten {
					now := time.Now()
					_, err := fmt.Fprintf(wr, "%d-%02d-%02dT%02d:%02d:%02d ",
						now.Year(), now.Month(), now.Day(),
						now.Hour(), now.Minute(), now.Second())
					if err != nil {
						return err
					}
					tsWritten = true
				}

				_, err = wr.Write(buf[prev : i+1])
				if err != nil {
					return err
				}
				prev = i + 1

				nlWritten = false
				if buf[i] == '\n' {
					tsWritten = false
					nlWritten = true
				}
			}
		}
	}

}

func getDownstream() (io.Writer, error) {
	wr := os.Stderr
	if len(os.Args) == 2 {
		var err error
		wr, err = os.OpenFile(os.Args[1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return nil, err
		}
	}

	return wr, nil
}
