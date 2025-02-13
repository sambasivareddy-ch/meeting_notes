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

# Expose 3000
EXPOSE 3000

CMD [ "npm", "start" ]