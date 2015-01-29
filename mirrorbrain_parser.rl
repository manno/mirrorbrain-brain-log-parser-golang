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

%% machine mirrorbrain;
%% write data;

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

	%%{
		action mark { marker = p }

		action setDate { entry.Date = parseTime(data[marker:p]) }
		action setIP { entry.Ip = data[marker:p] }
		action setRequestMethod { entry.RequestMethod = data[marker:p] }
		action setRequestPath { entry.RequestPath = data[marker:p] }
		action setRequestProto { entry.RequestProto = data[marker:p] }
		action setReturn { entry.ReturnCode = data[marker:p] }
		action setSize { entry.Size = data[marker:p] }
		action setReferer { entry.Referer = data[marker:p] }
		action setAgent { entry.Agent = data[marker:p] }
		action setRequestType { entry.RequestType = data[marker:p] }
		action setGivenType { entry.GivenType = data[marker:p] }
		action setRegion { entry.Region = data[marker:p] }
		action setMirror { entry.Mirror = data[marker:p] }
		action setCountry { entry.Country = data[marker:p] }
		action setASN { entry.Asn = data[marker:p] }
		action setNet { entry.Net = data[marker:p] }
		action setRedirSize { entry.RedirSize = data[marker:p] }
		action setRedirBytes { entry.RedirBytes = data[marker:p] }

		action LogLineFinished { entry.Parsed = true }

		ws = ' ';
		ws0 = ' '{0,};
		eol = /[\r\n]/ | '\r\n';
		date = [^\]]+                                                      >mark %setDate;
		request_method = ( 'GET' | 'POST' | 'HEAD' )                       >mark %setRequestMethod;
		request_path = [^ ]+                                               >mark %setRequestPath;
		request_proto = ( 'HTTP/1.0' | 'HTTP/1.1' )                        >mark %setRequestProto;
		return = digit+                                                    >mark %setReturn;
		size = digit+                                                      >mark %setSize;
		referer = [^"]+                                                    >mark %setReferer;
		useragent = [^"]+                                                  >mark %setAgent;
		request_type = ( 'file' | 'torrent' | 'redirect' | '-' )           >mark %setRequestType;
		give = ( 'file' | 'torrent' | 'redirect' | '-' )                   >mark %setGivenType;
		optional = ( alnum+ | '-' );
		region = optional                                                  >mark %setRegion;
		mirror = [^ ]+                                                     >mark %setMirror;
		country = optional ":" optional                                    >mark %setCountry;
		asn = 'ASN:' optional                                              >mark %setASN;
		ip = [0-9\.]+                                                      >mark %setIP;
		net = [0-9\.\/]+                                                   >mark %setNet; 
		redir_size = ( digit+ | '-' )                                      >mark %setRedirSize;
		redir_bytes = ( '-' | 'bytes=' [0-9\-]+ )                          >mark %setRedirBytes;

		log_line = ip ws '-' ws '-' ws '[' date ']' ws 
		'"' request_method ws request_path ws request_proto '"'
		ws return ws size ws '"' referer '"' ws 
		'"' useragent '"' ws ws0 
		'want:' request_type ws 'give:' give ws
		'r:' region ws mirror ws ws0 country ws asn ws 
		'P:' net ws ws0 
		'size:' redir_size ws
		redir_bytes eol @LogLineFinished;

		main := log_line;

		write init;
		write exec;

	}%%

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
