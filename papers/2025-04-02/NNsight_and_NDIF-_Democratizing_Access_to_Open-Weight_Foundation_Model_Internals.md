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

NNsight和NDIF是两项协同工作的技术，旨在支持对大型神经网络内部表示和计算机制的科学研究。NNsight是一个开源系统，扩展了PyTorch以支持远程延迟执行；NDIF则是一个可扩展的推理服务，执行NNsight请求并共享GPU资源和预训练模型。其核心创新是"干预图"架构，将实验设计与模型运行时解耦。该框架为研究大型语言模型等深度神经网络提供了透明高效的内部访问方案，无需承担定制化模型的部署成本。研究显示当前AI领域对大规模模型内部机制的研究存在缺口，该技术正致力于填补这一空白。相关资源已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T06:02:29Z
- **目录日期**: 2025-04-02
