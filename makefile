# For full Kind v0.17 release notes: https://github.com/kubernetes-sigs/kind/releases/tag/v0.17.0
#
# Other commands to install.
# go install github.com/divan/expvarmon@latest
# go install github.com/rakyll/hey@latest
#
# http://cart-service.cart-system.svc.cluster.local:4000/debug/pprof
# curl -il cart-service.cart-system.svc.cluster.local:4000/debug/vars
# curl -il cart-service.cart-system.svc.cluster.local:3000/status
#
# RSA Keys
# 	To generate a private/public key PEM file.
# 	$ openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# 	$ openssl rsa -pubout -in private.pem -out public.pem
#
# Testing Coverage
# 	$ go test -coverprofile p.out
# 	$ go tool cover -html p.out

db:
	go run app/scratch/db/main.go

run:
	go run app/services/app-api/main.go

status-debug:
	curl -sS localhost:4000/debug/liveness

status-api:
	curl -il localhost:3000/status

pgcli:
	pgcli postgresql://postgres:postgres@database-service.cart-system.svc.cluster.local

auth-local:
	curl -il -H "Authorization: Bearer ${TOKEN}" localhost:3000/auth

tidy:
	go mod tidy
	go mod vendor

# export TOKEN="COPY TOKEN STRING FROM LAST CALL"

test-users-local:
	curl -il -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/users/1/2

test-users:
	curl -il -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/users/1/2

# ==============================================================================
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	CGO_ENABLED=0 go test -count=1 ./...
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...
	govulncheck ./...
