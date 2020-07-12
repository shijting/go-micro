protoc --proto_path=src/protos --micro_out=src/users --go_out=src/users Users.proto
protoc --proto_path=src/protos --micro_out=src/courses --go_out=src/courses Courses.proto
protoc-go-inject-tag -input=src/courses/Courses.pb.go