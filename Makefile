default: generate

generate:
	swagger generate client -f spec.json -A registry

release:
	goreleaser