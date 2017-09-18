GOA=github.com/goadesign
REPOS=github.com/ReSTARTR/goa-sample
GOAGEN=goagen --design $(REPOS)/design
APPPKG=--app-pkg $(REPOS)/app

all: app client controller gorma swagger tool

app:
	$(GOAGEN) app

client:
	$(GOAGEN) client

controller:
	$(GOAGEN) $(APPPKG) controller \
		--pkg controller \
		--out controller \
		--force

swagger:
	$(GOAGEN) swagger

tool:
	$(GOAGEN) $(APPPKG) tool

gorma:
	$(GOAGEN) gen \
		--pkg-path=$(GOA)/gorma

swagger-ui:
	VERSION=3.2.1
	curl -L -o /tmp/swagger-ui-${VERSION}.tar.gz https://github.com/swagger-api/swagger-ui/archive/v${VERSION}.tar.gz
	tar zxf /tmp/swagger-ui-${VERSION}.tar.gz
	mv ./swagger-ui-${VERSION} swaggerui

run:
	go run *.go

clean:
	rm -rf app client controller models swagger tool
