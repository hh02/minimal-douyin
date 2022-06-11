gen_path = ../
package_name = response
import_path = github.com/hh02/minimal-douyin
idl_file = response.proto

response:
	rm -rf $(gen_path)/kitex_gen/$(package_name) && cd $(gen_path) && kitex -module $(import_path) -I ./idl/ -type protobuf $(idl_file)