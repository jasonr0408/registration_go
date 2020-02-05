FROM nginx
COPY dist/ /usr/share/nginx/html/
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY nginx/example.crt /etc/nginx/ssl/nginx.crt
COPY nginx/example.key /etc/nginx/ssl/nginx.key
