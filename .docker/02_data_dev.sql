INSERT INTO account (`id`,`name`) VALUES (123, 'Account 1');
INSERT INTO category (`name`,`is_fallback`) VALUES ('FOOD', false);
INSERT INTO category (`name`,`is_fallback`) VALUES ('MEAL', false);
INSERT INTO category (`name`,`is_fallback`) VALUES ('CASH', true);
INSERT INTO mcc (`mcc`,`category_id`) VALUES (5411, 1);
INSERT INTO mcc (`mcc`,`category_id`) VALUES (5412, 1);
INSERT INTO mcc (`mcc`,`category_id`) VALUES (5811, 2);
INSERT INTO mcc (`mcc`,`category_id`) VALUES (5812, 2);
INSERT INTO balance (`account_id`,`amount`,`category_id`) VALUES (123, 100, 1);
INSERT INTO balance (`account_id`,`amount`,`category_id`) VALUES (123, 100, 2);
INSERT INTO balance (`account_id`,`amount`,`category_id`) VALUES (123, 100, 3);

select * from balance;
