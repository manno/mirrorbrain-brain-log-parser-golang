
// line 1 "mirrorbrain_parser.rl"
package main

import (
	"os"
	"time"
	"log"
	"bufio"
	"fmt"
)

/*
#LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\" \        
#            want:%{WANT}e give:%{GIVE}e r:%{MB_REALM}e %{X-MirrorBrain-Mirror}o \
#            %{MB_CONTINENT_CODE}e:%{MB_COUNTRY_CODE}e ASN:%{ASN}e P:%{PFX}e \    
#            size:%{MB_FILESIZE}e %{Range}i" combined_redirect                    
#
#
#192.168.122.1 - - [29/Dec/2013:18:54:41 +0100] "GET /congress/2013/mp4/30c3-5415-de-en-Der_tiefe_Staat_h264-iprod.mp4.torrent HTTP/1.0" 200 11668 "-" "Transmission/2.77"     want:torrent give:torrent r:- -     EU:DE ASN:- P:-     size:126539847 -
*/

const ApacheDate = "02/Jan/2006:15:04:05 -0700"

type Entry struct {
	Date time.Time
	Agent string
	Asn string
	Country string
	GivenType string
	Ip string
	Mirror string
	Net string
	RedirBytes string
	RedirSize string
	Referer string
	Region string
	RequestMethod string
	RequestPath string
	RequestProto string
	RequestType string
	ReturnCode string
	Size string
	Parsed bool
}
type Entries []Entry


// line 47 "mirrorbrain_parser.rl"

// line 52 "mirrorbrain_parser.go"
var _mirrorbrain_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 19, 2, 18, 
	19, 
}

var _mirrorbrain_key_offsets []byte = []byte{
	0, 0, 3, 7, 8, 9, 10, 11, 
	12, 13, 14, 15, 16, 19, 20, 21, 
	22, 23, 24, 25, 26, 27, 28, 29, 
	30, 31, 33, 34, 35, 37, 40, 42, 
	45, 46, 47, 48, 49, 50, 51, 52, 
	53, 55, 56, 57, 58, 59, 63, 64, 
	65, 66, 67, 68, 69, 73, 74, 75, 
	76, 83, 84, 85, 86, 94, 95, 102, 
	103, 104, 105, 106, 107, 114, 115, 116, 
	117, 119, 122, 124, 125, 126, 127, 128, 
	131, 132, 134, 136, 137, 138, 139, 140, 
	141, 144, 149, 152, 159, 166, 173, 180, 
	181, 182, 183, 184, 185, 186, 187, 188, 
	189, 190, 191, 192, 193, 194, 195, 196, 
	197, 198, 199, 200, 201, 202, 203, 204, 
	205, 206, 207, 208, 209, 210, 211, 212, 
	213, 214, 215, 215, 
}

var _mirrorbrain_trans_keys []byte = []byte{
	46, 48, 57, 32, 46, 48, 57, 45, 
	32, 45, 32, 91, 93, 93, 32, 34, 
	71, 72, 80, 69, 84, 32, 32, 32, 
	72, 84, 84, 80, 47, 49, 46, 48, 
	49, 34, 32, 48, 57, 32, 48, 57, 
	48, 57, 32, 48, 57, 34, 34, 34, 
	32, 34, 34, 34, 32, 32, 119, 97, 
	110, 116, 58, 45, 102, 114, 116, 32, 
	103, 105, 118, 101, 58, 45, 102, 114, 
	116, 32, 114, 58, 45, 48, 57, 65, 
	90, 97, 122, 32, 32, 32, 32, 45, 
	48, 57, 65, 90, 97, 122, 58, 45, 
	48, 57, 65, 90, 97, 122, 32, 65, 
	83, 78, 58, 45, 48, 57, 65, 90, 
	97, 122, 32, 80, 58, 46, 57, 32, 
	46, 57, 32, 115, 105, 122, 101, 58, 
	45, 48, 57, 32, 45, 98, 10, 13, 
	121, 116, 101, 115, 61, 45, 48, 57, 
	10, 13, 45, 48, 57, 32, 48, 57, 
	32, 48, 57, 65, 90, 97, 122, 32, 
	48, 57, 65, 90, 97, 122, 58, 48, 
	57, 65, 90, 97, 122, 32, 48, 57, 
	65, 90, 97, 122, 105, 108, 101, 101, 
	100, 105, 114, 101, 99, 116, 111, 114, 
	114, 101, 110, 105, 108, 101, 101, 100, 
	105, 114, 101, 99, 116, 111, 114, 114, 
	101, 110, 69, 65, 68, 79, 83, 10, 
	
}

var _mirrorbrain_single_lengths []byte = []byte{
	0, 1, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 3, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 0, 1, 1, 0, 1, 0, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	2, 1, 1, 1, 1, 4, 1, 1, 
	1, 1, 1, 1, 4, 1, 1, 1, 
	1, 1, 1, 1, 2, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	0, 1, 2, 1, 1, 1, 1, 1, 
	1, 2, 2, 1, 1, 1, 1, 1, 
	1, 3, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 0, 1, 
}

