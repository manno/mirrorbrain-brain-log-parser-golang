# vim: ft=ragel
require 'ostruct'
require 'yaml'

#LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\" \        
#            want:%{WANT}e give:%{GIVE}e r:%{MB_REALM}e %{X-MirrorBrain-Mirror}o \
#            %{MB_CONTINENT_CODE}e:%{MB_COUNTRY_CODE}e ASN:%{ASN}e P:%{PFX}e \    
#            size:%{MB_FILESIZE}e %{Range}i" combined_redirect                    
#
#
#192.168.122.1 - - [29/Dec/2013:18:54:41 +0100] "GET /congress/2013/mp4/30c3-5415-de-en-Der_tiefe_Staat_h264-iprod.mp4.torrent HTTP/1.0" 200 11668 "-" "Transmission/2.77"     want:torrent give:torrent r:- -     EU:DE ASN:- P:-     size:126539847 -

$_VERBOSE=true
$DEBUG=false

%%{
  machine mirrorbrain;

  # action
  action mark { marker = p }

  action setDate { @result.date = data[marker..p-1] }
  action setIP { @result.ip = data[marker..p-1] }
  action setRequestMethod { @result.request_method = data[marker..p-1] }
  action setRequestPath { @result.request_path = data[marker..p-1] }
  action setRequestProto { @result.request_proto = data[marker..p-1] }
  action setReturn { @result.return_code = data[marker..p-1] }
  action setSize { @result.size = data[marker..p-1] }
  action setReferer { @result.referer = data[marker..p-1] }
  action setAgent { @result.agent = data[marker..p-1] }
  action setRequestType { @result.request_type = data[marker..p-1] }
  action setGivenType { @result.given_type = data[marker..p-1] }
  action setRegion { @result.region = data[marker..p-1] }
  action setMirror { @result.mirror = data[marker..p-1] }
  action setCountry { @result.country = data[marker..p-1] }
  action setASN { @result.asn = data[marker..p-1] }
  action setNet { @result.net = data[marker..p-1] }
  action setRedirSize { @result.redir_size = data[marker..p-1] }
  action setRedirBytes { @result.redir_bytes = data[marker..p-1] }

  action LogLineFinished {
    @result.parsed = true
  }

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

}%%

module Mirrorbrain
  class Tokenizer

    def initialize
      @result = OpenStruct.new
      @result.parsed = false

      %% write data;
    end

    def run(data)
      %% write init;
      %% write exec;

      @result
    end

  end
end

CONG = /\/congress\/2013\//

tokenizer = Mirrorbrain::Tokenizer.new
sums = Hash.new { |h,k| h[k] = 0 }
while gets do
  r = tokenizer.run($_)
  next unless r.request_method == 'GET'
  next unless r.given_type == 'redirect'
  next unless CONG.match r.request_path
  #puts [r.mirror, r.redir_size].join(',')
  sums[r.mirror] += r.redir_size.to_i
end

sums.each { |m,s| 
  printf "%s = %d mb\n", m, s/1024/1024
}
