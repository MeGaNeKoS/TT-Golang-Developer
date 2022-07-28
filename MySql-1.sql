SELECT worker.Name 
FROM Employee worker, 
	Employee manager
WHERE worker.ManagerId  = manager.Id 
AND worker.Salary> manager.Salary;