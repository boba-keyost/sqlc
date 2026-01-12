/* name: ListItemsAndSortSingle :many */
SELECT * FROM sort_items
ORDER BY sqlc.sort ( sort_by);

/* name: ListItemsAndSort :many */
SELECT * FROM sort_items
ORDER BY sqlc.sort ( sort_by, order_dir);

/* name: ListItemsAndSortDefField :many */
SELECT * FROM sort_items
ORDER BY sqlc.sort ( sort_by, order_dir, name);

/* name: ListItemsAndSortDefAll :many */
SELECT * FROM sort_items
ORDER BY sqlc.sort ( sort_by, order_dir, name, 'asc');