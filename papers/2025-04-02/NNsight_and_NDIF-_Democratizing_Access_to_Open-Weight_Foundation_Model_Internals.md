# NNsight and NDIF: Democratizing Access to Open-Weight Foundation Model Internals

**URL**: http://arxiv.org/abs/2407.14561v4

## 原始摘要

We introduce NNsight and NDIF, technologies that work in tandem to enable
scientific study of the representations and computations learned by very large
neural networks. NNsight is an open-source system that extends PyTorch to
introduce deferred remote execution. The National Deep Inference Fabric (NDIF)
is a scalable inference service that executes NNsight requests, allowing users
to share GPU resources and pretrained models. These technologies are enabled by
the Intervention Graph, an architecture developed to decouple experimental
design from model runtime. Together, this framework provides transparent and
efficient access to the internals of deep neural networks such as very large
language models (LLMs) without imposing the cost or complexity of hosting
customized models individually. We conduct a quantitative survey of the machine
learning literature that reveals a growing gap in the study of the internals of
large-scale AI. We demonstrate the design and use of our framework to address
this gap by enabling a range of research methods on huge models. Finally, we
conduct benchmarks to compare performance with previous approaches.
  Code, documentation, and tutorials are available at https://nnsight.net/.


## AI 摘要

NNsight和NDIF是两项协同工作的技术，旨在支持对大型神经网络内部表征和计算的科学研究。NNsight是一个开源系统，扩展了PyTorch以支持远程延迟执行；NDIF则是一个可扩展的推理服务，执行NNsight请求，实现GPU资源和预训练模型的共享。它们基于"干预图"架构，将实验设计与模型运行解耦，为研究大型语言模型等深度神经网络内部机制提供了透明高效的访问方式，无需单独托管定制模型。该框架填补了大规模AI内部研究日益扩大的空白，支持多种研究方法，并展示了优于以往方案的性能表现。详见nnsight.net。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T19:01:54Z
- **目录日期**: 2025-04-02
