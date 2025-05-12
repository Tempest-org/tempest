CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY, 
    name varchar(255) NOT NULL,
    description text,
    website_url varchar(255),
    primary_color varchar(255),
    secondary_color varchar(255),
    logo_url varchar(255),
    background_url varchar(255),
    created_by varchar(255) NOT NULL, 
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL
);
CREATE UNIQUE INDEX organizations_name_index ON organizations (name);

CREATE TABLE IF NOT EXISTS organizations_members (
    id SERIAL PRIMARY KEY,
    organization_id int NOT NULL,
    user_id varchar(255) NOT NULL, 
    role varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE
);
CREATE INDEX organizations_members_role ON organizations_members (role);
CREATE UNIQUE INDEX organizations_members_org_user ON organizations_members (organization_id, user_id);

CREATE TABLE IF NOT EXISTS organizations_invitations (
    id SERIAL PRIMARY KEY,
    organization_id int NOT NULL,
    email varchar(255) NOT NULL,
    invited_by_user_id varchar(255) NOT NULL, 
    invited_as_role varchar(255) NOT NULL,
    accepted boolean NOT NULL DEFAULT false,
    responded_at timestamp DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,
    FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE
);
CREATE INDEX organizations_invitations_invited_by_user_id_index ON organizations_invitations (invited_by_user_id);
CREATE UNIQUE INDEX organizations_invitations_org_email ON organizations_invitations (organization_id, email);