var _mirrorbrain_range_lengths []byte = []byte{
	0, 1, 1, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 1, 1, 1, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	3, 0, 0, 0, 3, 0, 3, 0, 
	0, 0, 0, 0, 3, 0, 0, 0, 
	1, 1, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 1, 3, 3, 3, 3, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 
}

var _mirrorbrain_index_offsets []int16 = []int16{
	0, 0, 3, 7, 9, 11, 13, 15, 
	17, 19, 21, 23, 25, 29, 31, 33, 
	35, 37, 39, 41, 43, 45, 47, 49, 
	51, 53, 55, 57, 59, 61, 64, 66, 
	69, 71, 73, 75, 77, 79, 81, 83, 
	85, 88, 90, 92, 94, 96, 101, 103, 
	105, 107, 109, 111, 113, 118, 120, 122, 
	124, 129, 131, 133, 135, 141, 143, 148, 
	150, 152, 154, 156, 158, 163, 165, 167, 
	169, 171, 174, 177, 179, 181, 183, 185, 
	188, 190, 193, 196, 198, 200, 202, 204, 
	206, 209, 214, 217, 222, 227, 232, 237, 
	239, 241, 243, 245, 247, 249, 251, 253, 
	255, 257, 259, 261, 263, 265, 267, 269, 
	271, 273, 275, 277, 279, 281, 283, 285, 
	287, 289, 291, 293, 295, 297, 299, 301, 
	303, 305, 307, 308, 
}

var _mirrorbrain_indicies []byte = []byte{
	0, 0, 1, 2, 3, 3, 1, 4, 
	1, 5, 1, 6, 1, 7, 1, 8, 
	1, 1, 9, 11, 10, 12, 1, 13, 
	1, 14, 15, 16, 1, 17, 1, 18, 
	1, 19, 1, 1, 20, 22, 21, 23, 
	1, 24, 1, 25, 1, 26, 1, 27, 
	1, 28, 1, 29, 1, 30, 1, 31, 
	1, 32, 1, 33, 1, 34, 35, 1, 
	36, 1, 37, 38, 1, 39, 1, 1, 
	40, 42, 41, 43, 1, 44, 1, 1, 
	45, 47, 46, 48, 1, 48, 49, 1, 
	50, 1, 51, 1, 52, 1, 53, 1, 
	54, 55, 56, 57, 1, 58, 1, 59, 
	1, 60, 1, 61, 1, 62, 1, 63, 
	1, 64, 65, 66, 67, 1, 68, 1, 
	69, 1, 70, 1, 71, 72, 72, 72, 
	1, 73, 1, 1, 74, 76, 75, 77, 
	78, 79, 79, 79, 1, 80, 1, 81, 
	82, 82, 82, 1, 83, 1, 84, 1, 
	85, 1, 86, 1, 87, 1, 88, 89, 
	89, 89, 1, 90, 1, 91, 1, 92, 
	1, 93, 1, 94, 95, 1, 96, 97, 
	1, 98, 1, 99, 1, 100, 1, 101, 
	1, 102, 103, 1, 104, 1, 105, 106, 
	1, 107, 108, 1, 109, 1, 110, 1, 
	111, 1, 112, 1, 113, 1, 114, 114, 
	1, 107, 108, 114, 114, 1, 104, 115, 
	1, 90, 116, 116, 116, 1, 83, 117, 
	117, 117, 1, 80, 79, 79, 79, 1, 
	73, 118, 118, 118, 1, 119, 1, 120, 
	1, 121, 1, 122, 1, 123, 1, 124, 
	1, 125, 1, 126, 1, 127, 1, 121, 
	1, 128, 1, 129, 1, 130, 1, 131, 
	1, 127, 1, 132, 1, 133, 1, 134, 
	1, 135, 1, 136, 1, 137, 1, 138, 
	1, 139, 1, 140, 1, 134, 1, 141, 
	1, 142, 1, 143, 1, 144, 1, 140, 
	1, 145, 1, 146, 1, 18, 1, 147, 
	1, 17, 1, 1, 148, 1, 
}

var _mirrorbrain_trans_targs []byte = []byte{
	2, 0, 3, 2, 4, 5, 6, 7, 
	8, 9, 9, 10, 11, 12, 13, 125, 
	128, 14, 15, 16, 17, 17, 18, 19, 
	20, 21, 22, 23, 24, 25, 26, 27, 
	28, 29, 30, 29, 31, 32, 31, 33, 
	34, 34, 35, 36, 37, 38, 38, 39, 
	40, 41, 42, 43, 44, 45, 46, 110, 
	113, 120, 47, 48, 49, 50, 51, 52, 
	53, 95, 98, 105, 54, 55, 56, 57, 
	94, 58, 59, 59, 60, 60, 61, 93, 
	62, 63, 92, 64, 65, 66, 67, 68, 
	69, 91, 70, 71, 72, 73, 74, 73, 
	74, 75, 76, 77, 78, 79, 80, 90, 
	81, 82, 83, 130, 131, 84, 85, 86, 
	87, 88, 89, 90, 91, 92, 94, 96, 
	97, 53, 99, 100, 101, 102, 103, 104, 
	106, 107, 108, 109, 111, 112, 46, 114, 
	115, 116, 117, 118, 119, 121, 122, 123, 
	124, 126, 127, 129, 130, 
}

