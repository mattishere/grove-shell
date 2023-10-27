# Intended for use on Linux machines.

# Local installation
install:
	@go build -ldflags "-s -w" -o grove ./cmd/grove-shell/
	@install -Dm755 grove $(DESTDIR)/usr/local/bin/grove

uninstall:
	@rm -rf $(DESTDIR)/usr/local/bin/grove

# Remove the binary
clean:
	@rm -f grove



# Installation & uninstallation for package managers only
global_install:
	@go build -ldflags "-s -w" -o grove
	@install -Dm755 grove $(DESTDIR)/usr/bin/grove

global_uninstall:
	@rm -rf $(DESTDIR)/usr/bin/grove
