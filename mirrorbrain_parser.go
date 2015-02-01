
// line 1 "mirrorbrain_parser.rl"
package main

import (
	"os"
	"time"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
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
var PathFilter = os.Getenv("FILTER")

type Entry struct {
	Date time.Time
	Agent string
	Asn string
	Country string
	GivenType string
	Ip string
	Mirror string
	Net string
	RedirRange string
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


// line 50 "mirrorbrain_parser.rl"

// line 55 "mirrorbrain_parser.go"
var _mirrorbrain_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 19, 2, 0, 
	8, 2, 0, 9, 2, 9, 0, 2, 
	10, 14, 2, 18, 19, 
}

var _mirrorbrain_key_offsets []int16 = []int16{
	0, 0, 3, 7, 8, 9, 10, 11, 
	12, 13, 14, 15, 16, 20, 21, 22, 
	23, 24, 25, 26, 27, 28, 29, 30, 
	31, 32, 34, 35, 36, 38, 41, 44, 
	45, 46, 47, 48, 49, 50, 51, 52, 
	54, 57, 59, 61, 63, 65, 71, 73, 
	75, 77, 79, 81, 83, 89, 91, 93, 
	95, 103, 106, 108, 110, 119, 122, 130, 
	133, 135, 137, 139, 141, 149, 152, 154, 
	156, 159, 163, 166, 168, 170, 172, 174, 
	178, 180, 183, 186, 188, 190, 192, 194, 
	196, 200, 207, 209, 211, 215, 223, 231, 
	239, 241, 251, 260, 269, 278, 286, 298, 
	301, 304, 313, 322, 331, 339, 349, 358, 
	367, 376, 385, 394, 403, 412, 421, 430, 
	439, 448, 457, 466, 475, 484, 493, 502, 
	511, 520, 529, 538, 547, 555, 557, 559, 
	561, 564, 566, 568, 570, 572, 574, 576, 
	578, 580, 582, 584, 586, 588, 590, 592, 
	594, 596, 598, 600, 602, 604, 606, 608, 
	610, 612, 614, 617, 619, 621, 623, 625, 
	627, 629, 631, 633, 635, 637, 639, 641, 
	643, 645, 647, 649, 651, 653, 655, 657, 
	659, 661, 664, 665, 666, 667, 668, 669, 
	670, 671, 672, 673, 675, 676, 677, 678, 
	679, 680, 681, 682, 
}

var _mirrorbrain_trans_keys []byte = []byte{
	46, 48, 57, 32, 46, 48, 57, 45, 
	32, 45, 32, 91, 93, 93, 32, 34, 
	71, 72, 79, 80, 69, 84, 32, 32, 
	32, 72, 84, 84, 80, 47, 49, 46, 
	48, 49, 34, 32, 48, 57, 32, 48, 
	57, 45, 48, 57, 32, 34, 34, 34, 
	32, 34, 34, 34, 32, 34, 32, 34, 
	119, 34, 97, 34, 110, 34, 116, 34, 
	58, 34, 45, 102, 109, 114, 116, 32, 
	34, 34, 103, 34, 105, 34, 118, 34, 
	101, 34, 58, 34, 45, 102, 109, 114, 
	116, 32, 34, 34, 114, 34, 58, 34, 
	45, 48, 57, 65, 90, 97, 122, 32, 
	34, 45, 32, 34, 32, 34, 32, 34, 
	45, 48, 57, 65, 90, 97, 122, 34, 
	45, 58, 34, 45, 48, 57, 65, 90, 
	97, 122, 32, 34, 45, 34, 65, 34, 
	83, 34, 78, 34, 58, 34, 45, 48, 
	57, 65, 90, 97, 122, 32, 34, 45, 
	34, 80, 34, 58, 34, 45, 57, 32, 
	34, 45, 57, 32, 34, 115, 34, 105, 
	34, 122, 34, 101, 34, 58, 34, 45, 
	48, 57, 32, 34, 34, 45, 98, 10, 
	13, 34, 34, 121, 34, 116, 34, 101, 
	34, 115, 34, 61, 34, 45, 48, 57, 
	10, 13, 34, 44, 45, 48, 57, 32, 
	34, 34, 98, 32, 34, 48, 57, 32, 
	34, 48, 57, 65, 90, 97, 122, 32, 
	34, 48, 57, 65, 90, 97, 122, 34, 
	58, 48, 57, 65, 90, 97, 122, 32, 
	34, 32, 34, 45, 119, 48, 57, 65, 
	90, 97, 122, 34, 58, 97, 48, 57, 
	65, 90, 98, 122, 34, 58, 110, 48, 
	57, 65, 90, 97, 122, 34, 58, 116, 
	48, 57, 65, 90, 97, 122, 34, 58, 
	48, 57, 65, 90, 97, 122, 34, 45, 
	102, 109, 114, 116, 48, 57, 65, 90, 
	97, 122, 32, 34, 45, 34, 65, 103, 
	32, 34, 105, 48, 57, 65, 90, 97, 
	122, 32, 34, 108, 48, 57, 65, 90, 
	97, 122, 32, 34, 101, 48, 57, 65, 
	90, 97, 122, 32, 34, 48, 57, 65, 
	90, 97, 122, 32, 34, 101, 105, 48, 
	57, 65, 90, 97, 122, 32, 34, 116, 
	48, 57, 65, 90, 97, 122, 32, 34, 
	97, 48, 57, 65, 90, 98, 122, 32, 
	34, 52, 48, 57, 65, 90, 97, 122, 
	32, 34, 114, 48, 57, 65, 90, 97, 
	122, 32, 34, 114, 48, 57, 65, 90, 
	97, 122, 32, 34, 111, 48, 57, 65, 
	90, 97, 122, 32, 34, 114, 48, 57, 
	65, 90, 97, 122, 32, 34, 108, 48, 
	57, 65, 90, 97, 122, 32, 34, 105, 
	48, 57, 65, 90, 97, 122, 32, 34, 
	115, 48, 57, 65, 90, 97, 122, 32, 
	34, 116, 48, 57, 65, 90, 97, 122, 
	32, 34, 101, 48, 57, 65, 90, 97, 
	122, 32, 34, 100, 48, 57, 65, 90, 
	97, 122, 32, 34, 105, 48, 57, 65, 
	90, 97, 122, 32, 34, 114, 48, 57, 
	65, 90, 97, 122, 32, 34, 101, 48, 
	57, 65, 90, 97, 122, 32, 34, 99, 
	48, 57, 65, 90, 97, 122, 32, 34, 
	111, 48, 57, 65, 90, 97, 122, 32, 
	34, 114, 48, 57, 65, 90, 97, 122, 
	32, 34, 114, 48, 57, 65, 90, 97, 
	122, 32, 34, 101, 48, 57, 65, 90, 
	97, 122, 32, 34, 110, 48, 57, 65, 
	90, 97, 122, 32, 34, 48, 57, 65, 
	90, 97, 122, 34, 105, 34, 108, 34, 
	101, 34, 101, 105, 34, 116, 34, 97, 
	34, 52, 34, 114, 34, 114, 34, 111, 
	34, 114, 34, 108, 34, 105, 34, 115, 
	34, 116, 34, 101, 34, 100, 34, 105, 
	34, 114, 34, 101, 34, 99, 34, 111, 
	34, 114, 34, 114, 34, 101, 34, 110, 
	34, 105, 34, 108, 34, 101, 34, 101, 
	105, 34, 116, 34, 97, 34, 52, 34, 
	114, 34, 114, 34, 111, 34, 114, 34, 
	108, 34, 105, 34, 115, 34, 116, 34, 
	101, 34, 100, 34, 105, 34, 114, 34, 
	101, 34, 99, 34, 111, 34, 114, 34, 
	114, 34, 101, 34, 110, 32, 48, 57, 
	69, 65, 68, 80, 84, 73, 79, 78, 
	83, 79, 82, 83, 79, 80, 70, 73, 
	78, 34, 10, 34, 
}

var _mirrorbrain_single_lengths []byte = []byte{
	0, 1, 2, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 4, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 0, 1, 1, 0, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 2, 
	3, 2, 2, 2, 2, 6, 2, 2, 
	2, 2, 2, 2, 6, 2, 2, 2, 
	2, 3, 2, 2, 3, 3, 2, 3, 
	2, 2, 2, 2, 2, 3, 2, 2, 
	1, 2, 3, 2, 2, 2, 2, 2, 
	2, 3, 3, 2, 2, 2, 2, 2, 
	2, 5, 2, 2, 2, 2, 2, 2, 
	2, 4, 3, 3, 3, 2, 6, 3, 
	3, 3, 3, 3, 2, 4, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 2, 2, 2, 2, 
	3, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 3, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 2, 1, 1, 1, 1, 
	1, 1, 1, 2, 
}

var _mirrorbrain_range_lengths []byte = []byte{
	0, 1, 1, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 1, 1, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	3, 0, 0, 0, 3, 0, 3, 0, 
	0, 0, 0, 0, 3, 0, 0, 0, 
	1, 1, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 0, 0, 1, 3, 3, 3, 
	0, 3, 3, 3, 3, 3, 3, 0, 
	0, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 
}

var _mirrorbrain_index_offsets []int16 = []int16{
	0, 0, 3, 7, 9, 11, 13, 15, 
	17, 19, 21, 23, 25, 30, 32, 34, 
	36, 38, 40, 42, 44, 46, 48, 50, 
	52, 54, 56, 58, 60, 62, 65, 68, 
	70, 72, 74, 76, 78, 80, 82, 84, 
	87, 91, 94, 97, 100, 103, 110, 113, 
	116, 119, 122, 125, 128, 135, 138, 141, 
	144, 150, 154, 157, 160, 167, 171, 177, 
	181, 184, 187, 190, 193, 199, 203, 206, 
	209, 212, 216, 220, 223, 226, 229, 232, 
	236, 239, 243, 247, 250, 253, 256, 259, 
	262, 266, 273, 276, 279, 283, 289, 295, 
	301, 304, 312, 319, 326, 333, 339, 349, 
	353, 357, 364, 371, 378, 384, 392, 399, 
	406, 413, 420, 427, 434, 441, 448, 455, 
	462, 469, 476, 483, 490, 497, 504, 511, 
	518, 525, 532, 539, 546, 552, 555, 558, 
	561, 565, 568, 571, 574, 577, 580, 583, 
	586, 589, 592, 595, 598, 601, 604, 607, 
	610, 613, 616, 619, 622, 625, 628, 631, 
	634, 637, 640, 644, 647, 650, 653, 656, 
	659, 662, 665, 668, 671, 674, 677, 680, 
	683, 686, 689, 692, 695, 698, 701, 704, 
	707, 710, 713, 715, 717, 719, 721, 723, 
	725, 727, 729, 731, 734, 736, 738, 740, 
	742, 744, 746, 748, 
}

var _mirrorbrain_indicies []byte = []byte{
	0, 0, 1, 2, 3, 3, 1, 4, 
	1, 5, 1, 6, 1, 7, 1, 8, 
	1, 1, 9, 11, 10, 12, 1, 13, 
	1, 14, 15, 16, 17, 1, 18, 1, 
	19, 1, 20, 1, 1, 21, 23, 22, 
	24, 1, 25, 1, 26, 1, 27, 1, 
	28, 1, 29, 1, 30, 1, 31, 1, 
	32, 1, 33, 1, 34, 1, 35, 36, 
	1, 37, 38, 1, 39, 1, 40, 1, 
	42, 41, 44, 43, 45, 1, 46, 1, 
	48, 47, 50, 49, 51, 50, 49, 51, 
	50, 52, 49, 50, 53, 49, 50, 54, 
	49, 50, 55, 49, 50, 56, 49, 50, 
	57, 58, 59, 60, 61, 49, 62, 50, 
	49, 50, 63, 49, 50, 64, 49, 50, 
	65, 49, 50, 66, 49, 50, 67, 49, 
	50, 68, 69, 70, 71, 72, 49, 73, 
	50, 49, 50, 74, 49, 50, 75, 49, 
	50, 76, 77, 77, 77, 49, 78, 50, 
	79, 49, 49, 81, 80, 83, 84, 82, 
	85, 50, 86, 87, 87, 87, 49, 50, 
	86, 88, 49, 50, 89, 90, 90, 90, 
	49, 91, 50, 92, 49, 50, 93, 49, 
	50, 94, 49, 50, 95, 49, 50, 96, 
	49, 50, 97, 98, 98, 98, 49, 99, 
	50, 100, 49, 50, 101, 49, 50, 102, 
	49, 50, 103, 49, 104, 50, 105, 49, 
	106, 50, 107, 49, 50, 108, 49, 50, 
	109, 49, 50, 110, 49, 50, 111, 49, 
	50, 112, 113, 49, 114, 50, 49, 50, 
	115, 116, 49, 117, 118, 50, 49, 50, 
	119, 49, 50, 120, 49, 50, 121, 49, 
	50, 122, 49, 50, 123, 49, 50, 124, 
	124, 49, 117, 118, 50, 125, 124, 124, 
	49, 126, 50, 49, 50, 127, 49, 114, 
	50, 128, 49, 99, 50, 129, 129, 129, 
	49, 91, 50, 130, 130, 130, 49, 50, 
	88, 87, 87, 87, 49, 131, 84, 82, 
	132, 50, 86, 133, 87, 87, 87, 49, 
	50, 88, 134, 87, 87, 87, 49, 50, 
	88, 135, 87, 87, 87, 49, 50, 88, 
	136, 87, 87, 87, 49, 50, 137, 87, 
	87, 87, 49, 50, 138, 139, 140, 141, 
	142, 90, 90, 90, 49, 143, 50, 92, 
	49, 50, 93, 63, 49, 91, 50, 144, 
	130, 130, 130, 49, 91, 50, 145, 130, 
	130, 130, 49, 91, 50, 146, 130, 130, 
	130, 49, 143, 50, 130, 130, 130, 49, 
	91, 50, 147, 148, 130, 130, 130, 49, 
	91, 50, 149, 130, 130, 130, 49, 91, 
	50, 150, 130, 130, 130, 49, 91, 50, 
	146, 130, 130, 130, 49, 91, 50, 151, 
	130, 130, 130, 49, 91, 50, 152, 130, 
	130, 130, 49, 91, 50, 153, 130, 130, 
	130, 49, 91, 50, 154, 130, 130, 130, 
	49, 91, 50, 155, 130, 130, 130, 49, 
	91, 50, 156, 130, 130, 130, 49, 91, 
	50, 157, 130, 130, 130, 49, 91, 50, 
	146, 130, 130, 130, 49, 91, 50, 158, 
	130, 130, 130, 49, 91, 50, 159, 130, 
	130, 130, 49, 91, 50, 160, 130, 130, 
	130, 49, 91, 50, 161, 130, 130, 130, 
	49, 91, 50, 162, 130, 130, 130, 49, 
	91, 50, 157, 130, 130, 130, 49, 91, 
	50, 163, 130, 130, 130, 49, 91, 50, 
	164, 130, 130, 130, 49, 91, 50, 165, 
	130, 130, 130, 49, 91, 50, 166, 130, 
	130, 130, 49, 91, 50, 157, 130, 130, 
	130, 49, 78, 50, 167, 167, 167, 49, 
	50, 168, 49, 50, 169, 49, 50, 170, 
	49, 50, 171, 172, 49, 50, 173, 49, 
	50, 174, 49, 50, 170, 49, 50, 175, 
	49, 50, 176, 49, 50, 177, 49, 50, 
	178, 49, 50, 179, 49, 50, 180, 49, 
	50, 181, 49, 50, 170, 49, 50, 182, 
	49, 50, 183, 49, 50, 184, 49, 50, 
	185, 49, 50, 186, 49, 50, 181, 49, 
	50, 187, 49, 50, 188, 49, 50, 189, 
	49, 50, 190, 49, 50, 181, 49, 50, 
	191, 49, 50, 192, 49, 50, 193, 49, 
	50, 194, 195, 49, 50, 196, 49, 50, 
	197, 49, 50, 193, 49, 50, 198, 49, 
	50, 199, 49, 50, 200, 49, 50, 201, 
	49, 50, 202, 49, 50, 203, 49, 50, 
	204, 49, 50, 193, 49, 50, 205, 49, 
	50, 206, 49, 50, 207, 49, 50, 208, 
	49, 50, 209, 49, 50, 204, 49, 50, 
	210, 49, 50, 211, 49, 50, 212, 49, 
	50, 213, 49, 50, 204, 49, 39, 214, 
	1, 215, 1, 216, 1, 19, 1, 217, 
	1, 218, 1, 219, 1, 220, 1, 221, 
	1, 19, 1, 222, 223, 1, 18, 1, 
	224, 1, 225, 1, 226, 1, 227, 1, 
	216, 1, 50, 49, 228, 50, 49, 
}

var _mirrorbrain_trans_targs []byte = []byte{
	2, 0, 3, 2, 4, 5, 6, 7, 
	8, 9, 9, 10, 11, 12, 13, 186, 
	189, 195, 14, 15, 16, 17, 17, 18, 
	19, 20, 21, 22, 23, 24, 25, 26, 
	27, 28, 29, 30, 29, 31, 185, 32, 
	33, 34, 35, 34, 35, 36, 37, 38, 
	39, 38, 39, 40, 41, 42, 43, 44, 
	45, 46, 159, 162, 174, 180, 47, 48, 
	49, 50, 51, 52, 53, 133, 136, 148, 
	154, 54, 55, 56, 57, 132, 58, 57, 
	59, 96, 59, 60, 96, 60, 61, 95, 
	62, 63, 94, 64, 63, 65, 66, 67, 
	68, 69, 93, 70, 69, 71, 72, 73, 
	74, 73, 74, 75, 76, 77, 78, 79, 
	80, 92, 81, 82, 83, 202, 203, 84, 
	85, 86, 87, 88, 89, 90, 91, 83, 
	92, 93, 94, 97, 97, 98, 99, 100, 
	101, 102, 103, 105, 109, 121, 127, 104, 
	106, 107, 108, 110, 113, 111, 112, 114, 
	115, 116, 117, 118, 119, 120, 122, 123, 
	124, 125, 126, 128, 129, 130, 131, 132, 
	134, 135, 53, 137, 140, 138, 139, 141, 
	142, 143, 144, 145, 146, 147, 149, 150, 
	151, 152, 153, 155, 156, 157, 158, 160, 
	161, 46, 163, 166, 164, 165, 167, 168, 
	169, 170, 171, 172, 173, 175, 176, 177, 
	178, 179, 181, 182, 183, 184, 185, 187, 
	188, 190, 191, 192, 193, 194, 196, 197, 
	198, 199, 200, 201, 202, 
}

var _mirrorbrain_trans_actions []byte = []byte{
	1, 0, 5, 0, 0, 0, 0, 0, 
	0, 1, 0, 3, 0, 0, 1, 1, 
	1, 1, 0, 0, 7, 1, 0, 9, 
	1, 0, 0, 0, 0, 0, 0, 0, 
	11, 0, 1, 13, 0, 1, 1, 15, 
	0, 1, 39, 0, 17, 0, 0, 1, 
	42, 0, 19, 0, 0, 0, 0, 0, 
	0, 1, 1, 1, 1, 1, 21, 0, 
	0, 0, 0, 0, 1, 1, 1, 1, 
	1, 23, 0, 0, 1, 1, 25, 0, 
	1, 45, 0, 27, 19, 0, 0, 0, 
	0, 1, 1, 29, 0, 0, 0, 0, 
	0, 1, 1, 31, 0, 0, 0, 1, 
	33, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 35, 1, 1, 51, 51, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 27, 0, 0, 0, 0, 
	0, 0, 1, 1, 1, 1, 1, 48, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 37, 
}

const mirrorbrain_start int = 1
const mirrorbrain_first_final int = 202
const mirrorbrain_error int = 0

const mirrorbrain_en_main int = 1


// line 51 "mirrorbrain_parser.rl"

func parseTime(timeStr string) time.Time {
	date, e :=  time.Parse(ApacheDate, timeStr)
	if e != nil {
		log.Fatal("Failed to parse date: ", timeStr)
	}
	return date
}

func mirrorbrain(data string) (entry *Entry, err error) {
	cs, p, pe := 0, 0, len(data)

	marker := 0
	entry = new(Entry)

	
// line 456 "mirrorbrain_parser.go"
	{
	cs = mirrorbrain_start
	}

// line 461 "mirrorbrain_parser.go"
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
// line 67 "mirrorbrain_parser.rl"

 marker = p 
		case 1:
// line 69 "mirrorbrain_parser.rl"

 entry.Date = parseTime(data[marker:p]) 
		case 2:
// line 70 "mirrorbrain_parser.rl"

 entry.Ip = data[marker:p] 
		case 3:
// line 71 "mirrorbrain_parser.rl"

 entry.RequestMethod = data[marker:p] 
		case 4:
// line 72 "mirrorbrain_parser.rl"

 entry.RequestPath = data[marker:p] 
		case 5:
// line 73 "mirrorbrain_parser.rl"

 entry.RequestProto = data[marker:p] 
		case 6:
// line 74 "mirrorbrain_parser.rl"

 entry.ReturnCode = data[marker:p] 
		case 7:
// line 75 "mirrorbrain_parser.rl"

 entry.Size = data[marker:p] 
		case 8:
// line 76 "mirrorbrain_parser.rl"

 entry.Referer = data[marker:p] 
		case 9:
// line 77 "mirrorbrain_parser.rl"

 entry.Agent = data[marker:p] 
		case 10:
// line 78 "mirrorbrain_parser.rl"

 entry.RequestType = data[marker:p] 
		case 11:
// line 79 "mirrorbrain_parser.rl"

 entry.GivenType = data[marker:p] 
		case 12:
// line 80 "mirrorbrain_parser.rl"

 entry.Region = data[marker:p] 
		case 13:
// line 81 "mirrorbrain_parser.rl"

 entry.Mirror = data[marker:p] 
		case 14:
// line 82 "mirrorbrain_parser.rl"

 entry.Country = data[marker:p] 
		case 15:
// line 83 "mirrorbrain_parser.rl"

 entry.Asn = data[marker:p] 
		case 16:
// line 84 "mirrorbrain_parser.rl"

 entry.Net = data[marker:p] 
		case 17:
// line 85 "mirrorbrain_parser.rl"

 entry.RedirSize = data[marker:p] 
		case 18:
// line 86 "mirrorbrain_parser.rl"

 entry.RedirRange = data[marker:p] 
		case 19:
// line 88 "mirrorbrain_parser.rl"

 entry.Parsed = true 
// line 620 "mirrorbrain_parser.go"
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

// line 129 "mirrorbrain_parser.rl"


	return entry, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var sizeByMirror =  make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		entry, err := mirrorbrain(line)
		if err != nil {
			log.Fatal(err)
		}
		if !entry.Parsed {
			fmt.Println("IP:", entry.Ip)
			fmt.Println("Date:", entry.Date)
			fmt.Println("RequestMethod:", entry.RequestMethod)
			fmt.Println("RequestPath:", entry.RequestPath)
			fmt.Println("RequestProto:", entry.RequestProto)
			fmt.Println("ReturnCode:", entry.ReturnCode)
			fmt.Println("Size:", entry.Size)
			fmt.Println("Referer:", entry.Referer)
			fmt.Println("Agent:", entry.Agent)
			fmt.Println("RequestType:", entry.RequestType)
			fmt.Println("GivenType:", entry.GivenType)
			fmt.Println("Region", entry.Region)
			fmt.Println("Mirror:", entry.Mirror)
			fmt.Println("Country:", entry.Country)
			fmt.Println("Asn:", entry.Asn)
			fmt.Println("Net:", entry.Net)
			fmt.Println("RedirSize:", entry.RedirSize)
			fmt.Println("RedirRange:", entry.RedirRange)

			log.Fatal("Failed to parse:\n", line)
		}

		if matches(*entry) {
			size, err := strconv.Atoi(entry.RedirSize)
			if err != nil {
				log.Print(entry)
				log.Fatal("Failed to parse file size: ", entry.RedirSize)
			}
			sizeByMirror[entry.Mirror] += size
		}
	}

	for key, value := range sizeByMirror {
		fmt.Println("Mirror:", key, "Size:", value/1024/1024/1024, "gb")
	}
}

func matches(entry Entry) (bool) {
	return  entry.RequestMethod == "GET" &&
		entry.GivenType == "redirect" &&
		//entry.RedirSize != "" &&
		strings.Index(entry.RequestPath, PathFilter) > -1
}

