# checksum-tool

## Build
```bash
git clone PROJECT
cd PROJECT/
go mod tidy
go build -o checksum .
./checksum -f FILE
# output like below:
CRC64 checksum: 2614761567380709856
MD5 checksum: 75938c249ac2ce60f4c1f79f3f89dd83
SHA1 checksum: d148e4ba0859306baded254739820c5f5df234b7
SHA256 checksum: 13bd032ffe2d166ae842e2627d27a1a9b996b7195af13c769ee621fe408fbab4
Spent time: 2.751002291s
```
