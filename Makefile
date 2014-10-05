test:
	go test && \
	cd zoopla && go test && cd .. && \
	cd rightmove && go test && cd ..

.SILENT:
.PHONY: test
