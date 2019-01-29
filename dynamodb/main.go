package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.planetmeican.com/titan/log"
)

var TableTaskDescription = dynamodb.TableDescription{
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("reportUUID"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("outerUUID"),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String("state"),
			AttributeType: aws.String("N"),
		},
	},
	GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndexDescription{
		{
			IndexName: aws.String("outerUUID"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("outerUUID"),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
		{
			IndexName: aws.String("state"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("state"),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("reportUUID"),
			KeyType:       aws.String("HASH"),
		},
	},
	TableName: aws.String("Task11"),
}

var createTable = dynamodb.CreateTableInput{
	AttributeDefinitions: TableTaskDescription.AttributeDefinitions,
	GlobalSecondaryIndexes: func(sid []*dynamodb.GlobalSecondaryIndexDescription) []*dynamodb.GlobalSecondaryIndex {
		si := make([]*dynamodb.GlobalSecondaryIndex, len(sid))
		for i := range sid {
			si[i] = &dynamodb.GlobalSecondaryIndex{
				IndexName: sid[i].IndexName,
				KeySchema: sid[i].KeySchema,
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
				Projection: sid[i].Projection,
			}
		}
		return si
	}(TableTaskDescription.GlobalSecondaryIndexes),
	KeySchema: TableTaskDescription.KeySchema,
	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	},
	TableName: TableTaskDescription.TableName,
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
	_, err := svc.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: TableTaskDescription.TableName,
	})
	if err != nil {
		log.Fatal(err)
	}
	result, err := svc.CreateTable(&createTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
