# StandardSQL

This is a query builder that is present in de (Capila) db package.

It will replace certain tags in a query with actual fieldnames, tablenames an so on.

## What functionalities are there?

Most of the following functionalities possibilities are driven by query options.

CountRecordsQuery
- gets the sql query for a count of more records
MoreRecordsQuery
- gets the sql query for more records
OneRecordQuery
- gets the sql query for one record
FullRecordQuery
- gets the sql query for one complete record
InsertRecordQuery
- gets the sql query for the insert of one record
UpdateRecordQuery
- gets the sql query for the update of one record
FileQuery
- gets the sql query from a file
HardDeleteRecordQuery
- gets the sql query for the deletion of one record

## QueryOptions

The following options are present and will be explained more in detail when needed.

type QueryOptions struct {
	Model        interface{}   // the main-model that the query works with
	UpdateFields []string      // the fields that are used in an update sql statement
	TableName    string        // the database tablename that corresponds with the main-model
	Joins        []JoinOptions // see below
	HardDelete   bool          // if the model has a harddelete (if not, a deleted_at should be present)
	ReplaceID    string        // if the id field is named differently
	PageSize     uint          // how many records at once for pagination
	Page         uint          // the page to start for pagination
	CountField   string        // the field that is used in the count
}

// JoinOptions are variable options for the join of the querybuilder
type JoinOptions struct {
	Model     interface{} // the (sub)model that the join works with
	TableName string      // the database tablename that corresponds with the sub-model
	SQLPart   string      // part of the sql where the join is made
}

## Tag replacements

The tags that get replaced, have a special format so Postgres thinks it is comment.

/*<countfield>*/
- this is one field from the database (Advice: use the id-field)
/*<dbfields>*/
- all field of the main-model (and sub-model) that have a valid db-tag
  - becomes: tablename.field1,tablename.field2,subtablename.field3
/*<tableName>*/
- the name of the table of the main-model
/*<joinTableName>*/
- tha name of the table of the sub-model
/*<joinText>*/
- this you have to write yourself and can be something like:
  - INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey
  - LEFT OUTER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey
/*<modelfields>*/
- fields that are used in an insert statement
- only fields of the mainmodel are used
  - becomes: :field1,:field2
/*<updatefields>*/
- fields that are used in an update statement
- only fields of the mainmodel are used
  - becomes: field1=:field1,field2:field2
/*<pagination>*/
- LIMIT and OFFSET will be set on the correct QueryOptions

### CountRecordsQuery

The standard query for this is the following:
`SELECT count(/*<countfield>*/) as count FROM /*<tableName>*/ /*<joinText>*/ WHERE deleted_at IS NULL`

With the following options:
	sql := db.CountRecordsQuery(&db.QueryOptions{
		TableName:  "myTable",
		CountField: "id",
	})
the same query will become something like:
`SELECT count(id) as count FROM myTable WHERE deleted_at IS NULL`

With the following options:
	sql := db.CountRecordsQuery(&db.QueryOptions{
		TableName:  "myTable",
		CountField: "id",
		HardDelete: true,
	})
the same query will become something like:
`SELECT count(id) as count FROM myTable WHERE TRUE`

This query can also have join options. See the MoreRecordQuery for those options.

### MoreRecordsQuery

You can use this as a oneliner:
	sql := db.MoreRecordsQuery(&db.QueryOptions{Model: datamodels.MyCompactModel{}, TableName: "myTable"})

But for readability, I advice the following:
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model: datamodels.MyCompactModel{},
		TableName: "myTable",
	})

> Hint: Notice the last option has a comma too
> Hint: on a MoreRecordsQuery it is logical to have a compact model, because it will probably be used in some list format

You tell the MoreRecordsQuery function, you need a sql statement with fields from datamodels.MyCompactModel
and the table to use is myTable.

The default query for this is:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ /*<joinText>*/ WHERE deleted_at IS NULL /*<pagination>*/`

Since there aren't any joins and no pagination given in the QueryOptions, there query will be:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ WHERE deleted_at IS NULL`

It will be replaced into something like this:
`SELECT myTable.field1,myTable.field2 FROM myTable WHERE deleted_at IS NULL`

