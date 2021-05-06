# platform guide todos
1. Define how ModuleName and proto module is defined
2. Naming RPC methods in proto files
3. Only define rest routs when GQL in not sufficient
4. Service Proto files should have messages which we can expose to public, no internal fields or messges should be exposed.
5. Do not hardcode keys, base urls etc, which could be different for different environments
6. Code will never have any conditional cases which can depend on different environments
7. Internal Services will never call a public registered service
8. Internal to internal service call will never involve GRPC dial
9. Protofile name should be sigular
10. Singular name for module no suffix "Service" needed in module and in proto file as well as in package name
11. each RPC would be implemented in seperate file, and will have same name as RPC in kebabcase
12. Internal services proto will follow same conventions as the main protos, except proto module will have '.intenals' suffix in them 
13. Internal protos can include messages from main protos but can not the other way around
14 Everyone install editorconfig plugins in IDEs

. Only include what is needed in service module




# Platform Framework code changes
1. Remove Custom errors defined in GRPC server, use all defaults
2. Add proto linter with buff

# Service Module code changes


# Database convention 
1. Database model names are always singular
2. Define Model structs seperately always
3. CreatedAt UpdatedAt DeletedAt, always have these fields in all tables
4. Encrypt all sensitive fields, always!!
5. TODO: Support incremental migration files
6. Add validation in all fields!!
7. Have seperate struct file for each model
8. Model file name should be same as struct name in kebabcase