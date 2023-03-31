test-all-keepers:
	go test ./x/.../keeper
test-all-types:
	go test ./x/.../types
test-all-keepers-types:
	go test ./x/.../keeper
	go test ./x/.../types
test-module-%:
	go test ./x/$*/keeper
	go test ./x/$*/types
