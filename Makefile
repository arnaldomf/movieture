BUILDER=go build
INSTALLER=go install
REPOSITORY=github.com/arnaldomf/moviture

movieture:
	$(BUILDER) $(REPOSITORY)

install: movieture
	$(INSTALLER) $(REPOSITORY)
