package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var TableJobDescription = dynamodb.TableDescription{
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("ReportUUID"),
			AttributeType: aws.String("S"),
		},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("ReportUUID"),
			KeyType:       aws.String("HASH"),
		},
	},
	TableName: aws.String("Task0"),
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
		Region:      aws.String("cn-north-1"),
		Endpoint:    aws.String("http://localhost:8000"),
		Credentials: theCredential,
	}
	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)
	result, err := svc.CreateTable(&createTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
	//genericDao, err := dynamodbdao.New(svc, "jobB")
	//if err != nil {
	//	panic(err)
	//}
	//address := Address{A: "d"}
	//err = genericDao.Put(&JobA{ThisId: "69d6190d-ec16-4dc5-b296-d021e4bbe805", OkId: "fuck", Address: address})
	//if err != nil {
	//	panic(err)
	//}
	//job := &JobA{}
	//err = genericDao.Get("69d6190d-ec16-4dc5-b296-d021e4bbe805", job)
	//fmt.Println(reflect.TypeOf(job))
	//fmt.Println(reflect.TypeOf(job.Address))
	//a := Address{}
	//err = mapstructure.Decode(&job.Address, &a)
	//fmt.Println(a.A)
}
