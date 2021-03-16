TARGET := iotdemo
SRCDIR := $(PWD)
CC := go

#REPO := registry.cn-shanghai.aliyuncs.com/advantech-k8s/
REPO := harbor.arfa.wise-paas.com/iotdemo/
TAG := v1.0.0
FLAGS := -ldflags "-s -w"

all: iotdemo-bin

%-image:
	docker build -t $(REPO)$*:$(TAG) -f $(SRCDIR)/cmd/$*/Dockerfile .
	docker push $(REPO)$*:$(TAG)

%-strip:
	$(CC) build $(FLAGS) -o $(SRCDIR)/bin/$* $(SRCDIR)/cmd/$*

%-bin:
	$(CC) build -mod=vendor -o $(SRCDIR)/bin/$* $(SRCDIR)/cmd/$*

.PHONY:clean
clean:
	rm -rf $(SRCDIR)/bin

