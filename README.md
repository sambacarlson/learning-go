How to set up go compiler for notebook;

1. install golang jupyter kernel like so \

```bash
   go install github.com/janpfeifer/gonb@latest && \
   go install golang.org/x/tools/cmd/goimports@latest && \
   go install golang.org/x/tools/gopls@latest && \
   gonb  -- install
```

-- Not that you will need go/bin in your path. If not already added, then add it like so (be sure to specify your default shell profile. I am assuming zsh here);

```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && \
   source ~/.zshrc
```

2. restart jupyter notebook/editor

ALTERNATIVE

- pull and run docker image: `docker pull janpfeifer/gonb_jupyterlab:latest docker run -it - rm -p 8888:8888 -v :/home/jovyan/work janpfeifer/gonb_jupyterlab:latest`
