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

NNsight与NDIF是一套协同技术，旨在支持大规模神经网络的内部表征与计算机制研究。NNsight是基于PyTorch的开源系统，支持延迟远程执行；NDIF则是可扩展的推理服务平台，提供共享GPU资源与预训练模型。核心技术"干预图"架构将实验设计与模型运行解耦，使研究者无需自行托管定制模型即可透明高效地访问大语言模型等深度神经网络内部状态。研究显示当前AI领域对大规模模型内部机制的研究存在明显缺口，该框架通过支持多种研究方法填补这一空白。性能测试表明其优于传统方案，相关资源已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T10:02:14Z
- **目录日期**: 2025-04-02
