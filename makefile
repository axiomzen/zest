all:
	echo "make install to install zest"

install:
	install -d /usr/local/share/zest
	install zest /usr/local/bin/zest
	install _zester /usr/local/bin/_zester
	install -m 644 Zestfile.example /usr/local/share/zest/Zestfile.example

uninstall:
	rm -rf /usr/local/share/zest
	rm /usr/local/bin/zest
	rm /usr/local/bin/_zester
