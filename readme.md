

https://blog.filippo.io/building-python-modules-with-go-1-5/

http://blog.golang.org/c-go-cgo

apt-get install python3-dev

go build -buildmode=c-shared -o makefiles.so makefiles.go makefiles_declarations.go 

time ./makefiles_plain_python.py 
time ./makefiles_go.py 
