-- +goose Up
-- +goose StatementBegin
INSERT INTO "spender" (name, email) VALUES ('John Doe', 'jonhdoe@gmail.com');
INSERT INTO "spender" (name, email) VALUES ('Jane Doe', 'janedoe@gmail.com');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "spender" WHERE  email = 'jonhdoe@gmail.com';
DELETE FROM "spender" WHERE  email = 'janedoe@gmail.com';
-- +goose StatementEnd
