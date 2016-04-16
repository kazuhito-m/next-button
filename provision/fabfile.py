#coding:utf-8
from fabric.api import local, run, sudo, put, env
from fabric.context_managers import cd
from fabric.contrib.project import rsync_project

env.shell="bash -c "

# 
# RaspberryPI セットアップ・プロビジョニングソース
# 
def setup_all():
#	all_upgrade()
#	japanize()
#	install_openjtalk()
#	basic_tools_setup()
#	install_nextbutton()
#	build_install_golang()
	build_nextbutton()
def japanize():
	## change locale
	# sudo("update-locale LANG=ja_JP.UTF-8")
	## change timezone
	# sudo("mv /etc/localtime{,.org}")
	# sudo("ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime")
	print("ここは今動いてません…。")

def all_upgrade():
	sudo("apt-get update", pty=False)
	# sudo("do-release-upgrade -d", pty=False)
	sudo("apt-get update", pty=False)
	sudo("apt-get upgrade -y", pty=False)
	sudo("apt-get dist-upgrade -y", pty=False)
def basic_tools_setup():
	sudo("apt-get install -y rsync git" , pty=False)
def install_openjtalk():
	sudo("apt-get install -y open-jtalk open-jtalk-mecab-naist-jdic hts-voice-nitech-jp-atr503-m001", pty=False)
	run("mkdir -p ./work")
	with cd("./work"):
		run("curl http://heanet.dl.sourceforge.net/project/mmdagent/MMDAgent_Example/MMDAgent_Example-1.6/MMDAgent_Example-1.6.zip > ./voicepack.zip")
		run("unzip voicepack.zip")
		sudo("mv ./MMDAgent*/Voice/mei /usr/share/hts-voice/")
		put("resources/scripts/readalound_text.bsh" , "/tmp/readalound_text.bsh")
		sudo("mv /tmp/readalound_text.bsh /usr/local/bin/readalound_text.bsh")
		sudo("chmod 755 /usr/local/bin/readalound_text.bsh")
	run("rm -rf ./work")
def install_nextbutton():
	sudo("rm -rf /usr/local/next-button/")
	sudo("rm -rf /usr/local/bin/next-button")
	rsync_project(
		local_dir="./resources/next-button",
		remote_dir="/tmp",
		exclude=[],
		delete=True
	)
	sudo("mv /tmp/next-button/ /usr/local/next-button/")
	sudo("chmod 755 /usr/local/next-button/*sh")
	sudo("ln -s /usr/local/next-button/next-button.bsh /usr/local/bin/next-button")
	# サービス的にsystemdに登録
	sudo("cp /usr/local/next-button/next-button.service /etc/systemd/system/next-button.service")
	sudo("systemctl enable next-button.service")
	sudo("systemctl stop next-button.service")
	sudo("systemctl start next-button.service")
def build_install_golang():
	sudo("rm -rf /usr/local/go")
	run("curl https://storage.googleapis.com/golang/go1.6.1.linux-armv6l.tar.gz > /tmp/go.tar.gz")
	run("tar -C /tmp -xzf /tmp/go.tar.gz")
	sudo("mv /tmp/go /usr/local/go")
	sudo("ln -s /usr/local/go/bin/go /usr/local/bin/go")
	# 最終的に、以下は「実行時指定する」から要らないのでは？ TODO サラで入れるときに検証してみる
#	put("./resources/bashrc_append.txt","/tmp/bashrc_append.txt")
#	run("cat /tmp/bashrc_append.txt >> ~/.bashrc")
def build_nextbutton():
	run("rm -rf ./go")
	run("mkdir -p ./go/src/github.com/kazuhito-m")
	run("git clone https://github.com/kazuhito-m/next-button.git ./go/src/github.com/kazuhito-m/next-button")
	with cd("./go/src/github.com/kazuhito-m/next-button"):
		run("GOPATH=~/go go get", pty=True)
		run("GOPATH=~/go go build", pty=True)
		run("chmod 755 ./next-button")
		sudo("cp ./next-button /usr/local/next-button/next-button")

