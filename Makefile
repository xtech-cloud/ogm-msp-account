.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/account/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/account/auth.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/account/profile.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/account/query.proto
