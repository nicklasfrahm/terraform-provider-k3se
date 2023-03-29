default: install

generate:
	go generate ./...

install:
	go install .

test:
	go test -count=1 -parallel=4 ./...

# Install a temporary override for the provider installation source.
plan: install
	@mv ~/.terraformrc ~/.terraformrc.bak 2> /dev/null || true
	@envsubst < config/.terraformrc > ~/.terraformrc
	terraform -chdir=examples/standalone plan
	terraform -chdir=examples/ha plan
	@rm ~/.terraformrc
	@mv ~/.terraformrc.bak ~/.terraformrc 2> /dev/null || true

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
