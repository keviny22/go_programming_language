package main

import (
	"encoding/json"
	"fmt"
)

type Statement struct {
	Sid    string `json:"Sid"`
	Effect string `json:"Effect"`
	Action string `json:"Action"`
	Resource string `json:"Resource"`
	Principal map[string]interface{}
}

//type Principal struct {
//	PrincipalArn PrincipalArn `json:"Aws"`
//}
//type Principal []string

//type PrincipalArn struct {
//	Arn string  `json:""`
//}

type Policy struct {
	Version   string
	Id        string
	Statement []Statement `"json: Statement"`
}

func main() {

	a := s3SourceBucketPolicy("codepipeline-bucket-us-east-1")
	fmt.Println(a)
}
func s3SourceBucketPolicy(bucket string) string {
	targets := [...]string{"108755494865","108755494866","243545454"}

    for i, target := range targets {
		targets[i] = "arn:aws:iam::" + target + ":root"
	}

	principals := map[string]interface{}{
		"AWS": targets,
	}

	listStatement := Statement{
		Sid:    "Allow source account to List Bucket in pipeline account",
		Effect: "Allow",
		Action: "s3:ListBucket",
		Resource: "arn:aws:s3:::" + bucket,
		Principal: principals,
	}
	getStatement := Statement{
		Sid:    "Allow source account to Get Objects in pipeline bucket",
		Effect: "Allow",
		Action: "s3:GetObject",
		Resource: "arn:aws:s3:::" + bucket + "/*",
		Principal: principals,
	}

	var statementsSlice []Statement

	statementsSlice = append(statementsSlice, listStatement, getStatement)

	resp := Policy{
		Version:   "2012-10-17",
		Id:        "SSEAndSSLPolicy",
		Statement: statementsSlice,
		}
	js, _ := json.Marshal(resp)
	return string(js)
}