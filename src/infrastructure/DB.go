package infrastructure

import (
    "github.com/guregu/dynamo"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "os"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

func NewDB() *dynamo.DB {
    //dynamoDBへ接続
    var db *dynamo.DB

    sess := session.Must(session.NewSession())

    if os.Getenv("GO_ENV") == "production" {
      db = dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
    }else{
      creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
      db = dynamo.New(sess,aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithCredentials(creds),)
    }

    return db
}
