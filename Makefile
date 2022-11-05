gen:
	rm -f ./pb/upload/*.go \
    rm -f ./doc/*.swagger.json
	\
    protoc -I ./proto \
    --go_out ./pb/ --go_opt paths=source_relative \
    --go-grpc_out ./pb/ --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./pb/ --grpc-gateway_opt paths=source_relative \
    --openapiv2_out ./doc --openapiv2_opt=allow_merge=true,merge_file_name=weather \
    ./proto/upload.proto

clean:
	rm -f ./pb/*.pb.go

d-up:
	docker run --name=redis-sb -p 6379:6379 -d --rm redis

d-stop:
	docker stop redis-sb

d-exec:
	docker exec -it redis-sb /bin/bash
