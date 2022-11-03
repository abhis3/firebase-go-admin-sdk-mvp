package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cloudStorage "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"

	//"firebase.google.com/go/auth"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Run the file with the command `go run hello.go`
func main() {
	// Set Environment Variables according to https://firebase.google.com/docs/emulator-suite/connect_storage#web-v9
	// COMMENT OUT THESE ENV VAR SETTERS IF YOU WANT TO TALK TO PROD
	err := os.Setenv("GCLOUD_PROJECT", "fir-photo-editor-8cc11")
	if err != nil {
		log.Fatalln(err)
	}

	err2 := os.Setenv("FIREBASE_STORAGE_EMULATOR_HOST", "localhost:9199")
	if err2 != nil {
		log.Fatalln(err2)
	}

	// Note: This is a bug in the current GO Admin SDK. It requires this environment
	// variable to be set to talk to the emulator despite the official documentation
	// stating otherwise. Issue is raised in the GO SDK repo here: TODO
	err3 := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9199")
	if err3 != nil {
		log.Fatalln(err3)
	}

	// Project Configs
	config := &firebase.Config{
		StorageBucket: "fir-photo-editor-8cc11.appspot.com", // REPLACE with your default bucket name
	}
	ctx := context.Background()

	// Uncomment this line and pass it in to the below line instead of `option.WithoutAuthentication()` to talk to prod
	// Note that you'll need the service account JSON that's generated in the Firebase Console in the correct directory to make this work

	// opt := option.WithCredentialsFile("/Users/abhisun/Documents/github/emulator-issue-go/fir-photo-editor-8cc11-firebase-adminsdk-kcf9w-16da57a892.json")

	app, err := firebase.NewApp(context.Background(), config, option.WithoutAuthentication())
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Some storage specific logic to fetch the default bucket

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	// If you want to run specific query, use the below lines and comment out the 2 lines below them

	//query := &cloudStorage.Query{Prefix: "images/"}
	//it := bucket.Objects(ctx, query)

	query := &cloudStorage.Query{Prefix: ""}
	it := bucket.Objects(ctx, query)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		// If you want to delete the object, feel free to use this code block

		// err = bucket.Object(attrs.Name).Delete(ctx)
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		fmt.Println(attrs.Name)
	}

	// END OF PROGRAM
	fmt.Println("============")
	fmt.Println("END")
}
