package_name = followrpc
service_name = followservice
import_path = github.com/hh02/minimal-douyin

server:
	rm -rf ./script ./build.sh ./handler.go ./main.go && kitex -module $(import_path) -I ../../idl/ -type protobuf -service $(service_name) -use $(import_path) follow.proto

client:
	rm -rf $(gen_path)/$(package_name) && cd ../../ && kitex -module $(import_path) -I ./idl/ -type protobuf follow.proto