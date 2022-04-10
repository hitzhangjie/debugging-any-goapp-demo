# README

---

- Test in docker container: `docker run -it --cap-add=SYS_PTRACE --rm -v $(pwd):/workspace hitzhangjie/linux101 /bin/bash`
- After container started: `cd /workspace`
- Then run `make` to build and run the test
- `Ctrl+C` to terminate the test.

---

This code works on Linux/amd64, while not on darwin/amd64.
