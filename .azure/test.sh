go test -coverprofile=coverage.txt -covermode count 2>&1 | tee testlogs.txt
go-junit-report < testlogs.txt > report.xml
go tool cover -html=coverage.txt -o cover.html