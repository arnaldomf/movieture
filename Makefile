BUILDER=go build
INSTALLER=go install
REPOSITORY=github.com/arnaldomf/movieture

movieture: main.go
	$(BUILDER) $(REPOSITORY)

install: movieture
	$(INSTALLER) $(REPOSITORY)