With the following options:
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model: datamodels.MyCompactModel{},
		TableName: "myTable",
		HardDelete: true,
	})
the same query will become something like:
`SELECT myTable.field1,myTable.field2 FROM myTable WHERE TRUE`

The end is like this so you can easy add stuff:
sql += ` and day=12`

With the following options:
	myJoin := db.JoinOptions{
		Model:     datamodels.MyCompactSubModel{},
		TableName: "mySubTable",
		SQLPart:   "INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey",
	}
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model:     datamodels.MyCompactModel{},
		TableName: "myTable",
		Joins:     []db.JoinOptions{myJoin},
	})
we now tell the query builder to add a join-part to the query.
We had this:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ /*<joinText>*/ WHERE deleted_at IS NULL`
So it will become
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey WHERE deleted_at IS NULL`
The other tag-replacments will turn it into something like this:
`SELECT myTable.field1,myTable.field2,mySubTable.field3 FROM myTable INNER JOIN mySubTable ON mySubTable.id = myTable.subkey WHERE deleted_at IS NULL`

You can even do multiple joins which act the same as the above:
	myJoin := db.JoinOptions{
		Model:     datamodels.MyCompactSubModel{},
		TableName: "mySubTable",
		SQLPart:   "INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey",
	}
	myOtherJoin := db.JoinOptions{
		Model:     datamodels.MyCompactOtherSubModel{},
		TableName: "myOtherSubTable",
		SQLPart:   "LEFT OUTER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey",
	}
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model:     injectMainModel{},
		TableName: "myTable",
		Joins:     []db.JoinOptions{myJoin,myOtherJoin},
	})

You can also use the pagination with the following options:
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model: datamodels.MyCompactModel{},
		TableName: "myTable",
		Page:      0,
		PageSize:  10,
	})
`SELECT myTable.field1,myTable.field2 FROM myTable WHERE deleted_at IS NULL LIMIT 10 OFFSET 0`

If you set page to 2, it will try to find record 21 to 30:
`SELECT myTable.field1,myTable.field2 FROM myTable WHERE deleted_at IS NULL LIMIT 10 OFFSET 20`

### OneRecordQuery

This is a typical usage:
	sql := db.OneRecordQuery(&db.QueryOptions{
		Model: datamodels.MyFullModel{},
		TableName: "myTable",
	})

> Hint: on a OneRecordQuery it is logical to have a richer model, because it is one detail record

You tell the OneRecordQuery function, you need a sql statement with fields from datamodels.MyFullModel
and the table to use is myTable.

The default query for this is:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ /*<joinText>*/ WHERE deleted_at IS NULL AND id = :id`

Since there aren't any joins given in the QueryOptions, there query will be:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ WHERE deleted_at IS NULL AND id = :id`

It will be replaced into something like this:
`SELECT myTable.field1,myTable.field2,myTable.field3,myTable.field4 FROM myTable WHERE deleted_at IS NULL AND id = :id`

With the following options:
	sql := db.OneRecordQuery(&db.QueryOptions{
		Model: datamodels.MyFullModel{},
		TableName: "myTable",
		HardDelete: true,
	})

The standard query will become something like:
`SELECT myTable.field1,myTable.field2,myTable.field3,myTable.field4 FROM myTable WHERE TRUE AND id = :id`

With the following options:
	sql := db.OneRecordQuery(&db.QueryOptions{
		Model: datamodels.MyFullModel{},
		TableName: "myTable",
		ReplaceID: "prefix_key",
	})
The standard query will become something like:
`SELECT myTable.field1,myTable.field2,myTable.field3,myTable.field4 FROM myTable WHERE deleted_at IS NULL AND prefix_key = :prefix_key`

With the following options
	myJoin := db.JoinOptions{
		Model:     datamodels.MyCompactSubModel{},
		TableName: "mySubTable",
		SQLPart:   "INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey",
	}
	sql := db.MoreRecordsQuery(&db.QueryOptions{
		Model:     datamodels.MyCompactModel{},
		TableName: "myTable",
		Joins:     []db.JoinOptions{myJoin},
	})
