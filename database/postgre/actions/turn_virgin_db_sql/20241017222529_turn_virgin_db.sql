-- +goose Up
-- +goose StatementBegin
DO $$
DECLARE
    r RECORD;
BEGIN
    -- Para cada tabela no schema public, execute o comando DROP TABLE, exceto para a tabela 'goose_db_version'
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        IF r.tablename != 'goose_db_version' THEN
            EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
        END IF;
    END LOOP;
END $$;


-- +goose StatementEnd
