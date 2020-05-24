rm -rf build/conf
cp -r conf build
go build -o build\zeus.exe ./cmd\zeus\