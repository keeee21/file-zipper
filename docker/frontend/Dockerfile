FROM node:22-alpine

#-------------------------------------------------
# Set timezone to Asia/Tokyo (for Alpine Linux)
#-------------------------------------------------
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "Asia/Tokyo" > /etc/timezone

#-------------------------------------------------
# Install pnpm
#-------------------------------------------------
RUN npm install -g pnpm

ENV PATH=./node_modules/.bin:$PATH

CMD ["/usr/bin/bash"]