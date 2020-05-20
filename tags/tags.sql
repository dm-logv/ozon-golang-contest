with goods_with_all_tags as (
    select goods_id
    from tags_goods
    group by goods_id
    having count(tag_id) = (
           select count(id)
           from tags)
)
    select goods.id, goods.name
    from goods goods
    join goods_with_all_tags selected
        on selected.goods_id = goods.id
;
