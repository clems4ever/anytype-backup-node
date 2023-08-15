
all:
	goreleaser build --snapshot

clean:
	rm -rf dist