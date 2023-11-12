-- task-2
CREATE OR REPLACE FUNCTION task_2_TG() RETURNS TRIGGER LANGUAGE PLPGSQL
AS 
$$
    DECLARE 
        t_sale INT;
        p_sale INT;
    BEGIN 
     
        SELECT s.total_amount FROM sale AS s  
        WHERE id = NEW.sale_id
        INTO t_sale;
        RAISE INFO 'T % ',t_sale;


        SELECT s.products_count FROM sale AS s
        WHERE id = NEW.sale_id
        INTO p_sale;
        RAISE INFO 'PC % ',p_sale;

        UPDATE sale 
        SET 
        total_amount = t_sale + (NEW.count * NEW.price),
        products_count = p_sale + NEW.count
        WHERE id = NEW.sale_id;
        RAISE INFO '%',NEW;
        
        RETURN NEW;
    END;
$$;
    CREATE TRIGGER saleTR
    AFTER INSERT ON sale_products
    FOR EACH ROW EXECUTE PROCEDURE task_2_TG();

-- CREATE TRIGGER  sale_sum_tg
-- AFTER INSERT ON sale_products
-- FOR EACH ROW EXECUTE PROCEDURE
-- sale_sum_trigger();

INSERT INTO sale_products(count,price,sale_id) VALUES
(5,120,'35de6e4a-1dbe-44fa-bb38-32929b72c701')
;
 INSERT INTO client (first_name,last_name) VALUES
 ('Jon','Jack')
 ; 

 INSERT INTO sale (date,client_id) VALUES
 ('2023-06-02','888a20e6-67e3-4e7d-bf55-405187aacb0a')
 ;


 CREATE OR REPLACE FUNCTION sale_payments_TG() RETURNS TRIGGER LANGUAGE PLPGSQL
 AS
 $$
    DECLARE 
        sale_q RECORD;
        sale_pr RECORD;
    BEGIN
        SELECT * FROM sale
        INTO sale_q
        WHERE id = NEW.sale_id ;
        RAISE INFO '%' , sale;

        SELECT * FROM sale_products
        INTO sale_pr
        WHERE id = sale_q.id;
        RAISE INFO '%' , sale_pr;
    RETURN NEW;
    END;
 
 $$;

 CREATE TRIGGER sale_pay
 ALTER INSERT ON sale_payments