var _mirrorbrain_trans_actions []byte = []byte{
	1, 0, 5, 0, 0, 0, 0, 0, 
	0, 1, 0, 3, 0, 0, 1, 1, 
	1, 0, 0, 7, 1, 0, 9, 1, 
	0, 0, 0, 0, 0, 0, 0, 11, 
	0, 1, 13, 0, 1, 15, 0, 0, 
	1, 0, 17, 0, 0, 1, 0, 19, 
	0, 0, 0, 0, 0, 0, 1, 1, 
	1, 1, 21, 0, 0, 0, 0, 0, 
	1, 1, 1, 1, 23, 0, 0, 1, 
	1, 25, 1, 0, 27, 0, 0, 0, 
	0, 1, 1, 29, 0, 0, 0, 0, 
	1, 1, 31, 0, 0, 1, 33, 0, 
	0, 0, 0, 0, 0, 0, 1, 1, 
	35, 1, 1, 39, 39, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 37, 
}

const mirrorbrain_start int = 1
const mirrorbrain_first_final int = 130
const mirrorbrain_error int = 0

const mirrorbrain_en_main int = 1


// line 48 "mirrorbrain_parser.rl"

func parseTime(timeStr string) time.Time {
	date, e :=  time.Parse(ApacheDate, timeStr)
	if e!=nil {
		return time.Now()
	}
	return date
}

func mirrorbrain(data string) (entry *Entry, err error) {
	cs, p, pe := 0, 0, len(data)

	marker := 0
	entry = new(Entry)

	
// line 283 "mirrorbrain_parser.go"
	{
	cs = mirrorbrain_start
	}

// line 288 "mirrorbrain_parser.go"
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_mirrorbrain_key_offsets[cs])
	_trans = int(_mirrorbrain_index_offsets[cs])

	_klen = int(_mirrorbrain_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _mirrorbrain_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _mirrorbrain_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_mirrorbrain_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _mirrorbrain_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _mirrorbrain_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_mirrorbrain_indicies[_trans])
	cs = int(_mirrorbrain_trans_targs[_trans])

	if _mirrorbrain_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_mirrorbrain_trans_actions[_trans])
	_nacts = uint(_mirrorbrain_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _mirrorbrain_actions[_acts-1] {
		case 0:
// line 64 "mirrorbrain_parser.rl"

 marker = p 
		case 1:
// line 66 "mirrorbrain_parser.rl"

 entry.Date = parseTime(data[marker:p]) 
		case 2:
// line 67 "mirrorbrain_parser.rl"

 entry.Ip = data[marker:p] 
		case 3:
// line 68 "mirrorbrain_parser.rl"

 entry.RequestMethod = data[marker:p] 
		case 4:
// line 69 "mirrorbrain_parser.rl"

 entry.RequestPath = data[marker:p] 
		case 5:
// line 70 "mirrorbrain_parser.rl"

 entry.RequestProto = data[marker:p] 
		case 6:
// line 71 "mirrorbrain_parser.rl"

 entry.ReturnCode = data[marker:p] 
		case 7:
// line 72 "mirrorbrain_parser.rl"

 entry.Size = data[marker:p] 
		case 8:
// line 73 "mirrorbrain_parser.rl"

 entry.Referer = data[marker:p] 
		case 9:
// line 74 "mirrorbrain_parser.rl"

 entry.Agent = data[marker:p] 
		case 10:
// line 75 "mirrorbrain_parser.rl"

 entry.RequestType = data[marker:p] 
		case 11:
// line 76 "mirrorbrain_parser.rl"

 entry.GivenType = data[marker:p] 
		case 12:
// line 77 "mirrorbrain_parser.rl"

 entry.Region = data[marker:p] 
		case 13:
// line 78 "mirrorbrain_parser.rl"

 entry.Mirror = data[marker:p] 
		case 14:
// line 79 "mirrorbrain_parser.rl"

 entry.Country = data[marker:p] 
		case 15:
// line 80 "mirrorbrain_parser.rl"

 entry.Asn = data[marker:p] 
		case 16:
// line 81 "mirrorbrain_parser.rl"

 entry.Net = data[marker:p] 
		case 17:
// line 82 "mirrorbrain_parser.rl"

 entry.RedirSize = data[marker:p] 
		case 18:
// line 83 "mirrorbrain_parser.rl"

 entry.RedirBytes = data[marker:p] 
		case 19:
// line 85 "mirrorbrain_parser.rl"

 entry.Parsed = true 
// line 447 "mirrorbrain_parser.go"
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

// line 125 "mirrorbrain_parser.rl"


	return entry, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		entry, err := mirrorbrain(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(entry)
	}

}

/*
next unless r.request_method == 'GET'
next unless r.given_type == 'redirect'
next unless CONG.match r.request_path
#puts [r.mirror, r.redir_size].join(',')
sums[r.mirror] += r.redir_size.to_i

sums.each { |m,s| 
printf "%s = %d mb\n", m, s/1024/1024
  }
  */