we now tell the query builder to add a join-part to the query.
We had this:
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ /*<joinText>*/ WHERE deleted_at IS NULL AND id = :id`
So it will become
`SELECT /*<dbfields>*/ FROM /*<tableName>*/ INNER JOIN /*<joinTableName>*/ ON /*<joinTableName>*/.id = /*<tableName>*/.subkey WHERE deleted_at IS NULL AND id = :id`
The other tag-replacments will turn it into something like this:
`SELECT myTable.field1,myTable.field2,mySubTable.field3 FROM myTable INNER JOIN mySubTable ON mySubTable.id = myTable.subkey WHERE deleted_at IS NULL AND id = :id`

And of course, this query can have multiple joins as well.

## FullRecordQuery

The standard query for this is just to get access to a specific record (which skips the soft delete part).
`SELECT * FROM /*<tableName>*/ WHERE id = :id`

Since tableName is the only variable part, the following call is enough:
	sql := db.FullRecordQuery(&db.QueryOptions{
		TableName: "myTable",
	})
which will result in something like this:
`SELECT * FROM myTable WHERE id = :id`

With the following options:
	sql := db.FullRecordQuery(&db.QueryOptions{
		TableName: "myTable",
		ReplaceID: "prefix_key",
	})
The standard query will become something like:
`SELECT * FROM myTable WHERE prefix_key = :prefix_key`

## InsertRecordQuery

The standard query for this is just to add a new record to the database.
`INSERT INTO /*<tableName>*/ (/*<dbfields>*/) VALUES (/*<modelfields>*/) RETURNING id;`

With the following options:
	sql := db.InsertRecordQuery(&db.QueryOptions{
		Model:     datamodels.MyModel{},
		TableName: "myTable",
	})
the query will result in something like this:
`INSERT INTO myTable (myTable.field1,myTable.field2) VALUES (:field1,:field2) RETURNING id;`

With the following options:
	sql := db.InsertRecordQuery(&db.QueryOptions{
		Model:     datamodels.MyModel{},
		TableName: "myTable",
		ReplaceID: "prefix_key",
	})
The standard query will become something like:
`INSERT INTO myTable (myTable.field1,myTable.field2) VALUES (:field1,:field2) RETURNING prefix_key;`

## UpdateRecordQuery

The standard query for this is just to change (some of the) fields of an existing record.
`INSERT INTO /*<tableName>*/ (/*<dbfields>*/) VALUES (/*<modelfields>*/) RETURNING id;`

With the following options:
	dbFields := []string{"myfield"}
	sql := db.UpdateRecordQuery(&db.QueryOptions{
		UpdateFields: dbFields,
		TableName:    "myTable",
	})
the query will result in something like this:
`UPDATE myTable SET myfield=:myfield WHERE id = :id`

Of course, you can infuence the id-field.

## HardDeleteRecordQuery

The standard query for this is just to get access to a specific record (which skips the soft delete part).
`DELETE FROM /*<tableName>*/ WHERE id = :id`

Since tableName is the only variable part, the following call is enough:
	sql := db.FullRecordQuery(&db.QueryOptions{
		TableName: "myTable",
	})
which will result in something like this:
`DELETE FROM myTable WHERE id = :id`

With the following options:
	sql := db.FullRecordQuery(&db.QueryOptions{
		TableName: "myTable",
		ReplaceID: "prefix_key",
	})
The standard query will become something like:
`DELETE FROM myTable WHERE prefix_key = :prefix_key`

## FileQuery

A FileQuery just gets the sql statement from an existing file.
This makes it quite useful to make a query in your database-client and put the end result in a file.

You use it like this (minimal format)
	sql, err := db.FileQuery("/mysql/myquery")
It will use the root directory of the project (or in the directory where you start your executable)
and then try to read the file, which is myquery.sql
If it it not found, you will get an error.

You can also use the QueryBuilder a bit more, by using options:
	sql, err = db.FileQuery("/mysql/myquery", &db.QueryOptions{
		TableName: "myTable",
		ReplaceID: "key",
	})
So if you have stored a query in that file with tag-replacments, it will result in a well formed query.
