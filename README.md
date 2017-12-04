# Image Spy

![imagespy](https://raw.githubusercontent.com/imagespy/imagespy/master/imagespy.gif)

## Usage

### Docker

```
$ docker pull imagespy/imagespy:latest
$ docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8888:8888 imagespy/imagespy:latest
# Open http://localhost:8888/ in your browser
```

## Development

Compile the server and run it:

```
$ cd app
$ go build
$ ./app
```

Start the UI:

```
$ cd ui
$ yarn install
$ yarn start
```
