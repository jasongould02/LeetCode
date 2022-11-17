# Write your MySQL query statement below
SELECT ActorDirector.actor_id, ActorDirector.director_id FROM ActorDirector
GROUP BY ActorDirector.actor_id, ActorDirector.director_id
HAVING COUNT(ActorDirector.timestamp) >= 3;
