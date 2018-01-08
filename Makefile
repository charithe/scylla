PARSER_PACKAGE:=graphql/internal
GENERATED_PARSER:=graphql/internal/generated_parser.go

.PHONY: run test clean

run: $(GENERATED_PARSER)
	@go run main.go

test: $(GENERATED_PARSER)
	@go test ./...

t: $(GENERATED_PARSER)
	@go test -run '/error_cases/invalid_int' ./...

$(GENERATED_PARSER):
	@go generate ./$(PARSER_PACKAGE)

clean:
	-@rm $(GENERATED_PARSER)
