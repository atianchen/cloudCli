# cloudCli
基于GO的持续化集成工具
   ________                __   _________ 
  / ____/ /___  __  ______/ /  / ____/ (_)
 / /   / / __ \/ / / / __  /  / /   / / / 
/ /___/ / /_/ / /_/ / /_/ /  / /___/ / /  
\____/_/\____/\__,_/\__,_/   \____/_/_/  

proto生成：protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\proto\simple.proto