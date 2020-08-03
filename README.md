### 常用git操作
```shell
#初始化一个git本地仓库
git init
touch README.md
#添加git管理
git add README.md

#提交
git commit -m "first commit"

#删除git 远程仓库地址
git remote rm origin

#添加git 远程仓库地址
git remote add origin https://github.com/nicklbx/shiyanlou-test.git

#推送到远程仓库
git push -u origin master


#查看所有分支
git branch -a

#获取最新分支
git pull -a

#获取tag
git tag

#切换分支
git checkout branch-v0.1

#强制清空本地修改
git reset --hard

```
