# 修改文件名为小写
# sudo find . -type f -name "*.go" -depth -exec sh -c 'mv "$0" "$(dirname "$0")/$(basename "$0" | tr "[:upper:]" "[:lower:]")"' {} \;

# 查看文件名
sudo find . -type f -name "solution*"