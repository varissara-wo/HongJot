.PHONY: new-migration

new-migration:
	@echo "$(name)" | grep -q -E '^[0-9]{3}_[a-z_]+$$' || (echo "Invalid name format. Should be 00X_name_with_underscore" && exit 1)
	@echo "Creating a new migration $(name)..."
	@cp ./migration/000_example.sql.template ./migration/$(name).sql