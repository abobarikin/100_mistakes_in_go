package main

import "errors"

func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}
					
	return concat, nil
}

func concatenate(_, _ string) (string, error) {
	return "", nil
}