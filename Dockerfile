FROM       scratch
ADD        books books
ENV        PORT 8080
EXPOSE     8080
ENTRYPOINT ["/books"]