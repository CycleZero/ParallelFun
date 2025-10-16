CREATE TABLE servers (
     id BIGSERIAL PRIMARY KEY,
     created_at TIMESTAMP WITH TIME ZONE,
     updated_at TIMESTAMP WITH TIME ZONE,
     deleted_at TIMESTAMP WITH TIME ZONE,
     name VARCHAR(255) NOT NULL,
     address VARCHAR(255) NOT NULL,
     port INTEGER NOT NULL,
     status INTEGER NOT NULL,
     owner_id BIGINT NOT NULL
);

-- 创建索引
CREATE INDEX idx_servers_deleted_at ON servers(deleted_at);
CREATE INDEX idx_servers_name ON servers(name);
CREATE INDEX idx_servers_status ON servers(status);
CREATE INDEX idx_servers_owner_id ON servers(owner_id);
