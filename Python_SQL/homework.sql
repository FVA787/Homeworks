CREATE TABLE materials
(
    id            varchar PRIMARY KEY,
    material_name VARCHAR(500),
    unit          VARCHAR(10),

    price_2023_02 DECIMAL(10, 2),
    price_2023_03 DECIMAL(10, 2),
    price_2023_04 DECIMAL(10, 2),
    price_2023_05 DECIMAL(10, 2),
    price_2023_06 DECIMAL(10, 2),
    price_2023_08 DECIMAL(10, 2),
    price_2023_09 DECIMAL(10, 2),
    price_2023_10 DECIMAL(10, 2),
    price_2023_11 DECIMAL(10, 2),
    price_2023_12 DECIMAL(10, 2)
);
select material_name
from materials
where price_2023_03 > 30000;
select avg(price_2023_02),
       avg(price_2023_03),
       avg(price_2023_04),
       avg(price_2023_05),
       avg(price_2023_06),
       avg(price_2023_08),
       avg(price_2023_09),
       avg(price_2023_10),
       avg(price_2023_11),
       avg(price_2023_12)
from materials;
select material_name
from materials
where price_2023_04 > price_2023_03;
select material_name
from materials
where price_2023_12 = (SELECT MAX(price_2023_12)
                       FROM materials);
--.	Увеличьте все цены в столбце price_2023_06 на 5%.
update materials
set price_2023_06 = price_2023_06 * 1.05
where price_2023_06 is not null;
-- 4.	Создайте представление materials_summary, которое будет содержать:
-- a.	Наименование материала.
-- b.	Среднюю цену за год.
-- c.	Минимальную цену за год.
-- d.	Максимальную цену за год.
-- e.	Общее количество записей по каждому материалу.--
CREATE VIEW materials_summary AS
SELECT material_name,
       -- средняя цена за год
       (
           price_2023_02 +
           price_2023_03 +
           price_2023_04 +
           price_2023_05 +
           price_2023_06 +
           price_2023_08 +
           price_2023_09 +
           price_2023_10 +
           price_2023_11 +
           price_2023_12
           )
           /
       NULLIF(
               (price_2023_02 IS NOT NULL)::int +
               (price_2023_03 IS NOT NULL)::int +
               (price_2023_04 IS NOT NULL)::int +
               (price_2023_05 IS NOT NULL)::int +
               (price_2023_06 IS NOT NULL)::int +
               (price_2023_08 IS NOT NULL)::int +
               (price_2023_09 IS NOT NULL)::int +
               (price_2023_10 IS NOT NULL)::int +
               (price_2023_11 IS NOT NULL)::int +
               (price_2023_12 IS NOT NULL)::int,
               0
       ) AS avg_price_year,
       --минимальная цена в году:
       LEAST(
               price_2023_02,
               price_2023_03,
               price_2023_04,
               price_2023_05,
               price_2023_06,
               price_2023_08,
               price_2023_09,
               price_2023_10,
               price_2023_11,
               price_2023_12
       ) AS min_price_year,
       --максимальная цена в году
       GREATEST(
               price_2023_02,
               price_2023_03,
               price_2023_04,
               price_2023_05,
               price_2023_06,
               price_2023_08,
               price_2023_09,
               price_2023_10,
               price_2023_11,
               price_2023_12
       ) AS max_price_year,
       -- количество записей (не NULL цен)
       (price_2023_02 IS NOT NULL)::int +
       (price_2023_03 IS NOT NULL)::int +
       (price_2023_04 IS NOT NULL)::int +
       (price_2023_05 IS NOT NULL)::int +
       (price_2023_06 IS NOT NULL)::int +
       (price_2023_08 IS NOT NULL)::int +
       (price_2023_09 IS NOT NULL)::int +
       (price_2023_10 IS NOT NULL)::int +
       (price_2023_11 IS NOT NULL)::int +
       (price_2023_12 IS NOT NULL)::int
         AS records_count
from materials;

drop view materials_summary;