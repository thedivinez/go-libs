package storage

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database interface defines all storage operations
type Database interface {
	// Basic CRUD operations
	DeleteOne(ctx context.Context, collection string, filter any) error
	DeleteMany(ctx context.Context, collection string, filter any) (int64, error)
	Count(ctx context.Context, collection string, filter any) (int64, error)
	InsertOne(ctx context.Context, collection string, data any) (string, error)
	InsertMany(ctx context.Context, collection string, data []interface{}) error
	UpdateOne(ctx context.Context, collection string, filter any, update any, opts ...*options.UpdateOptions) error
	UpdateMany(ctx context.Context, collection string, filter any, update any, opts ...*options.UpdateOptions) (int64, error)
	ReplaceOne(ctx context.Context, collection string, filter any, replacement any, opts ...*options.ReplaceOptions) error

	// Legacy methods for compatibility
	Find(ctx context.Context, collection string, filter any, results interface{}, opts ...*options.FindOptions) error
	Aggregate(ctx context.Context, collection string, pipeline any, results any, opts ...*options.AggregateOptions) error
	FindOne(ctx context.Context, collection string, filter any, results interface{}, opts ...*options.FindOneOptions) error

	// Schema management
	CreateCollection(ctx context.Context, name string, validator bson.M, opts ...*options.CreateCollectionOptions) error
	CreateIndexes(ctx context.Context, collection string, indexes []mongo.IndexModel) error
	DropCollection(ctx context.Context, name string) error
	CollectionExists(ctx context.Context, name string) (bool, error)

	// Utility
	GenerateID(ctx context.Context, collection string, letters int, digits int) (*ID, error)
	Transaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error

	// Health and monitoring
	Ping(ctx context.Context) error
	GetDatabase() *mongo.Database
	Disconnect(ctx context.Context) error
}

// MongoStorage implements the Database interface
type MongoStorage struct {
	client *mongo.Client
	db     *mongo.Database
}

