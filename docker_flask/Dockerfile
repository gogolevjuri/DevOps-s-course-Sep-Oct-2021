FROM python:3.9-alpine AS build-stage
RUN mkdir -p ./templates/
COPY requirements.txt flask_app.py ./
COPY 404 game ./templates/
RUN pip3 install -r requirements.txt && \
    rm -rf /.cache/* /tmp/* /usr/local/lib/python*/multiprocessing/* /usr/local/lib/python*/pydoc_data/* /usr/local/lib/python*/*/pkg_resources/* /usr/local/lib/python*/turtledemo/* usr/local/lib/python*/__pycache__/* /usr/local/lib/python*/*/pip/* /usr/local/lib/python*/*/wheel/* /usr/local/lib/python*/*/setuptools/* /usr/local/lib/python*/xml/* /usr/local/lib/python*/xmlrpc/* /usr/lib/python*/ensurepip /usr/local/include/python*/* /usr/local/lib/python*/concurrent/* /usr/local/lib/python*/unittest/* /usr/local/lib/python*/*/werkzeug/debug* /usr/local/lib/python*/*/werkzeug/middleware* /usr/local/lib/python*/idlelib/* /usr/local/lib/python*/ensurepip 
FROM alpine
COPY --from=build-stage . .
ENTRYPOINT ["python3"]
CMD ["flask_app.py"]
