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

NNsight和NDIF是两项协同工作的技术，旨在支持对大型神经网络内部表征与计算的科学研究。NNsight是一个开源系统，扩展了PyTorch以实现延迟远程执行；NDIF则是一个可扩展的推理服务，执行NNsight请求，支持共享GPU资源和预训练模型。该框架通过"干预图"架构解耦实验设计与模型运行时，为研究者提供透明高效的神经网络内部访问（如大语言模型），无需承担定制化模型的托管成本。研究显示当前机器学习文献中大型AI内部机制研究存在缺口，该框架通过支持多种研究方法填补这一空白。代码和文档详见nnsight.net。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T16:02:06Z
- **目录日期**: 2025-04-02
