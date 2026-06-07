ALTER TABLE namespaces
    ADD CONSTRAINT namespaces_name_unique UNIQUE (name);