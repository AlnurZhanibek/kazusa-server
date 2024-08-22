swag:
	@echo 'Generating swagger'
	swag init --parseDependency --parseInternal --generalInfo main.go
