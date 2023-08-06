CREATE TABLE todos
(
    id bigint NOT NULL AUTO_INCREMENT,
    description varchar(255) NOT NULL,
    is_completed tinyint DEFAULT false,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);
