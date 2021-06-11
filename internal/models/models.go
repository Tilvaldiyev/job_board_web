package models

import "os"

type ApplyJob struct {
	FirstName   string
	Email       string
	WebsiteLink string
	Portfolio   *os.File
	Coverletter string
}
