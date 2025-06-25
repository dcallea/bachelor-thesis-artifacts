#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main() {
	setuid(0);

	char dir[256];

	printf("Enter directory to list: ");
	fflush(stdout);

	if (fgets(dir, sizeof(dir), stdin) == NULL) {
		fprintf(stderr, "Error reading input\n");
		return 1;
	}

	dir[strcspn(dir, "\n")] = 0;

	char cmd[512];
	snprintf(cmd, sizeof(cmd), "ls %s", dir);

	int ret = system(cmd);
	if (ret != 0) {
		printf("Error running command\n");
	}

	return 0;
}
