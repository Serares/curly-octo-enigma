# docker file to crate the sqlite container
FROM alpine:latest
# Install SQLite
RUN apk --no-cache add sqlite
# Create a directory to store the database
WORKDIR /db
# Copy your SQLite database file into the container
COPY data/questions-app /db/
# Expose the port if needed
EXPOSE 8080:8080
# Command to run when the container starts
CMD ["sqlite3", "/data/questions-app"]
