#!/bin/bash

# 配置
local_program_path="/Users/fcihpy/Desktop/anypay-monorepo/script/build/contractBot"     # 本地程序路径
remote_server="ubuntu@IM"      # 远程服务器地址
remote_program_name="/opt/contract/contractBot"         # 远程服务器程序路径

# 执行本地程序的打包
echo "执行本地程序的打包..."
make all

# 检查 make 命令的返回值，如果成功（返回值为0），则执行 scp 命令
if [ $? -eq 0 ]; then
    # 传输程序到远程服务器
    echo "传输程序到远程服务器..."
    ssh -t "$remote_server" "sudo pkill contractBot"
    scp "$local_program_path" "$remote_server:$remote_program_name"
else
    echo "make all 命令执行失败"
fi

# 启动新的程序
echo "执行重启进程操作..."
ssh -t "$remote_server" "/opt/contract/restart.sh"

echo "部署完成"

