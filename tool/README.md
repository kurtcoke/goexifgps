cmdtool:
-------

$ bash loop > locations.txt

Will print GPS locations to text file and errors to console.

$ bash loop 2> errorlog.txt

Print errors to text file and GPS locations to console.


Perhaps use a more descriptive name when you "go install" it on your own machine.

Install as binary:
-----------------

$ go install cmdtool.go

$ sudo cp $GOBIN/cmdtool /usr/bin
