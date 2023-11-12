CREATE OR REPLACE PROCEDURE payment( sender_id INT ,receiver_id INT ,price INT)LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        sender RECORD ;
        receiver RECORD;
        foiz DECIMAL = 0.5;
    BEGIN 
        SELECT * FROM users INTO sender WHERE id = sender_id;
        SELECT * FROM users INTO receiver WHERE id = receiver_id;

        RAISE INFO 'receiver % sender %',receiver.id,sender.id;
        
        IF sender IS NULL OR receiver IS NULL THEN
        RAISE EXCEPTION 'XATO';
            ROLLBACK;
        END IF;
        sender.balance = sender.balance - price - price*foiz/100;
        receiver.balance = receiver.balance + price;

        UPDATE users SET balance = sender.balance::INT  WHERE id = sender_id;
        UPDATE users SET balance = receiver.balance::INT  WHERE id = receiver_id;
        
        INSERT INTO click (sender_id, receiver_id,price,komissiya)VALUES
        (sender_id,receiver_id,price,price*foiz/100);
        

        IF sender.balance >= 0  THEN
            COMMIT;
        else 
            ROLLBACK;
            raise notice '%da % buncha pul yo''q ',sender.name,price;
         END IF;
        
        IF price+price*foiz/100 < sender.balance THEN
            COMMIT;
        else 
            raise notice '% sizda % cha pul bo''lishi kerak biz sizdan % shuncha foiz olamiz',sender.name,price,foiz;
            ROLLBACK;
        END IF;
    END;
$$;


CALL payment(1, 2, 10000)

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    balance NUMERIC
);

INSERT INTO users(name, balance)
VALUES
('Baxodir', 100000),
('Izzat', 50000);






2. Komissiya pul yechib olasiz. 0.5%

100_000 -> 100%
 ?      -> 5%
500 som;

click
    id
    sender_id
    receiver_id
    price
    komissiya
    created_at
CREATE TABLE click(
     id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id),
    receiver_id INT NOT NULL REFERENCES users(id),
    price INT,
    komissiya INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
--    3.1 Har 5 marta otkazmada CACHBACK 0.25 qaytaring. 

CREATE OR REPLACE PROCEDURE task_3_1(sender_id_arg INT, receiver_id_arg INT,price INT)LANGUAGE PLPGSQL
AS  
$$
    DECLARE 
        sender RECORD ;
        receiver RECORD;
        foiz DECIMAL = 0.5;
    BEGIN 
        SELECT * FROM users INTO sender WHERE id = sender_id_arg;
        SELECT * FROM users INTO receiver WHERE id = receiver_id_arg;

        RAISE INFO 'receiver % sender %',receiver.id,sender.id;
        
        IF sender IS NULL OR receiver IS NULL THEN
        RAISE EXCEPTION 'XATO';
            ROLLBACK;
        END IF;
        sender.balance = sender.balance - price - price*foiz/100;
        receiver.balance = receiver.balance + price;

        UPDATE users SET balance = sender.balance::INT  WHERE id = sender_id_arg;
        UPDATE users SET balance = receiver.balance::INT  WHERE id = receiver_id_arg;
        
        INSERT INTO click (sender_id, receiver_id,price,komissiya)VALUES
        (sender_id_arg,receiver_id_arg,price,price*foiz/100);
        

        IF sender.balance >= 0  THEN
            COMMIT;
        else 
            ROLLBACK;
            raise notice '%da % buncha pul yo''q ',sender.name,price;
         END IF;
        
        IF price+price*foiz/100 < sender.balance THEN
            COMMIT;
        else 
            raise notice '% sizda % cha pul bo''lishi kerak biz sizdan % shuncha foiz olamiz',sender.name,price,foiz;
            ROLLBACK;
        END IF;
    END;
$$;

CREATE OR REPLACE PROCEDURE py(sender_id_arg INT,receiver_id_arg INT ,price INT)LANGUAGE PLPGSQL
AS 
$$
    DECLARE 
        sender RECORD;
        receiver RECORD;
        cash_VR RECORD;
        cash_count int = 1 ;
        money_sender_VR RECORD;
        percent DECIMAL = 0.5/100;
        percent_cashback DECIMAL = 0.25/100;
        MAX_PRICE INT;

    BEGIN
        SELECT * FROM users WHERE id = sender_id_arg INTO sender;
        SELECT * FROM users WHERE id = receiver_id_arg INTO receiver;

        INSERT INTO money_sender(money,cash_sender_id)VALUES (price,sender_id_arg) 
        ON CONFLICT(id) DO NOTHING;
        SELECT * FROM  money_sender AS M  WHERE M.cash_sender_id = sender_id_arg INTO money_sender_VR ORDER BY id DESC;
        
        raise info 'money_sender % ',money_sender_VR;

        INSERT INTO cash(sender_id,count,money_id)VALUES (sender_id_arg,cash_count,money_sender_VR.id) 
        ON CONFLICT(sender_id) DO NOTHING;
        SELECT * FROM cash AS C WHERE C.sender_id = sender_id_arg INTO cash_VR;



        raise info 'sender % receiver % MONEY % CASH %' ,sender,receiver,money_sender_VR,cash_VR;


        IF cash_VR.count<5 THEN
        
        sender.balance = sender.balance - price - price * percent;
        receiver.balance = receiver.balance + price;
        cash_VR.count = cash_VR.count + 1;
        cash_count = cash_VR.count;
        UPDATE users SET balance = sender.balance WHERE id = sender_id_arg;
        UPDATE users SET balance = receiver.balance WHERE id = receiver_id_arg;
        cash_VR.count = cash_VR.count + 1;
        cash_count = cash_VR.count;
        UPDATE cash SET count = cash_VR.count WHERE sender_id = sender.id;
        -- UPDATE users 

        INSERT INTO  checks(sender_name ,receiver_name, sender_price,commission_price ,cashback )VALUES
        (sender.name,receiver.name,price,price * percent,MAX_PRICE::INT*percent_cashback);
        RAISE INFO 'MAX %',MAX_PRICE::DECIMAL*percent_cashback;

        END IF ;
        
        IF cash_VR.count = 5 THEN 
        -- MAX
            SELECT MAX(money) FROM money_sender WHERE cash_sender_id = cash_VR.money_id INTO MAX_PRICE; 
            IF MAX_PRICE < price THEN
                MAX_PRICE = price;
            END IF;
            sender.balance = sender.balance - price  + MAX_PRICE*percent_cashback;
            receiver.balance = receiver.balance + price;
            
            INSERT INTO  checks(sender_name ,receiver_name, sender_price,commission_price ,cashback )VALUES
            (sender.name,receiver.name,price,price * percent,MAX_PRICE::INT*percent_cashback);


            
        -- DELETE MONEY_SENDER ROW 
            DELETE FROM money_sender WHERE id IN (SELECT id FROM money_sender WHERE cash_sender_id=1)RETURNING *;
            

            UPDATE cash SET count = cash_VR.count WHERE sender_id = sender.id;
            cash_VR.count = cash_VR.count + 1;
            cash_count = cash_VR.count;




        END IF; 



        IF sender.balance >= 0  THEN
            COMMIT;
        END IF;

        
    END;

$$;

CREATE TABLE checks(
    id SERIAL PRIMARY KEY,
    sender_name VARCHAR,
    receiver_name VARCHAR,
    sender_price INT,
    commission_price INT,
    cashback INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE history(
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES users(id),
    send_money INT ,
    commission_price INT,
    cashback INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);
