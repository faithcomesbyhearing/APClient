package main

import "github.com/faithcomesbyhearing/fcbh-dataset-io/decode_yaml/request"

func templateA(req request.Request) request.Request {
	req.IsNew = true
	req.NotifyOk = []string{"jbarndt@fcbhmail.org",
		"ezornes@fcbhmail.org",
		"gfiddes@fcbhmail.org",
		"edomschot@fcbhmail.org"}
	req.NotifyErr = []string{"jbarndt@fcbhmail.org",
		"gary@shortsands.com",
		"ezornes@fcbhmail.org"}
	return req
}
