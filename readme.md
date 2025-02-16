# This is a repo to document all the work done going through the book "Mastering Go" by Mihalis Tsoukalous

## The program contains a few packages, where the relevant topics being covered are listed

- compiling
    - learning how to either compile files into statically linked binaries (files where all required dependencies are bundled into the binary) or running the files using go run

- downloading_packages
    - understanding how when working with external packages, it is required to download the actual files from the external repositories, and where these files are stored on the local machine

- using os.Stdout
    - details on how os.stdout, os.stdin and os.stderr are all available to go programs, and how they can be used wherever the interface io.Writer is used

- using os.Stdin
    - details on using bufio.NewScanner to read from os.Stdin, and how os.Stdin is just a file like every other file in the unix system. Because it is just a file, like every other File it has the Read method attached to it and so fulfils the io.Reader interface that must be passed into bufio.NewScanner as an argument

- getting arguments from command line arguments

- understanding how messages from programs can be written to either the stdout or stderror streams, and how each of these streams can be captured and their output can be written to files