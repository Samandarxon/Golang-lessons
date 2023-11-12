SELECT e.name AS EMPLOYEE FROM EMPLOYEE AS e, EMPLOYEE  AS m
WHERE e.managerId = m.id AND e.salary>m.salary
