# scrape-file

This script searches a directory for lines in files matching a certain pattern prefix.
It omits the prefix match and saves the match to a file of your choosing.
```
usage: go run scrape.go <directory> <keyword> <output_file>
example: go run scrape.go /tmp/test/ /tmp/test/output
```
