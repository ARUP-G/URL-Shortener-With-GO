# Use the official Nginx image from the Docker Hub
FROM nginx:alpine

# Replace the default index.html with your own
COPY ./frontend/index.html /usr/share/nginx/html/index.html

# Expose port 80 to the outside world
EXPOSE 80

# Start Nginx when the container has started
CMD ["nginx", "-g", "daemon off;"]
