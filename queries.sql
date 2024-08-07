-- name: GetIndivSummary :many
WITH "RankedScores" AS (
    SELECT 
        s.event_id,
        u.username,
        u.country,
        s.timems,
        ROW_NUMBER() OVER (PARTITION BY s.event_id ORDER BY min(s.timems) ASC) AS "rank"
    FROM 
        scores s
LEFT    JOIN 
        users u ON s.user_id = u.user_id
    group by s.event_id, s.user_id
)
SELECT 
    rs.username,
    rs.country,
    SUM(CASE WHEN rs.rank = 1 THEN 1 ELSE 0 END) AS gold,
    SUM(CASE WHEN rs.rank = 2 THEN 1 ELSE 0 END) AS silver,
    SUM(CASE WHEN rs.rank = 3 THEN 1 ELSE 0 END) AS bronze,
    SUM(CASE WHEN rs.rank IN (1, 2, 3) THEN 1 ELSE 0 END) AS total_medals,
    SUM(CASE WHEN rs.rank = 1 THEN 3 WHEN rs.rank = 2 THEN 2 WHEN rs.rank = 3 THEN 1 ELSE 0 END) AS total_points
FROM 
    RankedScores rs
GROUP BY 
    rs.username
ORDER BY total_points DESC
LIMIT 10;

-- name: GetCountrySummary :many
WITH "RankedScores" AS (
    SELECT 
        s.event_id,
        u.country,
        s.timems,
        ROW_NUMBER() OVER (PARTITION BY s.event_id ORDER BY min(s.timems) ASC) AS "rank"
    FROM 
        scores s
    JOIN 
        users u ON s.user_id = u.user_id
    GROUP BY 
        s.user_id, s.event_id
)
SELECT 
    rs.country,
    SUM(CASE WHEN rs.rank = 1 THEN 1 ELSE 0 END) AS gold,
    SUM(CASE WHEN rs.rank = 2 THEN 1 ELSE 0 END) AS silver,
    SUM(CASE WHEN rs.rank = 3 THEN 1 ELSE 0 END) AS bronze,
    SUM(CASE WHEN rs.rank IN (1, 2, 3) THEN 1 ELSE 0 END) AS total_medals,
    SUM(CASE WHEN rs.rank = 1 THEN 3 WHEN rs.rank = 2 THEN 2 WHEN rs.rank = 3 THEN 1 ELSE 0 END) AS total_points
FROM 
    RankedScores rs
GROUP BY 
    rs.country
ORDER BY 
total_points DESC
LIMIT 10;

-- name: GetIndivDetails :many
WITH "RankedScores" AS (
    SELECT 
        s.event_id,
        u.username,
        u.country,
        s.timems,
        s.date_entered,
        ROW_NUMBER() OVER (PARTITION BY s.event_id ORDER BY min(s.timems) ASC) AS "rank"
    FROM 
        scores s
LEFT    JOIN 
        users u ON s.user_id = u.user_id
    GROUP BY u.user_id, s.event_id
)
    SELECT country, timems, date_entered, rank, event_type_name, event_level FROM RankedScores
LEFT JOIN events e on RankedScores.event_id = e.event_id
LEFT JOIN events_types et on e.event_type = et.event_type_id
WHERE username = ?
ORDER BY rank asc;


-- name: GetCountryDetails :many
WITH "RankedScores" AS (
    SELECT 
        s.event_id,
        u.username,
        u.country,
        s.timems,
        s.date_entered,
        ROW_NUMBER() OVER (PARTITION BY s.event_id ORDER BY min(s.timems) ASC) AS "rank"
    FROM 
        scores s
LEFT    JOIN 
        users u ON s.user_id = u.user_id
    GROUP BY u.user_id, s.event_id
)
    SELECT username, country, timems, date_entered, rank, event_type_name, event_level FROM RankedScores
LEFT JOIN events e on RankedScores.event_id = e.event_id
LEFT JOIN events_types et on e.event_type = et.event_type_id
WHERE country = ?
ORDER BY rank asc;

-- name: SubmitScore :execresult
INSERT INTO scores (user_id, event_id, timems)
SELECT u.user_id, e.event_id, ?
FROM users u
LEFT JOIN events e ON e.event_level = ? AND e.event_type = ?
WHERE u.username = ?;
