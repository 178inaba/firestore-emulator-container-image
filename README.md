# Firestore Emulator Container Image

[![Publish Docker Image](https://github.com/178inaba/firestore-emulator-container-image/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/178inaba/firestore-emulator-container-image/actions/workflows/docker-publish.yml)

## Usage

```console
$ docker run -d --name datastore-emulator -e DATABASE_MODE=datastore-mode -p 8080:8080 ghcr.io/178inaba/firestore-emulator
```

### GitHub Actions

```yaml
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

### Docker Compose

```yaml
services:
  firestore:
    image: ghcr.io/178inaba/firestore-emulator
    ports:
      - 8080:8080
    healthcheck:
      test: ['CMD', 'curl', '-f', 'localhost:8080']

  datastore:
    image: ghcr.io/178inaba/firestore-emulator
    ports:
      - 8081:8080
    environment:
      DATABASE_MODE: datastore-mode
    healthcheck:
      test: ['CMD', 'curl', '-f', 'localhost:8080']

  app:
    ...
    environment:
      FIRESTORE_EMULATOR_HOST: firestore:8080
      DATASTORE_EMULATOR_HOST: datastore:8080
    depends_on:
      firestore:
        condition: service_healthy
      datastore:
        condition: service_healthy
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
