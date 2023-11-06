CREATE TABLE users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE roles(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE users_roles(
    user_id UUID,
    role_id UUID,
    PRIMARY KEY(user_id, role_id),
    CONSTRAINT fk_users_roles_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_users_roles_role FOREIGN KEY(role_id) REFERENCES roles(id)
);

CREATE TYPE permission AS ENUM ('create', 'read', 'update', 'delete');

CREATE TABLE permissions(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    role_id UUID,
    type permission NOT NULL,
    CONSTRAINT fk_permissons_role FOREIGN KEY(role_id) REFERENCES roles(id),
    UNIQUE (role_id, type)
);
