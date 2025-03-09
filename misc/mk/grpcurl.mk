
# Manual tests on gRPC endpoints

# While fields can be declared optional in Proto3 (experimental feature),
# all fields are essentially optional in the sense that missing fields are accepted,
# and are assigned the default value. When an optional field is missing, it is
# assigned a null pointer value.

# Nominal - required fields only
test-nominal-1:
	grpcurl \
    -import-path api/proto \
    -proto service.proto \
    -plaintext -d '{"user_id":"1", "category": "food", "amount": 42.5}' \
    localhost:5000 \
    expense.Expenses/CreateExpense

# Nominal - all fields
test-nominal-2:
	grpcurl \
    -import-path api/proto \
    -proto service.proto \
    -plaintext -d '{"user_id":"1", "category": "food", "amount": 42.5, "description": "breakfast", "timestamp": "2025-03-08T15:04:05Z"}' \
    localhost:5000 \
    expense.Expenses/CreateExpense

# Error - category field is an empty string
test-error-1:
	grpcurl \
    -import-path api/proto \
    -proto service.proto \
    -plaintext -d '{"user_id":"1", "category": ""}' \
    localhost:5000 \
    expense.Expenses/CreateExpense

# Error - amount field is less than 0
test-error-2:
	grpcurl \
    -import-path api/proto \
    -proto service.proto \
    -plaintext -d '{"user_id":"1", "category": "breakfast", "amount": -2}' \
    localhost:5000 \
    expense.Expenses/CreateExpense

# Error - timestamp is not ISO 8601
test-error-3:
	grpcurl \
    -import-path api/proto \
    -proto service.proto \
    -plaintext -d '{"user_id":"1", "category": "breakfast", "amount": 22.4, "timestamp": "January 2"}' \
    localhost:5000 \
    expense.Expenses/CreateExpense