SELECT salary FROM employee 
WHERE salary= (SELECT DISTINCT(salary) 
FROM employee ORDER BY salary LIMIT 1,1);