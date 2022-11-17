package api

import (
	"database/sql"
)

//ResourceID is an alias for the sql package's NullInt64.
//Every model in  the API uses a ResourceID as its primary key.
//ResourceID is used to make the service extensible. If we want to use DB.
type ResourceID sql.NullInt64
