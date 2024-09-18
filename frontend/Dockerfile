# Use the official Nginx image from the Docker Hub
FROM nginx:alpine 

COPY nginx.conf /etc/nginx/nginx.conf

COPY static/images/home.mp4 /usr/share/nginx/html/static/images/home.mp4 
# Replace the default index.html with your own
COPY static/index.html /usr/share/nginx/html/static/index.html

# Expose port 80 to the outside world
EXPOSE 80

# Start Nginx when the container has started
CMD ["nginx", "-g", "daemon off;"]
