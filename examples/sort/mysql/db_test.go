//go:build examples

package sort

import (
	"context"
	"database/sql"
	"path"
	"reflect"
	"runtime"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

	"github.com/boba-keyost/sqlc/internal/sqltest/local"
)

func testList[ArgType any](
	t *testing.T,
	expectedNames []string,
	queryFunc func(ctx context.Context, args ArgType) ([]SortItem, error),
	ctx context.Context,
	arg ArgType,
) {
	items, err := queryFunc(ctx, arg)
	if err != nil {
		t.Fatal(err)
	}
	actual := make([]string, len(items))
	for i, itm := range items {
		actual[i] = itm.Name
	}
	if diff := cmp.Diff(actual, expectedNames); diff != "" {
		qn := path.Base(runtime.FuncForPC(reflect.ValueOf(queryFunc).Pointer()).Name())
		t.Errorf("%s get items names mismatch:\n%s", qn, diff)
	}
}

func TestSort(t *testing.T) {
	ctx := context.Background()
	uri := local.MySQL(t, []string{"schema.sql"})
	sdb, err := sql.Open("mysql", uri)
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()

	db := New(sdb)

	null := sql.NullString{}
	created := sql.NullString{}
	asc := sql.NullString{}
	desc := sql.NullString{}

	_ = created.Scan("created")
	_ = asc.Scan("desc")
	_ = desc.Scan("desc")

	testList(t, []string{"d", "a", "c", "b"}, db.ListItemsAndSortSingle, ctx, created)

	testList(
		t,
		[]string{"b", "c", "a", "d"},
		db.ListItemsAndSort,
		ctx,
		ListItemsAndSortParams{created, desc},
	)

	testList(
		t,
		[]string{"d", "c", "b", "a"},
		db.ListItemsAndSortDefField,
		ctx,
		ListItemsAndSortDefFieldParams{null, desc},
	)

	testList(
		t,
		[]string{"a", "b", "c", "d"},
		db.ListItemsAndSortDefAll,
		ctx,
		ListItemsAndSortDefAllParams{null, null},
	)
}
