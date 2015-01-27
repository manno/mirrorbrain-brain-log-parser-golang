all:
	ragel -s -Z mirrorbrain_parser.rl
	go run mirrorbrain_parser.go sample.log 
