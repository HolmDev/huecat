package main

import (
	"strconv"
)

func hex2dec(input string) int {
	dec, err := strconv.ParseInt(input, 16, 32);

	if err != nil {
		panic(err);
	}

	return int(dec);
}

func colorify(input string) string {
	r, g, b := 0, 0, 0;

	r = hex2dec(input[1:3]);
	g = hex2dec(input[3:5]);
	b = hex2dec(input[5:7]);

	average := (r+g+b)/3;

	sr := strconv.Itoa(r)
	sg := strconv.Itoa(g)
	sb := strconv.Itoa(b)

	seq := "\033[48;2;"+sr+";"+sg+";"+sb+"m";

	if average > 128 {
		seq += "\033[30m"
	}


	return seq+input+"\033[m";
}
