tailregexp
==========

install
 
	$ cd tailregexp
	$ go install ./...
use

	$ bin/tailregexp 
			-file='/path/to/logfile.log' -regex='(.*)ERROR(.*)'
get help
 
	$ bin/tailregexp -help