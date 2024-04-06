FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:emulators

LABEL org.opencontainers.image.description Local Firestore emulator.

ENV DATABASE_MODE=firestore-native
ENV HOST_PORT=0.0.0.0:8080

CMD ["sh", "-c", "gcloud emulators firestore start --database-mode $DATABASE_MODE --host-port $HOST_PORT"]
