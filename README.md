# E-commerce API build with Go lang.

## Pre-requisites

- Must have go lang installed in your local machine. 
- Create an account in Mongodb atlast and start a free server. ( for ref : https://www.youtube.com/watch?v=JNr5noDp6EM ) 
- Get POSTMAN for sending requests to API to check if it is working. ( Or you may use Thunderclient VSCode extension )

### Commands to input in terminal to make this code run : 

- ` go mod init github.com/<yourgitname>/<foldername> `

Is it necessary? If you have worked in node.js or build package.json file before, this is very similar but go lang style. 
It contains your path to your modules. you can have any path instead of 'github.com/...' but since it is this way in our project
so it's recommended just to make it run.

- ` export = GO111MODULE="on" `
- ` go mod tidy `
- ` go build `
- ` go run main.go `

That's it! make any changes and implement any new feature you want. ~happy coding.
