INSERT INTO user (NicName, Email, Password) VALUES ("qwer", "ASD", "z");


INSERT INTO posts(title, content, create_at, update_at, id_user) VALUES("Hello", "World! and more history", "123.112.34", "123.123.1", 1);

INSERT INTO tags(title) VALUES("News");

INSERT INTO tag_posts(id_tag, id_post) VALUES(1, 1);

INSERT INTO categories(title, description) VALUES("hello", "World!!!!!!");

INSERT INTO category_posts(id_category, id_post) VALUES(1,1);

INSERT INTO comments(text, id_user, id_post) VALUES("Такую хрень сотворил", 1, 1);

INSERT INTO likes_posts(id_post, id_user) VALUES(1, 1);