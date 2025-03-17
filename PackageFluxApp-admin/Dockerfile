# 基础镜像
FROM ubuntu:20.04
# 工作目录
WORKDIR /app

# 把编译后的打包进来这个镜像，放到工作目录 /app。你随便换
COPY package /app/package
COPY config-dev.yaml /app/config-dev.yaml
# CMD 是执行命令
# 最佳
ENTRYPOINT ["/app/package"]