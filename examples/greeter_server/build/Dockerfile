FROM ubuntu:20.04

WORKDIR /data

COPY greeter_server ./

# 固定环境变量注入
# LIBRA_API_ACCESS_POINT libra接口地址
CMD ["sh", "-c", "./greeter_server -api $LIBRA_API_ACCESS_POINT"]

