package awscommands

import (
	"RUSLANGO/commands"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bwmarrin/discordgo"
)

var bucket string
var sess *session.Session
var region = "us-east-2"

func init() {
	bucket = "ruslanbot"

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ = session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

}
func downloadFromS3Bucket(item string) string {
	//downloader for s3 items
	downloader := s3manager.NewDownloader(sess)
	//for some reason when you create a file with full item name/path it gives you an error
	ss := strings.Split(item, "/")
	fileName := ss[len(ss)-1] //to get just the file name without folders
	file, err := os.Create("tmp/" + fileName)
	if err != nil {
		fmt.Println("error creating a file", fileName, err)
	}
	defer file.Close()
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(strings.ToLower(item)),
		})
	if err != nil {
		fmt.Println("error downloading the file", item, err)
	} else {
		fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	}
	return file.Name()
}

func downloadFromS3BucketFolder(folder string) string {

	//lister (?) to list items in a bucket
	svc := s3.New(sess)
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		//because when I created the bucket I put the files into subfolders
		Prefix: aws.String(strings.ToLower(folder)),
	})

	if err != nil {
		fmt.Println("Unable to list items in bucket", bucket, err)
	}

	//get a random item from the list of items in the bucket
	items := resp.Contents
	item := *commands.GetRandomItem(items).Key
	return downloadFromS3Bucket(item)
}

func SendRandomFileFromFolder(s *discordgo.Session, c string, folder string) {

	//download file from s3 bucket and folder, returns the name of the file
	fileName := downloadFromS3BucketFolder(folder)
	sendFile(s, c, fileName)
}

func SendFileFromS3(s *discordgo.Session, c string, item string) {
	fileName := downloadFromS3Bucket(item)
	sendFile(s, c, fileName)
}

func sendFile(s *discordgo.Session, channel string, fileName string) {
	//open file to give it to discord
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening a file "+file.Name(), err)
	}
	s.ChannelFileSend(channel, filepath.Base(file.Name()), file)
	//to not leave any leftovers
	os.Remove(file.Name())
}
