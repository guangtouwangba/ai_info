# Open-Source LLM-Driven Federated Transformer for Predictive IoV Management

**URL**: http://arxiv.org/abs/2505.00651v2

## 原始摘要

The proliferation of connected vehicles within the Internet of Vehicles (IoV)
ecosystem presents critical challenges in ensuring scalable, real-time, and
privacy-preserving traffic management. Existing centralized IoV solutions often
suffer from high latency, limited scalability, and reliance on proprietary
Artificial Intelligence (AI) models, creating significant barriers to
widespread deployment, particularly in dynamic and privacy-sensitive
environments. Meanwhile, integrating Large Language Models (LLMs) in vehicular
systems remains underexplored, especially concerning prompt optimization and
effective utilization in federated contexts. To address these challenges, we
propose the Federated Prompt-Optimized Traffic Transformer (FPoTT), a novel
framework that leverages open-source LLMs for predictive IoV management. FPoTT
introduces a dynamic prompt optimization mechanism that iteratively refines
textual prompts to enhance trajectory prediction. The architecture employs a
dual-layer federated learning paradigm, combining lightweight edge models for
real-time inference with cloud-based LLMs to retain global intelligence. A
Transformer-driven synthetic data generator is incorporated to augment training
with diverse, high-fidelity traffic scenarios in the Next Generation Simulation
(NGSIM) format. Extensive evaluations demonstrate that FPoTT, utilizing
EleutherAI Pythia-1B, achieves 99.86% prediction accuracy on real-world data
while maintaining high performance on synthetic datasets. These results
underscore the potential of open-source LLMs in enabling secure, adaptive, and
scalable IoV management, offering a promising alternative to proprietary
solutions in smart mobility ecosystems.


## AI 摘要

为解决车联网（IoV）中实时性、扩展性和隐私保护等挑战，本研究提出联邦提示优化交通转换器（FPoTT）框架。该框架利用开源大语言模型（LLM），通过动态提示优化机制提升轨迹预测精度，采用双层联邦学习架构（边缘轻量模型实时推理+云端LLM全局智能），并集成基于Transformer的合成数据生成器增强训练。测试显示，FPoTT使用Pythia-1B模型在真实数据中达到99.86%预测准确率，同时保持对合成数据集的高性能，为车联网管理提供了安全、自适应且可扩展的开源解决方案替代方案。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T13:10:23Z
- **目录日期**: 2025-05-14
