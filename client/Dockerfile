FROM grpcweb/common

ARG EXAMPLE_DIR=/github/grpc-web/net/grpc/gateway/examples/echo

RUN protoc -I=$EXAMPLE_DIR echo.proto \
--js_out=import_style=commonjs:\
$EXAMPLE_DIR/commonjs-example \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:\
$EXAMPLE_DIR/commonjs-example

RUN cd $EXAMPLE_DIR/commonjs-example && \
  npm install && \
  npx webpack && \
  cp echotest.html /var/www/html && \
  cp dist/main.js /var/www/html/dist

EXPOSE 8081
WORKDIR /var/www/html
CMD ["python", "-m", "SimpleHTTPServer", "8081"]