#List of commands
# Create a service - new service
#example : kit new service user
kit new service service_name
#now create endpoints for gorilla
kit g s user -w --gorilla
#if we added new endpoints run and compile
kit g s todo --gorilla -w
# regenerate the file
# kit g s user

