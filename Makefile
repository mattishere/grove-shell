# Intended for use on Linux machines.

# Local installation
install:
	@go build -ldflags "-s -w" -o grove ./cmd/grove-shell/
	@install -Dm755 grove $(DESTDIR)/usr/local/bin/grove
	@echo "Installed Grove."

uninstall:
	@rm -rf $(DESTDIR)/usr/local/bin/grove
	@echo "Uninstalled Grove."

register:
	@echo "$(DESTDIR)/usr/local/bin/grove" | tee -a /etc/shells > /dev/null
	@echo "Registered Grove as a shell."


# Remove the binary
clean:
	@rm -f grove
	@echo "Cleaned the local Grove binary."



# Installation & uninstallation for package managers only
global_install:
	@go build -ldflags "-s -w" -o grove
	@install -Dm755 grove $(DESTDIR)/usr/bin/grove

global_uninstall:
	@rm -rf $(DESTDIR)/usr/bin/grove
