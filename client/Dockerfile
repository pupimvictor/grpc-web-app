FROM grpcweb/common

ADD src/index.html /var/www/html
ADD dist/main.js /var/www/html/dist

EXPOSE 8081
WORKDIR /var/www/html
CMD ["python", "-m", "SimpleHTTPServer", "8081"]