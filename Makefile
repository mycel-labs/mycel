mock-expected-keepers:
	mockgen -source=x/registry/types/expected_keepers.go \
        -package testutil \
        -destination=x/registry/testutil/expected_keepers_mocks.go 
	mockgen -source=x/incentives/types/expected_keepers.go \
        -package testutil \
        -destination=x/incentives/testutil/expected_keepers_mocks.go 
