# MESSENGER CLEANER

This app will automate to delete your unwanted threads on Facebook - all of them.

## Dependencies

In order to use this application, you have to install the dependency.

### PhantomJS

A headless webkit browser that will make the automation possible.

Installation instruction, [please follow this link](http://phantomjs.org/download.html)

And make sure that the phantomjs is available to your %PATH% (for windows uers only)

## Usage

In order to use the application, you have to fill out the **settings.json** file.

For the account object of the json file, you have to input your username and password on facebook so the app will automatically login for your behalf.

The messages object, under excluded object, if there's a conversations you don't want to be deleted, please put their names. It should be comma-separated.

Please check below the example filled out settings.json

```json
{
    "account": {
        "username": "your_username@mail.com",
        "password": "password1q2w@@"
    },
    "messages": {
        "excluded": [
            "John Doe",
            "John McNeil"
        ]
    }
}
```
Then run the messengerdeleter app.

#### Windows Users

Just double click the messengerdeleterapp.exe

#### Mac OSX and Linux Users

Just run the messengerdeleterapp on terminal.

```bash
$ ./messengerdeleterapp
```

## Building

You could checkout this repository and build the tool.

```bash
$ GOOS=windows GOARCH=386 go build -o dist/win32-messengerdeleter/messengerdeleter.exe *.go
$ GOOS=windows GOARCH=amd64 go build -o dist/win64-messengerdeleter/messengerdeleter.exe *.go
$ GOOS=linux GOARCH=386 go build -o dist/linux32-messengerdeleter/messengerdeleter *.go
$ GOOS=linux GOARCH=amd64 go build -o dist/linux64-messengerdeleter/messengerdeleter *.go
```

And don't forget to copy-and-paste the **settings.json** file on each distribution folders.