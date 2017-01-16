package godb

// Database is reprsenting the database
type Database interface {
	// Close closes the connection to the database
	Close()

	// Collection is getting collection by it's name
	Collection(name string) Collection

	// Gets Indexer for the provided collection name
	Indexer(cname string) Indexer
}

// Collection is a collection in the database. Single database is having
// multiple collections and this interface is representing one of them.
type Collection interface {
	Find(query interface{}) Query

	FindID(id interface{}) Query

	Insert(doc interface{}) error

	Update(selector interface{}, update interface{}) error

	Upsert(selector interface{}, update interface{}) (*ChangeInfo, error)

	Remove(selector interface{}) error

	RemoveID(id interface{}) error

	RemoveAll(selector interface{}) (*ChangeInfo, error)

	Bulk() Bulk
}

// Indexer is responsible for creation of indexes
type Indexer interface {
	// CreateAll creates all provided indexes
	CreateAll([]Index) error
}

type Query interface {
	All(result interface{}) error

	Apply(change Change, result interface{}) (*ChangeInfo, error)

	One(result interface{}) error

	Iter() Iter

	Count() (int, error)

	Limit(n int) Query

	Sort(fields ...string) Query

	Select(selector interface{}) Query
}

type Iter interface {
	Next(result interface{}) bool
}

type Bulk interface {
	Run() (*BulkResult, error)

	Insert(docs ...interface{})

	Update(pairs ...interface{})
}

// Config is a configuration object used for the communication with
// the database
type Config struct {
	Addrs []string // slice of hosts

	Database string // name of the database

	MaxRetryAttempts int // number of max retry attempts
}

type Change struct {
	Update    interface{}
	Upsert    bool
	Remove    bool
	ReturnNew bool
}

type ChangeInfo struct {
	Updated    int
	Removed    int
	UpsertedID interface{}
}

// BulkResult is a result of the Bulk operation. It indicates
// how many records are affected
type BulkResult struct {
	Matched  int
	Modified int
}

// Index is representing a single index in the datastore
type Index struct {
	// The key which to participates in the index. It's a slice
	// to support and composite indexes
	Key []string

	// Unique is flag which indicates whether index should
	// be unique or not
	Unique bool
}

