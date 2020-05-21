CREATE TABLE TODO (
    id BIGINT NOT NULL AUTO_INCREMENT,
    content VARCHAR (255) DEFAULT "",
    PRIMARY KEY (id)
);

INSERT into TODO (id, content) VALUES (1, "a content 1");
INSERT into TODO (id, content) VALUES (2, "a content 2");
INSERT into TODO (id, content) VALUES (3, "a content 3");
INSERT into TODO (id, content) VALUES (4, "a content 4");