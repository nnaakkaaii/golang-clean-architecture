PS=`find ./app | grep *.go`
for P in $PS
do
  go fmt $P
done