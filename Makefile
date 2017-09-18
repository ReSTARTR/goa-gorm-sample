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

run:
	go run *.go

clean:
	rm -rf app client controller models swagger tool
