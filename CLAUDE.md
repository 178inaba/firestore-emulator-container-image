# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Docker container images for Google Cloud emulators (Firestore/Datastore and Pub/Sub). This repository contains no application code — only Dockerfiles and CI/CD workflows.

## Build and Test

No build system or test framework. Everything runs through Docker and GitHub Actions.

### Firestore Emulator

```bash
# Build
docker build -t firestore-emulator firestore/

# Run (Firestore Native mode)
docker run -p 8080:8080 firestore-emulator

# Run (Datastore mode)
docker run -e DATABASE_MODE=datastore-mode -p 8080:8080 firestore-emulator

# Health check
curl -f localhost:8080
```

### Pub/Sub Emulator

```bash
# Build
docker build -t pubsub-emulator pubsub/

# Run
docker run -p 8085:8085 pubsub-emulator

# Health check
curl -f localhost:8085
```

## Repository Structure

- `firestore/Dockerfile` — Firestore emulator image (port 8080)
- `pubsub/Dockerfile` — Pub/Sub emulator image (port 8085)
- `.github/workflows/` — CI/CD workflows
  - `firestore-publish.yml` — Build and publish Firestore image on merge to main (only on `firestore/` path changes)
  - `pubsub-publish.yml` — Build and publish Pub/Sub image on merge to main (only on `pubsub/` path changes)
  - `test.yml` — Build and health-check test both emulator images on PR
- `examples/docker-compose/` — Example Go application using Firestore/Datastore

## CI/CD

- **PR tests**: Build image, start container, verify HTTP 200 via curl health check
- **Publish**: On merge to main, push multi-platform (amd64/arm64) images to both Docker Hub (`178inaba/*`) and GHCR (`ghcr.io/178inaba/*`)
- **Image tags**: `latest`, gcloud version number (e.g., `555.0.0`), commit SHA
- **Dependabot**: Daily automated updates for gcloud base images and GitHub Actions

## Important Notes

- Firestore Dockerfile uses `sh -c` to run the command because it needs shell expansion for the `DATABASE_MODE` environment variable
- Pub/Sub emulator requires the `beta` subcommand: `gcloud beta emulators pubsub start`
- Publish workflows use path filters — they only trigger on changes to their respective Dockerfile and corresponding workflow files
