# Firestore Emulator Container Image

## Usage

```console
$ docker run -d --name datastore-emulator -e DATABASE_MODE=datastore-mode -p 8080:8080 ghcr.io/178inaba/firestore-emulator
```

### GitHub Actions

```yml
jobs:
  test:
    runs-on: ubuntu-latest
    services:
      datastore-emulator:
        image: ghcr.io/178inaba/firestore-emulator
        ports:
          - 8080:8080
        env:
          DATABASE_MODE: datastore-mode
    env:
      DATASTORE_EMULATOR_HOST: localhost:8080
    steps:
      - name: Test
        run: |
          # Replace this with running tests for your app.
          curl -f localhost:8080
```

## Environment Variables

### `DATABASE_MODE`

The database mode to start the Firestore Emulator in.

The valid options are:

- `firestore-native` (default): start the emulator in Firestore Native
- `datastore-mode`: start the emulator in Datastore Mode

## License

[MIT](LICENSE)

## Author

Masahiro Furudate (a.k.a. [178inaba](https://github.com/178inaba))  
<178inaba.git@gmail.com>
