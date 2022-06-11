package dynamo

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
)

// Store provides functionalities to access dynamodb.
type Store struct {
	db         dynamodbiface.DynamoDBAPI
	tableName  string
	now        func() time.Time
	generateID func() string
}

func generateID() string {
	return uuid.New().String()
}

func NewStore(tableName string) Store {
	return NewCustomStore(
		dynamodb.New(session.Must(session.NewSession())), //GetLocalAWSSession()),
		tableName,
		time.Now,
		generateID,
	)
}

func NewCustomStore(db dynamodbiface.DynamoDBAPI, tableName string, now func() time.Time, generateIDFunc func() string) Store {
	s := Store{
		db:         db,
		tableName:  tableName,
		now:        now,
		generateID: generateIDFunc,
	}
	if s.now == nil {
		s.now = time.Now
	}
	if s.generateID == nil {
		s.generateID = generateID
	}
	return s
}
