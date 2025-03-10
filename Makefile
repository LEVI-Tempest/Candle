ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	PROTO_FILES=$(shell $(Git_Bash) -c "find pkg/proto -name *.proto")
else
	PROTO_FILES=$(shell find pkg/proto -name *.proto)
endif


config:
	@echo "config"
	protoc --proto_path=./third_party \
           --proto_path=./pkg/proto \
           --go_out=paths=source_relative:./pkg/proto \
		   --go-http_out=paths=source_relative:./pkg/proto \
		   --go-grpc_out=paths=source_relative:./pkg/proto \
		   $(PROTO_FILES)
#		   --openapi_out=fq_schema_naming=true,default_response=false,paths=source_relative:./pkg/protoc/api/openapi \
		   $(PROTO_FILES)