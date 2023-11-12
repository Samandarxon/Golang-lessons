CREATE OR REPLACE PROCEDURE task_3_2(sender_id_arg INT ,receiver_id_arg INT,price INT) LANGUAGE PLPGSQL
AS 
$$
    DECLARE
        sender_v RECORD;
        receiver_v RECORD;
        commission_price_v DECIMAL = price*0.5/100;
        cashback_PR DECIMAL = 0;
        max_price INT;
        is_module_5 BOOLEAN;
        history_cashback_pr_id RECORD;
    BEGIN 
        SELECT * FROM users WHERE id = sender_id_arg INTO sender_v;
        SELECT * FROM users WHERE id = receiver_id_arg INTO receiver_v;
        RAISE INFO '%   %',sender_v,receiver_v;
        SELECT ((SELECT COUNT(sender_id) FROM history WHERE sender_id = sender_id_arg ) +1)  % 5 = 0 INTO is_module_5;
        RAISE INFO ' IS %  ',is_module_5;

    IF is_module_5 THEN 
        RAISE INFO '>>>>>>>>>>>>>>>>>>>>>>>>>>>>salom<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<';
        
        INSERT INTO history(sender_id,send_money,commission_price,cashback) VALUES
        (sender_id_arg,price,commission_price_v,cashback_PR) RETURNING id INTO history_cashback_pr_id ;

        RAISE INFO 'history_cashback_pr_id % ',history_cashback_pr_id;


        SELECT MAX(send_money) FROM (SELECT * FROM history WHERE sender_id = sender_id_arg ORDER BY id DESC LIMIT 5) INTO max_price;
        RAISE INFO 'max_price % CASH_PR %',max_price,cashback_PR;

        sender_v.balance = sender_v.balance - price + commission_price_v::INT + max_price * 0.25 / 100;
        receiver_v.balance = receiver_v.balance + price;
        cashback_PR = max_price * 0.25 / 100;
        RAISE INFO 'SENDER % receiver_v %',sender_v,receiver_v;


        UPDATE users SET balance = sender_v.balance::INT WHERE id = sender_id_arg;
        UPDATE users SET balance = receiver_v.balance::INT WHERE id = receiver_id_arg;
        UPDATE history SET cashback=cashback_PR WHERE id = history_cashback_pr_id.id;

        
        INSERT INTO checks(sender_name , receiver_name , sender_price , commission_price , cashback ) VALUES
        (sender_v.name,receiver_v.name,price,commission_price_v::INT,cashback_PR);
        RAISE INFO '>>>>>>>>>>>>>>>>>>>>>>>>>>>>salom<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<';
    ELSE 
        
        sender_v.balance = sender_v.balance - price - commission_price_v;
        receiver_v.balance = receiver_v.balance + price;

        UPDATE users SET balance = sender_v.balance::INT WHERE id = sender_id_arg;
        UPDATE users SET balance = receiver_v.balance::INT WHERE id = receiver_id_arg;

        INSERT INTO history(sender_id,send_money,commission_price,cashback) VALUES
        (sender_id_arg,price,commission_price_v,cashback_PR);

        INSERT INTO checks(sender_name , receiver_name , sender_price , commission_price , cashback ) VALUES
        (sender_v.name,receiver_v.name,price,commission_price_v::INT,cashback_PR);
    END IF ;
    END;
$$;


