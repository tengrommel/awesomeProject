package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.planetmeican.com/planet/dynamodbdao"
	"reflect"
)

type JobA struct {
	ThisId  string
	OkId    string
	Name    string
	Address []byte
}

type Address struct {
	A string
}

func (j *JobA) GetKey() string {
	return j.OkId
}

var TableJobDescription = dynamodb.TableDescription{
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("ThisId"),
			AttributeType: aws.String("S"),
		},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("ThisId"),
			KeyType:       aws.String("HASH"),
		},
	},
	TableName: aws.String("jobB"),
}

var createTable = dynamodb.CreateTableInput{
	AttributeDefinitions: TableJobDescription.AttributeDefinitions,
	KeySchema:            TableJobDescription.KeySchema,
	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	},
	TableName: TableJobDescription.TableName,
}

func main() {
	theCredential := credentials.NewStaticCredentials("222", "222", "")
	config := &aws.Config{
		Region:      aws.String("us-west-2"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: theCredential,
	}
	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)
	genericDao, err := dynamodbdao.New(svc, "jobB")
	if err != nil {
		panic(err)
	}
	address := Address{A: "d"}
	jsonAddress, _ := json.Marshal(address)
	err = genericDao.Put(&JobA{ThisId: "69d6190d-ec16-4dc5-b296-d021e4bbe805", OkId: "fuck", Address: jsonAddress})
	if err != nil {
		panic(err)
	}
	job := &JobA{}
	err = genericDao.Get("69d6190d-ec16-4dc5-b296-d021e4bbe805", job)
	fmt.Println(reflect.TypeOf(job))
	fmt.Println(reflect.TypeOf(job.Address))
	aa := &Address{}
	err = json.Unmarshal(job.Address, aa)
	fmt.Println(err)
	fmt.Println(aa.A)
}