// PageResult holds paginated results
type PageResult[T any] struct {
	Data       []T   `json:"data"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

// ID represents a generated identifier
type ID struct {
	Value string `json:"value"`
}

// NewMongoStorage creates a new MongoDB storage instance
func NewMongoStorage(ctx context.Context, uri, dbname string) (*MongoStorage, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB")
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping MongoDB")
	}

	return &MongoStorage{client: client, db: client.Database(dbname)}, nil
}

// Disconnect closes the MongoDB connection
func (s *MongoStorage) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

// GetDatabase returns the underlying mongo.Database
func (s *MongoStorage) GetDatabase() *mongo.Database {
	return s.db
}

// Ping checks if the database is accessible
func (s *MongoStorage) Ping(ctx context.Context) error {
	return s.client.Ping(ctx, nil)
}

/*******************************************
*           SCHEMA MANAGEMENT              *
*******************************************/

// CreateCollection creates a new collection with optional validator
func (s *MongoStorage) CreateCollection(ctx context.Context, name string, validator bson.M, opts ...*options.CreateCollectionOptions) error {
	collOpts := options.CreateCollection()
	if validator != nil {
		collOpts.SetValidator(validator)
	}

	// Merge any additional options
	for _, opt := range opts {
		if opt.Capped != nil && *opt.Capped {
			collOpts = collOpts.SetCapped(true)
		}
		if opt.SizeInBytes != nil {
			collOpts.SetSizeInBytes(*opt.SizeInBytes)
		}
		if opt.MaxDocuments != nil {
			collOpts.SetMaxDocuments(*opt.MaxDocuments)
		}
	}

	err := s.db.CreateCollection(ctx, name, collOpts)
	if err != nil {
		return errors.Wrapf(err, "failed to create collection %s", name)
	}
	return nil
}

// CreateIndexes creates multiple indexes on a collection
func (s *MongoStorage) CreateIndexes(ctx context.Context, collection string, indexes []mongo.IndexModel) error {
	if len(indexes) == 0 {
		return nil
	}

	_, err := s.db.Collection(collection).Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return errors.Wrapf(err, "failed to create indexes on collection %s", collection)
	}
	return nil
}

// DropCollection drops a collection
func (s *MongoStorage) DropCollection(ctx context.Context, name string) error {
	err := s.db.Collection(name).Drop(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to drop collection %s", name)
	}
	return nil
}

// CollectionExists checks if a collection exists
func (s *MongoStorage) CollectionExists(ctx context.Context, name string) (bool, error) {
	names, err := s.db.ListCollectionNames(ctx, bson.M{"name": name})
	if err != nil {
		return false, errors.Wrap(err, "failed to list collections")
	}
	return len(names) > 0, nil
}

/*******************************************
*        LEGACY COMPATIBLE METHODS         *
*******************************************/

// Find finds multiple documents (legacy compatible)
func (s *MongoStorage) Find(ctx context.Context, collection string, filter any, results interface{}, opts ...*options.FindOptions) error {
	cursor, err := s.db.Collection(collection).Find(ctx, filter, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to find documents")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, results); err != nil {
		return errors.Wrap(err, "failed to decode documents")
	}
	return nil
}

// FindOne finds a single document (legacy compatible)
func (s *MongoStorage) FindOne(ctx context.Context, collection string, filter any, results interface{}, opts ...*options.FindOneOptions) error {
	err := s.db.Collection(collection).FindOne(ctx, filter, opts...).Decode(results)
	if err != nil {
		return errors.Wrap(err, "failed to find document")
	}
	return nil
}

/*******************************************
*         STANDARD CRUD OPERATIONS         *
*******************************************/

// DeleteOne deletes a single document
func (s *MongoStorage) DeleteOne(ctx context.Context, collection string, filter any) error {
	result, err := s.db.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "failed to delete document")
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found to delete")
	}
	return nil
}

// DeleteMany deletes multiple documents and returns count
func (s *MongoStorage) DeleteMany(ctx context.Context, collection string, filter any) (int64, error) {
	result, err := s.db.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "failed to delete documents")
	}
	return result.DeletedCount, nil
}

// Count counts documents matching filter
func (s *MongoStorage) Count(ctx context.Context, collection string, filter any) (int64, error) {
	count, err := s.db.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count documents")
	}
	return count, nil
}

// InsertOne inserts a single document and returns its ID
func (s *MongoStorage) InsertOne(ctx context.Context, collection string, data any) (string, error) {
	result, err := s.db.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return "", errors.Wrap(err, "failed to insert document")
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return fmt.Sprintf("%v", result.InsertedID), nil
}

// InsertMany inserts multiple documents
func (s *MongoStorage) InsertMany(ctx context.Context, collection string, data []interface{}) error {
	if len(data) == 0 {
		return nil
	}

	_, err := s.db.Collection(collection).InsertMany(ctx, data)
	if err != nil {
		return errors.Wrap(err, "failed to insert documents")
	}
	return nil
}

// UpdateOne updates a single document
func (s *MongoStorage) UpdateOne(ctx context.Context, collection string, filter any, update any, opts ...*options.UpdateOptions) error {
	result, err := s.db.Collection(collection).UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to update document")
	}
	if result.MatchedCount == 0 {
		return errors.New("no document found to update")
	}
	return nil
}

// UpdateMany updates multiple documents and returns count
func (s *MongoStorage) UpdateMany(ctx context.Context, collection string, filter any, update any, opts ...*options.UpdateOptions) (int64, error) {
	result, err := s.db.Collection(collection).UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to update documents")
	}
	return result.ModifiedCount, nil
}

// ReplaceOne replaces a single document
func (s *MongoStorage) ReplaceOne(ctx context.Context, collection string, filter any, replacement any, opts ...*options.ReplaceOptions) error {
	result, err := s.db.Collection(collection).ReplaceOne(ctx, filter, replacement, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to replace document")
	}
	if result.MatchedCount == 0 {
		return errors.New("no document found to replace")
	}
	return nil
}

/*******************************************
*          TRANSACTION SUPPORT             *
*******************************************/

// Transaction executes a function within a transaction
func (s *MongoStorage) Transaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	session, err := s.client.StartSession()
	if err != nil {
		return errors.Wrap(err, "failed to start session")
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})
	return err
}

func (s *MongoStorage) Aggregate(ctx context.Context, collection string, pipeline, results any, opts ...*options.AggregateOptions) error {
	cursor, err := s.db.Collection(collection).Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to execute aggregation")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, results); err != nil {
		return errors.Wrap(err, "failed to decode aggregation results")
	}
	return nil
}

/*******************************************
*           UTILITY FUNCTIONS              *
*******************************************/

// GenerateID generates a unique ID with specified format
func (s *MongoStorage) GenerateID(ctx context.Context, collection string, letters, digits int) (*ID, error) {
	generate := func() (*ID, error) {
		letterBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterPart := make([]byte, letters)

		for i := range letterPart {
			idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
			if err != nil {
				return nil, err
			}
			letterPart[i] = letterBytes[idx.Int64()]
		}

		digitPart, err := generateRandomNumber(digits)
		if err != nil {
			return nil, err
		}

		return &ID{
			Value: fmt.Sprintf("%s%0*d", string(letterPart), digits, digitPart),
		}, nil
	}

	// Try up to 10 times to generate a unique ID
	for i := 0; i < 10; i++ {
		id, err := generate()
		if err != nil {
			continue
		}

		// Check if ID already exists
		count, err := s.Count(ctx, collection, bson.M{"id": id.Value})
		if err != nil {
			continue
		}
		if count == 0 {
			return id, nil
		}
	}

	return nil, errors.New("failed to generate unique ID after 10 attempts")
}

// generateRandomNumber generates a random number with specified digits
func generateRandomNumber(numberOfDigits int) (int, error) {
	if numberOfDigits < 1 {
		return 0, errors.New("number of digits must be at least 1")
	}

	maxLimit := int64(math.Pow10(numberOfDigits) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}

	randomNumberInt := int(randomNumber.Int64())
	if randomNumberInt < lowLimit {
		randomNumberInt += lowLimit
	}
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}

	return randomNumberInt, nil
}

// Transcode converts between data structures using JSON marshaling
func (s *MongoStorage) Transcode(in, out any) error {
	data, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "failed to marshal input")
	}
	if err := json.Unmarshal(data, out); err != nil {
		return errors.Wrap(err, "failed to unmarshal output")
	}
	return nil
}

/*******************************************
*    STANDALONE GENERIC FUNCTIONS          *
*******************************************/

// FindOneTyped finds a single document and returns it as type T
func FindOneTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, opts ...*options.FindOneOptions) (*T, error) {
	var result T
	err := s.db.Collection(collection).FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to find document")
	}
	return &result, nil
}

// FindTyped finds multiple documents and returns them as []T
func FindTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, opts ...*options.FindOptions) ([]T, error) {
	cursor, err := s.db.Collection(collection).Find(ctx, filter, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find documents")
	}
	defer cursor.Close(ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode documents")
	}
	return results, nil
}

// FindOneAndUpdateTyped finds and updates a document, returning the updated document
func FindOneAndUpdateTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, update any, opts ...*options.FindOneAndUpdateOptions) (*T, error) {
	var result T
	err := s.db.Collection(collection).FindOneAndUpdate(ctx, filter, update, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to find and update document")
	}
	return &result, nil
}

// FindOneAndReplaceTyped finds and replaces a document, returning the replaced document
func FindOneAndReplaceTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, replacement any, opts ...*options.FindOneAndReplaceOptions) (*T, error) {
	var result T
	err := s.db.Collection(collection).FindOneAndReplace(ctx, filter, replacement, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to find and replace document")
	}
	return &result, nil
}

// FindOneAndDeleteTyped finds and deletes a document, returning the deleted document
func FindOneAndDeleteTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, opts ...*options.FindOneAndDeleteOptions) (*T, error) {
	var result T
	err := s.db.Collection(collection).FindOneAndDelete(ctx, filter, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to find and delete document")
	}
	return &result, nil
}

// GetPageTyped returns paginated results with type safety
func GetPageTyped[T any](ctx context.Context, s *MongoStorage, collection string, filter any, page int64, limit int64, sortField string, sortOrder int) (*PageResult[T], error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if sortOrder != 1 && sortOrder != -1 {
		sortOrder = -1
	}

	skip := (page - 1) * limit

	findOpts := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetSort(bson.M{sortField: sortOrder})

	results, err := FindTyped[T](ctx, s, collection, filter, findOpts)
	if err != nil {
		return nil, err
	}

	total, err := s.Count(ctx, collection, filter)
	if err != nil {
		return nil, err
	}

	totalPages := int64(math.Ceil(float64(total) / float64(limit)))

	return &PageResult[T]{
		Data:       results,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

// AggregateTyped performs aggregation and returns typed results
func AggregateTyped[T any](ctx context.Context, s *MongoStorage, collection string, pipeline any) ([]T, error) {
	cursor, err := s.db.Collection(collection).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute aggregation")
	}
	defer cursor.Close(ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "failed to decode aggregation results")
	}
	return results, nil
}

/*******************************************
*           HELPER FUNCTIONS               *
*******************************************/

// NewObjectID creates a new MongoDB ObjectID
func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// NewObjectIDFromHex creates an ObjectID from hex string
func NewObjectIDFromHex(hex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hex)
}

// IsValidObjectID checks if a string is a valid ObjectID
func IsValidObjectID(hex string) bool {
	_, err := primitive.ObjectIDFromHex(hex)
	return err == nil
}

// WithTimeout creates a context with timeout
func WithTimeout(parent context.Context, duration time.Duration) (context.Context, context.CancelFunc) {
	if duration == 0 {
		duration = 30 * time.Second
	}
	return context.WithTimeout(parent, duration)
}

// MustObjectID panics if ObjectID creation fails (useful for testing)
func MustObjectID(hex string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return oid
}

// ParsePage parses page number from string
func ParsePage(page string) int64 {
	if page == "" {
		return 1
	}
	p, err := strconv.ParseInt(page, 10, 64)
	if err != nil || p < 1 {
		return 1
	}
	return p
}
