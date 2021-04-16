# Garbage Factory 1.0
Docker image that generates disk garbage for testing purposes.

![alt text](https://github.com/lucaskenda/garbagefactory/blob/main/garbagefactory.png?raw=true)

## To create the image execute
```
  > sh build.sh
```

## Environment Variables

**GARBAGE_FACTORY_FILE_SIZE**

The size of each file in MiB (default: 100MiB).

**GARBAGE_FACTORY_FILES_TO_CREATE**

The amount of files to create (default: 1).
The files are named as follow: load-0, load-1, load-2 ...

**GARBAGE_FACTORY_FOLDER**

The base path where the files should be saved (default: ./garbage)

**GARBAGE_FACTORY_KEEP_ALIVE**

By default, after the files creation the docker image stops running.
To avoid this, you can set this flag to true in order to keep it running.
