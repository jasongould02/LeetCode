SELECT Email FROM (SELECT email, COUNT(email) AS num FROM Person
GROUP BY email) AS Duplicates WHERE num > 1;
