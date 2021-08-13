WITH A AS (
  SELECT
    *
  FROM `hoge-11111.dataset.t`
  WHERE id = 1 AND name = "mike" OR account_type = 3
  ORDER BY id ASC
),
B AS(
  SELECT
    id AS id,
    SUM(cost) AS total_cost,
    COUNT(*) AS total_count
  FROM (
    SELECT
      cost AS cost,
      id AS id,
      SPLIT(hoge, ",") AS hoge_arr
    FROM payment p
    INNER JOIN customer c
    ON p.customer_id = c.id
  )
  GROUP BY id
  UNION ALL
  SELECT
    100 AS total_cost,
    300 AS total_count
)
SELECT
  A.*,
  B.*
FROM A
LEFT JOIN B ON A.id = B.id
