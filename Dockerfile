FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:emulators

ENV DATABASE_MODE=firestore-native

EXPOSE 8080

CMD ["sh", "-c", "gcloud emulators firestore start --database-mode $DATABASE_MODE --host-port 0.0.0.0:8080"]
