package main

const PORT = "50001"
const APP_NAME = "Content Fetcher"
const GITHUB_API_CONTENTS_ENDPOINT = "https://api.github.com/repos/abyanmajid/codemore.io/contents/"

func main() {
	ListenAndServe()
}
