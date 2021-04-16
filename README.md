# Garbage Factory 1.0
Docker image that generates disk garbage for testing purposes.

## To create the image execute
```
  > sh build.sh
```

## Environment Variables

**GARBAGE_FACTORY_FILE_SIZE**

The size of each file in megabytes.

**GARBAGE_FACTORY_FILES_TO_CREATE**

The amount of files to create.
The files are named as follow: load-0, load-1, load-2 ...

**GARBAGE_FACTORY_FOLDER**

The base path where the files should be saved.

**GARBAGE_FACTORY_KEEP_ALIVE**

By default, after the files creation the docker image stops running.
To avoid this, you can set this flag to true in order to keep it running.
