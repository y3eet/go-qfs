# Go QFS(Quick File Share)

A lightweight desktop file-sharing application built in Go. Run the binary in any directory to instantly share its contents with any device on the same network. no installation, no configuration.

## Usage

1. Place the executable file in the directory you want to share
2. Run it
3. A local HTTP server starts and the app displays a QR code and connection address
4. Connect from any device on the same network and scan the QR code or enter the address in a browser
   > [!NOTE]
   > You may need to expose your port if you are using a firewall
5. Upload or download files directly from that shared directory

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
