gen_path = ../..
package_name = userrpc
service_name = userservice
import_path = github.com/hh02/minimal-douyin
idl_file = user.proto

server:
	rm -rf ./script ./build.sh ./handler.go ./main.go && kitex -module $(import_path) -I ../../idl/ -type protobuf -service $(service_name) -use $(import_path)/kitex_gen $(idl_file)

client:
	rm -rf $(gen_path)/kitex_gen/$(package_name) && cd $(gen_path) && kitex -module $(import_path) -I ./idl/ -type protobuf $(idl_file)