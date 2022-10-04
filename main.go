package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

var Version = "0.0"

func main() {
	region := flag.String("region", "us-east-1", "Region to Use")
	version := flag.Bool("version", false, "print version of aws-env")
	flag.Parse()

	if *version {
		fmt.Printf("%s version %s\n", os.Args[0], Version)
		return
	}

	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Region: region,
			},
		),
	)

	ecrClient := ecr.New(sess)

	authToken, err := ecrClient.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
	if nil != err {
		panic(err)
	}

	endpointsAsByteArray, err := base64.StdEncoding.DecodeString(*authToken.AuthorizationData[0].AuthorizationToken)
	if nil != err {
		panic(err)
	}

	elements := strings.SplitN(string(endpointsAsByteArray), ":", 2)

	command := make([]string, 0)

	command = append(command, "docker", "login", "-u", elements[0], "-p", elements[1], *authToken.AuthorizationData[0].ProxyEndpoint)

	fmt.Println(strings.Join(command, " "))
}
