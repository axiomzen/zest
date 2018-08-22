all:
	echo "make install to install zest"

install:
#	docker: Error response from daemon: Mounts denied:
#	The path /usr/local/bin/_zester
#	is not shared from OS X and is not known to Docker.
# 	You can configure shared paths from Docker -> Preferences... -> File Sharing.
#	See https://docs.docker.com/docker-for-mac/osxfs/#namespaces for more info.

	install -d $(HOME)/.zest
	install zest /usr/local/bin/zest
	install _zester $(HOME)/.zest/_zester

uninstall:
	rm -rf /usr/local/share/zest
	rm -rf $(HOME)/.zest
	rm /usr/local/bin/zest
	rm /usr/local/bin/_zester
