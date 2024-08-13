swag:
	@echo 'Generating swagger'
	swag init --parseDependency --parseInternal --generalInfo cmd/app/main.go
