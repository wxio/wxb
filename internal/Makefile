

goprotoc:
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/build/build.proto 
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/workspace/workspace.proto 
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/semantics/semantics.proto 
