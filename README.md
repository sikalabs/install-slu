<p align="center">
  <h1 align="center">install-slu: SikaLabs Utils Installer</h1>
  <p align="center">
    <a href="https://opensource.sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/OPENSOURCE BY-SIKALABS-131480?style=for-the-badge"></a>
    <a href="https://sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-sikalabs.com-gray?style=for-the-badge"></a>
    <a href="mailto:opensource@sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-opensource@sikalabs.com-gray?style=for-the-badge"></a>
  </p>
</p>

## Install

### Linux AMD64

```bash
# Check the current version on Github https://github.com/sikalabs/install-slu/releases
VERSION=v0.1.0 && \
OS=linux && \
ARCH=amd64 && \
BIN=install-slu && \
curl -L https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xvzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
mv ${BIN} /usr/local/bin/
```
