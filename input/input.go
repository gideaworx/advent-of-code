package input

import (
	"bufio"
	"io"
)

func ReadLines(in io.Reader) ([]string, error) {
	lines := []string{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadByteMatrix(in io.Reader) ([][]byte, error) {
	lines, err := ReadLines(in)
	if err != nil {
		return nil, err
	}

	m := make([][]byte, len(lines))

	for i := range lines {
		m[i] = []byte(lines[i])
	}

	return m, nil
}
