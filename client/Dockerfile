FROM node:latest

# Set working directory to /app folder
WORKDIR /app 

# At first copy the package.json folder & Run "npm install"
# to all the required dependencies ready
COPY package.json package-lock.json ./

# Running "npm install" to install all dependencies
RUN npm install 

# Now copy the remaining code to working directory
COPY . ./

# Building the react app
RUN npm run build

FROM nginx:alpine

# Copying the build from /app/build to nginx/html folder
COPY --from=build /app/build /usr/share/nginx/html

# Exposing port 80 (exposed by nginx)
EXPOSE 80

# Command to run on container start/run
CMD ["nginx", "-g", "daemon off;"]