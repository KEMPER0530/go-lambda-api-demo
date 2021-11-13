package infrastructure

import (
    "github.com/guregu/dynamo"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "os"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

func NewDB() *dynamo.DB {
    //DBへ接続
    sess := session.Must(session.NewSession())
    creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
    db := dynamo.New(sess,aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithCredentials(creds),)

    return db
}
