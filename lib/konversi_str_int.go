package lib

import "strconv"

func KonversiInt(str string) (int, error) {
	var number int
	var err error

	number, err = strconv.Atoi(str)

	if err == nil {
		return number, nil
	}
	return number, err
}
