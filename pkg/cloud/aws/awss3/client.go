package awss3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/util/cmdutil"
	"os"
)

func S3Client() *s3.S3 {
	config := cmdutil.GetConfig("aws")
	if config.AccessKeyId == "" {
		log.Warnln("需要先配置访问凭证 (Access Key need to be configured first)")
		os.Exit(0)
		return nil
	} else {
		cfg := &aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(config.AccessKeyId, config.AccessKeySecret, config.STSToken),
		}
		sess := session.Must(session.NewSession(cfg))
		svc := s3.New(sess)
		log.Traceln("S3 Client 连接成功 (S3 Client connection successful)")
		return svc
	}
}
